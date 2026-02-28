package ecerp

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// maxResponseBodySize 响应体最大允许大小 (10MB)
const maxResponseBodySize = 10 << 20

// Caller 定义了 SDK 的核心调用接口，各子模块依赖此接口
type Caller interface {
	// Do 执行 API 调用并将响应解析到 result 中
	Do(ctx context.Context, method string, bizContent interface{}, result interface{}) error
}

// compile-time interface check
var _ Caller = (*Client)(nil)

// Client 易仓ERP API客户端
//
// Client 是并发安全的，应当被复用而非每次请求创建新实例。
type Client struct {
	config *Config
	closed atomic.Bool
}

// NewClient 创建新的易仓ERP客户端
//
// 必填参数:
//   - appKey:    应用Key
//   - appSecret: 应用密钥
//   - serviceID: 授权服务ID
//
// 可选配置通过 Option 函数传入。
// 如果必填参数为空，则返回 nil 和 error。
func NewClient(appKey, appSecret, serviceID string, opts ...Option) (*Client, error) {
	if appKey == "" {
		return nil, errors.New("ecerp: appKey is required")
	}
	if appSecret == "" {
		return nil, errors.New("ecerp: appSecret is required")
	}
	if serviceID == "" {
		return nil, errors.New("ecerp: serviceID is required")
	}

	cfg := newDefaultConfig(appKey, appSecret, serviceID)
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.HTTPClient == nil {
		transport := &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		}
		cfg.HTTPClient = &http.Client{
			Timeout:   cfg.Timeout,
			Transport: transport,
		}
	}

	return &Client{config: cfg}, nil
}

// MustNewClient 创建客户端，失败时 panic。适用于初始化阶段。
func MustNewClient(appKey, appSecret, serviceID string, opts ...Option) *Client {
	c, err := NewClient(appKey, appSecret, serviceID, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// Do 执行 API 调用
//
// 参数:
//   - ctx:        上下文，用于超时控制和取消
//   - method:     接口方法名（interface_method，如 "getOrderList"）
//   - bizContent: 业务参数，会被序列化为 JSON（传 nil 为空对象）
//   - result:     响应 data 字段的目标结构体指针（传 nil 忽略返回数据）
func (c *Client) Do(ctx context.Context, method string, bizContent interface{}, result interface{}) error {
	resp, err := c.doRequest(ctx, method, bizContent)
	if err != nil {
		return err
	}

	// 检查业务错误
	if resp.Code != CodeSuccess {
		return &APIError{
			Code:    resp.Code,
			Message: resp.Message,
		}
	}

	// 解析业务数据
	if result != nil && len(resp.Data) > 0 {
		if err := json.Unmarshal(resp.Data, result); err != nil {
			return fmt.Errorf("ecerp: unmarshal data error: %w, raw: %s", err, truncateStr(string(resp.Data), 500))
		}
	}

	return nil
}

// DoRaw 执行 API 调用，返回原始响应
//
// 即使 API 返回非 200 状态码，Response 仍会被返回以供调试。
func (c *Client) DoRaw(ctx context.Context, method string, bizContent interface{}) (*Response, error) {
	resp, err := c.doRequest(ctx, method, bizContent)
	if err != nil {
		return nil, err
	}

	if resp.Code != CodeSuccess {
		return resp, &APIError{
			Code:    resp.Code,
			Message: resp.Message,
		}
	}

	return resp, nil
}

// doRequest 内部方法，执行 HTTP 请求
func (c *Client) doRequest(ctx context.Context, method string, bizContent interface{}) (*Response, error) {
	if method == "" {
		return nil, errors.New("ecerp: interface_method is required")
	}

	if c.closed.Load() {
		return nil, errors.New("ecerp: client is closed")
	}

	// 序列化 bizContent
	var bizContentStr string
	if bizContent != nil {
		bizBytes, err := json.Marshal(bizContent)
		if err != nil {
			return nil, fmt.Errorf("ecerp: marshal biz_content error: %w", err)
		}
		bizContentStr = string(bizBytes)
	} else {
		bizContentStr = "{}"
	}

	// 构建公共参数
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	nonceStr, err := generateNonceStr()
	if err != nil {
		return nil, fmt.Errorf("ecerp: generate nonce error: %w", err)
	}

	params := map[string]string{
		"app_key":          c.config.AppKey,
		"service_id":       c.config.ServiceID,
		"interface_method": method,
		"timestamp":        timestamp,
		"nonce_str":        nonceStr,
		"charset":          c.config.Charset,
		"version":          c.config.Version,
		"sign_type":        c.config.SignType,
		"biz_content":      bizContentStr,
	}

	// 生成签名
	sign := GenerateSign(params, c.config.AppSecret)

	// 构建请求体
	reqBody := Request{
		AppKey:          params["app_key"],
		ServiceID:       params["service_id"],
		InterfaceMethod: params["interface_method"],
		Timestamp:       params["timestamp"],
		NonceStr:        params["nonce_str"],
		Charset:         params["charset"],
		Version:         params["version"],
		SignType:        params["sign_type"],
		BizContent:      params["biz_content"],
		Sign:            sign,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("ecerp: marshal request error: %w", err)
	}

	// 调试日志
	if c.config.Logger != nil {
		c.config.Logger.Debug("ecerp: request",
			slog.String("method", method),
			slog.Int("body_size", len(bodyBytes)),
		)
	}

	// 发送 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.BaseURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("ecerp: create http request error: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	if c.config.UserAgent != "" {
		httpReq.Header.Set("User-Agent", c.config.UserAgent)
	}

	httpResp, err := c.config.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("ecerp: http request error: %w", err)
	}
	defer httpResp.Body.Close()

	// 检查 HTTP 状态码
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(httpResp.Body, 1024))
		return nil, fmt.Errorf("ecerp: unexpected http status %d, body: %s", httpResp.StatusCode, string(body))
	}

	// 读取响应（限制大小防止 OOM）
	respBody, err := io.ReadAll(io.LimitReader(httpResp.Body, maxResponseBodySize))
	if err != nil {
		return nil, fmt.Errorf("ecerp: read response error: %w", err)
	}

	if len(respBody) == 0 {
		return nil, errors.New("ecerp: empty response body")
	}

	// 调试日志
	if c.config.Logger != nil {
		c.config.Logger.Debug("ecerp: response",
			slog.String("method", method),
			slog.Int("status", httpResp.StatusCode),
			slog.Int("body_size", len(respBody)),
		)
	}

	// 解析响应
	var apiResp Response
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("ecerp: unmarshal response error: %w, body: %s", err, truncateStr(string(respBody), 500))
	}

	return &apiResp, nil
}

// Close 关闭客户端，释放底层连接
func (c *Client) Close() {
	if c.closed.CompareAndSwap(false, true) {
		if transport, ok := c.config.HTTPClient.Transport.(*http.Transport); ok {
			transport.CloseIdleConnections()
		}
	}
}

// generateNonceStr 生成加密安全的随机字符串
func generateNonceStr() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("crypto/rand failed: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// truncateStr 截断字符串，用于日志输出
func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...(truncated)"
}
