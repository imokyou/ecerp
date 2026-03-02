package ecerp

import (
	"context"
	"errors"
	"testing"
	"time"
)

// ─────────────────────────────────────────────
// isRetryable
// ─────────────────────────────────────────────

func TestIsRetryable(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil 不重试", nil, false},
		{"context.Canceled 不重试", context.Canceled, false},
		{"context.DeadlineExceeded 可重试", context.DeadlineExceeded, true},
		{"APIError 5xx 可重试", &APIError{Code: 500, Message: "server error"}, true},
		{"APIError 429 可重试", &APIError{Code: 429, Message: "rate limit"}, true},
		{"APIError 400 不重试", &APIError{Code: 400, Message: "bad request"}, false},
		{"APIError 401 不重试", &APIError{Code: 401, Message: "unauthorized"}, false},
		{"APIError 404 不重试", &APIError{Code: 404, Message: "not found"}, false},
		{"普通错误不重试", errors.New("some network error"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRetryable(tt.err)
			if got != tt.want {
				t.Errorf("isRetryable(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

// ─────────────────────────────────────────────
// RetryConfig.calcDelay
// ─────────────────────────────────────────────

func TestCalcDelay_InBounds(t *testing.T) {
	cfg := RetryConfig{
		BaseDelay:  100 * time.Millisecond,
		MaxDelay:   5 * time.Second,
		Multiplier: 2.0,
	}

	for attempt := 0; attempt < 10; attempt++ {
		d := cfg.calcDelay(attempt)
		// 加了 ±25% 抖动，实际延迟应在 [BaseDelay*0.75, MaxDelay*1.25] 范围内
		lower := time.Duration(float64(cfg.BaseDelay) * 0.75)
		upper := time.Duration(float64(cfg.MaxDelay) * 1.25)
		if d < lower || d > upper {
			t.Errorf("attempt=%d: calcDelay=%v, 期望在 [%v, %v] 内", attempt, d, lower, upper)
		}
	}
}

// ─────────────────────────────────────────────
// doWithRetry
// ─────────────────────────────────────────────

func TestDoWithRetry_SuccessOnFirstAttempt(t *testing.T) {
	calls := 0
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}
	err := doWithRetry(context.Background(), cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("期望成功, 收到: %v", err)
	}
	if calls != 1 {
		t.Errorf("期望调用1次, 实际 %d 次", calls)
	}
}

func TestDoWithRetry_RetriesOnServerError(t *testing.T) {
	calls := 0
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}
	serverErr := &APIError{Code: 500, Message: "internal error"}

	err := doWithRetry(context.Background(), cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		if calls < 3 {
			return serverErr
		}
		return nil // 第3次成功
	})
	if err != nil {
		t.Fatalf("期望第3次成功, 收到: %v", err)
	}
	if calls != 3 {
		t.Errorf("期望调用3次, 实际 %d 次", calls)
	}
}

func TestDoWithRetry_ExhaustsMaxAttempts(t *testing.T) {
	calls := 0
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}
	serverErr := &APIError{Code: 503, Message: "unavailable"}

	err := doWithRetry(context.Background(), cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		return serverErr
	})
	if err == nil {
		t.Fatal("期望耗尽重试后返回错误")
	}
	if calls != 3 {
		t.Errorf("期望调用3次, 实际 %d 次", calls)
	}
}

func TestDoWithRetry_NoRetryOnClientError(t *testing.T) {
	calls := 0
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}
	clientErr := &APIError{Code: 400, Message: "bad request"}

	err := doWithRetry(context.Background(), cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		return clientErr
	})
	if err == nil {
		t.Fatal("期望返回错误")
	}
	if calls != 1 {
		t.Errorf("客户端错误不应重试, 期望调用1次, 实际 %d 次", calls)
	}
}

func TestDoWithRetry_StopsOnContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	calls := 0
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}

	err := doWithRetry(ctx, cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		return &APIError{Code: 500, Message: "error"}
	})
	// context 已取消，doWithRetry 的等待会立即返回 ctx.Err()
	// 第一次调用会发生（不等待），之后等待时发现 ctx 取消
	if err == nil {
		t.Fatal("期望 context 取消后返回错误")
	}
}

func TestDoWithRetry_MaxAttemptsOne_NoRetry(t *testing.T) {
	calls := 0
	cfg := RetryConfig{MaxAttempts: 1, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}

	err := doWithRetry(context.Background(), cfg, nil, NoopMetrics{}, "test", func() error {
		calls++
		return &APIError{Code: 500, Message: "error"}
	})
	if err == nil {
		t.Fatal("期望返回错误")
	}
	if calls != 1 {
		t.Errorf("MaxAttempts=1 时不应重试, 期望调用1次, 实际 %d 次", calls)
	}
}

// TestDoWithRetry_RecordRetryCallback 验证重试时 Metrics.RecordRetry 被正确调用
func TestDoWithRetry_RecordRetryCallback(t *testing.T) {
	cfg := RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Multiplier: 1}
	serverErr := &APIError{Code: 500, Message: "error"}

	retryCalls := 0
	mockM := &mockMetrics{
		onRetry: func(_ string, _ int) { retryCalls++ },
	}

	_ = doWithRetry(context.Background(), cfg, nil, mockM, "testMethod", func() error {
		return serverErr
	})

	// MaxAttempts=3 要兤1首+2重试，应调用 RecordRetry 2 次
	if retryCalls != 2 {
		t.Errorf("期望 RecordRetry 被调用2次, 实际 %d 次", retryCalls)
	}
}

// mockMetrics 用于测试的指标实现
type mockMetrics struct {
	onRequest func(string, int, time.Duration, error)
	onRetry   func(string, int)
	onBreaker func(string)
}

func (m *mockMetrics) RecordRequest(method string, code int, d time.Duration, err error) {
	if m.onRequest != nil {
		m.onRequest(method, code, d, err)
	}
}
func (m *mockMetrics) RecordRetry(method string, attempt int) {
	if m.onRetry != nil {
		m.onRetry(method, attempt)
	}
}
func (m *mockMetrics) RecordCircuitBreaker(state string) {
	if m.onBreaker != nil {
		m.onBreaker(state)
	}
}
