package ecerp

import (
	"log/slog"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL 易仓ERP开放平台默认API地址
	DefaultBaseURL = "https://openapi-web.eccang.com/openApi/api/unity"

	// DefaultTimeout 默认超时时间
	DefaultTimeout = 30 * time.Second

	// DefaultCharset 默认字符集
	DefaultCharset = "UTF-8"

	// DefaultVersion 默认API版本
	DefaultVersion = "V1.0.0"

	// DefaultSignType 默认签名类型
	DefaultSignType = "MD5"

	// DefaultUserAgent 默认 User-Agent
	DefaultUserAgent = "ecerp-go-sdk/1.0"
)

// Config 客户端配置
type Config struct {
	// AppKey 应用Key（必填）
	AppKey string

	// AppSecret 应用密钥，用于签名（必填）
	AppSecret string

	// ServiceID 授权服务ID（必填）
	ServiceID string

	// BaseURL API基础地址
	BaseURL string

	// HTTPClient 自定义HTTP客户端。
	// 如果不设置，将自动创建一个带有连接池的 http.Client。
	HTTPClient *http.Client

	// Timeout 请求超时时间
	Timeout time.Duration

	// Charset 字符集
	Charset string

	// Version API版本
	Version string

	// SignType 签名类型 (MD5 或 AES)
	SignType string

	// UserAgent 请求 User-Agent
	UserAgent string

	// Logger 日志记录器。设为 nil 关闭日志。
	// 使用 slog.Logger 以兼容 Go 标准库。
	Logger *slog.Logger

	// Retry 重试配置。
	// 默认启用（MaxAttempts=3）。若要禁用重试，可使用 WithDisableRetry()。
	Retry RetryConfig

	// Breaker 熔断器配置。
	// 默认启用（Threshold=5，Timeout=60s）。若要禁用熔断器，可使用 WithDisableBreaker()。
	Breaker BreakerConfig

	// Metrics 可观测性指标接口。
	// 默认为 NoopMetrics（空实现，零开销）。
	// 可使用 WithMetrics() 接入 Prometheus / OTEL 等。
	Metrics Metrics
}

// Option 函数式选项
type Option func(*Config)

// WithBaseURL 设置自定义API基础地址
func WithBaseURL(url string) Option {
	return func(c *Config) {
		c.BaseURL = url
	}
}

// WithHTTPClient 设置自定义HTTP客户端
//
// 当设置此选项时，WithTimeout 将不会影响 HTTP 客户端的超时配置，
// 请在自定义 http.Client 中自行设置。
func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = client
	}
}

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithCharset 设置字符集
func WithCharset(charset string) Option {
	return func(c *Config) {
		c.Charset = charset
	}
}

// WithVersion 设置API版本
func WithVersion(version string) Option {
	return func(c *Config) {
		c.Version = version
	}
}

// WithSignType 设置签名类型
func WithSignType(signType string) Option {
	return func(c *Config) {
		c.SignType = signType
	}
}

// WithUserAgent 设置 User-Agent
func WithUserAgent(ua string) Option {
	return func(c *Config) {
		c.UserAgent = ua
	}
}

// WithLogger 设置日志记录器
func WithLogger(logger *slog.Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// WithRetry 自定义重试配置
func WithRetry(retry RetryConfig) Option {
	return func(c *Config) {
		c.Retry = retry
	}
}

// WithDisableRetry 禁用自动重试（每次请求最多尝试一次）
func WithDisableRetry() Option {
	return func(c *Config) {
		c.Retry.MaxAttempts = 1
	}
}

// WithBreaker 自定义熔断器配置
func WithBreaker(cfg BreakerConfig) Option {
	return func(c *Config) {
		c.Breaker = cfg
	}
}

// WithDisableBreaker 禁用熔断器
func WithDisableBreaker() Option {
	return func(c *Config) {
		c.Breaker.Threshold = 0
	}
}

// WithMetrics 设置可观测性指标接口
func WithMetrics(m Metrics) Option {
	return func(c *Config) {
		c.Metrics = m
	}
}

// newDefaultConfig 创建默认配置
func newDefaultConfig(appKey, appSecret, serviceID string) *Config {
	return &Config{
		AppKey:    appKey,
		AppSecret: appSecret,
		ServiceID: serviceID,
		BaseURL:   DefaultBaseURL,
		Timeout:   DefaultTimeout,
		Charset:   DefaultCharset,
		Version:   DefaultVersion,
		SignType:  DefaultSignType,
		UserAgent: DefaultUserAgent,
		Retry:     defaultRetryConfig(),
		Breaker:   defaultBreakerConfig(),
		Metrics:   NoopMetrics{},
	}
}
