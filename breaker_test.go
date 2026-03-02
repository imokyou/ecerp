package ecerp

import (
	"sync"
	"testing"
	"time"
)

// ─────────────────────────────────────────────
// 基础状态转换
// ─────────────────────────────────────────────

func newTestBreaker(threshold int32, timeout time.Duration) *CircuitBreaker {
	return newCircuitBreaker(BreakerConfig{
		Threshold: threshold,
		Timeout:   timeout,
	}, nil)
}

func TestBreaker_InitialStateClosed(t *testing.T) {
	cb := newTestBreaker(3, time.Second)
	if cb.State() != "closed" {
		t.Errorf("期望初始状态 closed, 收到 %s", cb.State())
	}
	if !cb.Allow() {
		t.Error("Closed 状态应允许请求")
	}
}

func TestBreaker_TripsAfterThreshold(t *testing.T) {
	cb := newTestBreaker(3, time.Second)

	// 前 2 次失败不触发熔断
	cb.RecordFailure()
	if cb.State() != "closed" {
		t.Errorf("2次失败后应仍为 closed, 收到 %s", cb.State())
	}
	cb.RecordFailure()

	// 第 3 次失败触发熔断
	cb.RecordFailure()
	if cb.State() != "open" {
		t.Errorf("3次连续失败后期望 open, 收到 %s", cb.State())
	}
}

func TestBreaker_OpenRejectsRequests(t *testing.T) {
	cb := newTestBreaker(1, time.Minute) // 长冷却时间确保不会进入 half-open
	cb.RecordFailure()

	if cb.State() != "open" {
		t.Fatalf("期望 open, 收到 %s", cb.State())
	}
	if cb.Allow() {
		t.Error("Open 状态应拒绝请求")
	}
}

// ─────────────────────────────────────────────
// Half-Open 状态转换
// ─────────────────────────────────────────────

func TestBreaker_TransitionsToHalfOpenAfterTimeout(t *testing.T) {
	cb := newTestBreaker(1, time.Millisecond) // 极短冷却时间
	cb.RecordFailure()

	if cb.State() != "open" {
		t.Fatalf("期望 open, 收到 %s", cb.State())
	}

	// 等待冷却期结束
	time.Sleep(5 * time.Millisecond)

	// Allow() 应将状态切换为 half_open 并放行请求
	if !cb.Allow() {
		t.Error("冷却期后应放行试探请求")
	}
	if cb.State() != "half_open" {
		t.Errorf("冷却期后期望 half_open, 收到 %s", cb.State())
	}
}

func TestBreaker_HalfOpen_SuccessCloses(t *testing.T) {
	cb := newTestBreaker(1, time.Millisecond)
	cb.RecordFailure()
	time.Sleep(5 * time.Millisecond)
	cb.Allow() // 触发 open → half_open

	cb.RecordSuccess()

	if cb.State() != "closed" {
		t.Errorf("Half-Open 成功后期望 closed, 收到 %s", cb.State())
	}
	if !cb.Allow() {
		t.Error("Closed 状态应允许请求")
	}
}

func TestBreaker_HalfOpen_FailureReopens(t *testing.T) {
	cb := newTestBreaker(1, time.Millisecond)
	cb.RecordFailure()
	time.Sleep(5 * time.Millisecond)
	cb.Allow() // 触发 open → half_open

	cb.RecordFailure() // 试探失败，重新打开

	if cb.State() != "open" {
		t.Errorf("Half-Open 失败后期望重新 open, 收到 %s", cb.State())
	}
	if cb.Allow() {
		t.Error("重新 Open 后应拒绝请求")
	}
}

// ─────────────────────────────────────────────
// 成功重置失败计数
// ─────────────────────────────────────────────

func TestBreaker_SuccessResetsFailures(t *testing.T) {
	cb := newTestBreaker(3, time.Second)

	// 连续失败 2 次
	cb.RecordFailure()
	cb.RecordFailure()

	// 成功一次，重置计数
	cb.RecordSuccess()

	// 再失败 2 次，不应触发熔断（因为已重置）
	cb.RecordFailure()
	cb.RecordFailure()

	if cb.State() != "closed" {
		t.Errorf("成功重置后2次失败不应触发熔断, 收到 %s", cb.State())
	}
}

// ─────────────────────────────────────────────
// 禁用熔断器（Threshold = 0）
// ─────────────────────────────────────────────

func TestBreaker_Disabled_ReturnsNil(t *testing.T) {
	cb := newCircuitBreaker(BreakerConfig{Threshold: 0}, nil)
	if cb != nil {
		t.Error("Threshold=0 应返回 nil 熔断器")
	}
}

// ─────────────────────────────────────────────
// 状态变化回调（onState）
// ─────────────────────────────────────────────

func TestBreaker_StateChangeCallback(t *testing.T) {
	states := []string{}
	cb := newCircuitBreaker(BreakerConfig{Threshold: 1, Timeout: time.Millisecond}, func(s string) {
		states = append(states, s)
	})

	cb.RecordFailure() // → open
	time.Sleep(5 * time.Millisecond)
	cb.Allow()         // → half_open
	cb.RecordSuccess() // → closed

	if len(states) < 3 {
		t.Fatalf("期望至少3次状态变化回调, 收到 %d 次: %v", len(states), states)
	}
	if states[0] != "open" {
		t.Errorf("第1次变化期望 open, 收到 %s", states[0])
	}
	if states[1] != "half_open" {
		t.Errorf("第2次变化期望 half_open, 收到 %s", states[1])
	}
	if states[2] != "closed" {
		t.Errorf("第3次变化期望 closed, 收到 %s", states[2])
	}
}

// ─────────────────────────────────────────────
// 并发安全（-race 检测）
// ─────────────────────────────────────────────

func TestBreaker_ConcurrentSafe(t *testing.T) {
	cb := newTestBreaker(10, 10*time.Millisecond)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.Allow()
			cb.RecordFailure()
			cb.RecordSuccess()
			_ = cb.State()
		}()
	}
	wg.Wait()
	// 只要没有 race detector 报错，即为通过
}
