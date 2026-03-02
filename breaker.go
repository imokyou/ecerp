package ecerp

import (
	"errors"
	"sync/atomic"
	"time"
)

// ErrCircuitOpen 熔断器开路错误：当前请求被熔断器拒绝，下游服务正在保护冷却中。
var ErrCircuitOpen = errors.New("ecerp: circuit breaker is open, request rejected")

// circuitState 熔断器状态枚举
type circuitState int32

const (
	stateClosed   circuitState = iota // 正常，放行所有请求
	stateOpen                         // 熔断，拒绝所有请求
	stateHalfOpen                     // 试探，放行单个请求
)

// BreakerConfig 熔断器配置
type BreakerConfig struct {
	// Threshold 触发熔断的连续失败次数，默认 5。
	// 设为 0 表示禁用熔断器。
	Threshold int32

	// Timeout 熔断后到 Half-Open 的冷却时间，默认 60s。
	Timeout time.Duration
}

// defaultBreakerConfig 返回默认熔断器配置
func defaultBreakerConfig() BreakerConfig {
	return BreakerConfig{
		Threshold: 5,
		Timeout:   60 * time.Second,
	}
}

// CircuitBreaker 三态熔断器（并发安全）
//
// 状态转换：
//
//	Closed ──(连续失败 >= Threshold)──► Open
//	Open   ──(冷却 Timeout 到期)     ──► Half-Open
//	Half-Open ──(请求成功)            ──► Closed
//	Half-Open ──(请求失败)            ──► Open（重置计时器）
type CircuitBreaker struct {
	cfg       BreakerConfig
	state     atomic.Int32       // circuitState
	failures  atomic.Int32       // 连续失败次数
	nextRetry atomic.Int64       // Open→Half-Open 的时间点（UnixNano）
	onState   func(state string) // 状态变化回调（用于 Metrics 上报）
}

// newCircuitBreaker 创建熔断器；cfg.Threshold==0 时返回 nil（禁用）
func newCircuitBreaker(cfg BreakerConfig, onState func(string)) *CircuitBreaker {
	if cfg.Threshold <= 0 {
		return nil
	}
	cb := &CircuitBreaker{cfg: cfg, onState: onState}
	cb.state.Store(int32(stateClosed))
	return cb
}

// Allow 判断当前请求是否被允许通过
//
// - Closed：始终允许
// - Open：若冷却期已过则切换到 Half-Open 并允许一个试探请求，否则拒绝
// - Half-Open：拒绝（只放行冷却后第一个请求，后续等待结果）
func (cb *CircuitBreaker) Allow() bool {
	switch circuitState(cb.state.Load()) {
	case stateClosed:
		return true
	case stateOpen:
		if time.Now().UnixNano() >= cb.nextRetry.Load() {
			// 冷却期结束，切换到 Half-Open，放行一个试探请求
			if cb.state.CompareAndSwap(int32(stateOpen), int32(stateHalfOpen)) {
				cb.notifyState("half_open")
			}
			return true
		}
		return false
	case stateHalfOpen:
		// Half-Open 状态只允许一个试探请求（已在 stateOpen→stateHalfOpen 时放行）
		return false
	default:
		return true
	}
}

// RecordSuccess 记录一次成功请求
func (cb *CircuitBreaker) RecordSuccess() {
	switch circuitState(cb.state.Load()) {
	case stateHalfOpen:
		// 试探成功，恢复关闭
		cb.failures.Store(0)
		cb.state.Store(int32(stateClosed))
		cb.notifyState("closed")
	case stateClosed:
		// 重置连续失败计数
		cb.failures.Store(0)
	}
}

// RecordFailure 记录一次失败请求
func (cb *CircuitBreaker) RecordFailure() {
	switch circuitState(cb.state.Load()) {
	case stateClosed:
		newFailures := cb.failures.Add(1)
		if newFailures >= cb.cfg.Threshold {
			cb.trip()
		}
	case stateHalfOpen:
		// 试探失败，重新打开熔断
		cb.trip()
	}
}

// trip 切换到 Open 状态并设置冷却期
func (cb *CircuitBreaker) trip() {
	cb.nextRetry.Store(time.Now().Add(cb.cfg.Timeout).UnixNano())
	cb.state.Store(int32(stateOpen))
	cb.notifyState("open")
}

// State 返回当前熔断器状态字符串（用于日志/监控）
func (cb *CircuitBreaker) State() string {
	switch circuitState(cb.state.Load()) {
	case stateClosed:
		return "closed"
	case stateOpen:
		return "open"
	case stateHalfOpen:
		return "half_open"
	default:
		return "unknown"
	}
}

func (cb *CircuitBreaker) notifyState(state string) {
	if cb.onState != nil {
		cb.onState(state)
	}
}
