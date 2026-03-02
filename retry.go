package ecerp

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"math/rand/v2"
	"time"
)

// RetryConfig 重试配置
//
// 采用指数退避（Exponential Backoff）加随机抖动（Jitter）策略，
// 防止多客户端同时重试导致的惊群效应（Thundering Herd）。
type RetryConfig struct {
	// MaxAttempts 最大尝试次数（含首次），默认值 3。
	// 设为 1 表示不重试（只尝试一次）。
	MaxAttempts int

	// BaseDelay 首次重试前的基础等待时间，默认值 500ms。
	BaseDelay time.Duration

	// MaxDelay 单次等待时间上限，默认值 30s。
	MaxDelay time.Duration

	// Multiplier 退避乘数，默认值 2.0。
	Multiplier float64
}

// defaultRetryConfig 返回默认重试配置
func defaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts: 3,
		BaseDelay:   500 * time.Millisecond,
		MaxDelay:    30 * time.Second,
		Multiplier:  2.0,
	}
}

// calcDelay 计算第 attempt 次重试（从0开始）的等待时间（含抖动）
func (r RetryConfig) calcDelay(attempt int) time.Duration {
	delay := float64(r.BaseDelay) * math.Pow(r.Multiplier, float64(attempt))
	if delay > float64(r.MaxDelay) {
		delay = float64(r.MaxDelay)
	}
	// 加入 ±25% 随机抖动
	jitter := delay * 0.25 * (rand.Float64()*2 - 1)
	return time.Duration(delay + jitter)
}

// isRetryable 判断错误是否可重试
//
// 可重试场景：
//   - 服务端错误（HTTP 5xx / API code >= 500）
//   - 频率限制（HTTP 429 / APIError 429）
//   - context 超时（仅区别于用户主动 Cancel）
//
// 不可重试场景：
//   - 用户主动取消（context.Canceled）
//   - 客户端参数错误（4xx, 排除429）
//   - 网络连接被拒绝（对端未启动，无需重试）
func isRetryable(err error) bool {
	if err == nil {
		return false
	}
	// 用户主动取消不重试
	if errors.Is(err, context.Canceled) {
		return false
	}
	// API 业务错误按 code 判断
	if apiErr, ok := IsAPIError(err); ok {
		return apiErr.IsRateLimitError() || apiErr.IsServerError()
	}
	// 网络超时错误可重试（context.DeadlineExceeded 由上层 context 控制）
	return errors.Is(err, context.DeadlineExceeded)
}

// doWithRetry 在指定重试配置下执行 fn，自动处理退避、日志和 Metrics 上报
func doWithRetry(ctx context.Context, cfg RetryConfig, logger *slog.Logger, metrics Metrics, method string, fn func() error) error {
	var lastErr error
	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		// 若非首次尝试则等待退避时间
		if attempt > 0 {
			delay := cfg.calcDelay(attempt - 1)
			if logger != nil {
				logger.Warn("ecerp: retrying request",
					slog.String("method", method),
					slog.Int("attempt", attempt+1),
					slog.Int("max_attempts", cfg.MaxAttempts),
					slog.Duration("delay", delay),
					slog.String("last_error", lastErr.Error()),
				)
			}
			// 上报重试指标
			if metrics != nil {
				metrics.RecordRetry(method, attempt)
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}
		}

		lastErr = fn()
		if lastErr == nil {
			return nil
		}
		if !isRetryable(lastErr) {
			return lastErr
		}
	}
	return lastErr
}
