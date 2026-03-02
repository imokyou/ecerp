package ecerp

import (
	"log/slog"
	"time"
)

// Metrics 可观测性指标接口
//
// 用户实现此接口并通过 WithMetrics 注入，即可接入任意监控后端
// （如 Prometheus、DataDog、OpenTelemetry 等）。
//
// 接口故意保持极简：仅3个方法，无需导入外部包。
// SDK 内部在请求完成、发生重试、熔断器状态变化时分别回调对应方法。
type Metrics interface {
	// RecordRequest 在每次 API 请求完成后调用（无论成功还是失败）。
	//   - method:     接口方法名（如 "getOrderList"）
	//   - statusCode: HTTP 状态码 或 API 业务 code。成功时为 200。
	//   - latency:    本次请求的端到端耗时（含网络+反序列化）
	//   - err:        请求错误（nil 表示成功）
	RecordRequest(method string, statusCode int, latency time.Duration, err error)

	// RecordRetry 在发生重试时调用。
	//   - method:  接口方法名
	//   - attempt: 当前第几次重试（从1开始）
	RecordRetry(method string, attempt int)

	// RecordCircuitBreaker 在熔断器状态发生变化时调用。
	//   - state: 新状态，取值 "open" | "half_open" | "closed"
	RecordCircuitBreaker(state string)
}

// NoopMetrics 空实现（默认值），所有方法均为空操作，不产生任何运行时开销。
type NoopMetrics struct{}

func (NoopMetrics) RecordRequest(_ string, _ int, _ time.Duration, _ error) {}
func (NoopMetrics) RecordRetry(_ string, _ int)                             {}
func (NoopMetrics) RecordCircuitBreaker(_ string)                           {}

// LogMetrics 基于 slog.Logger 的内置 Metrics 实现。
//
// 适用于开发调试阶段的快速接入，不需要配置任何监控系统。
// 生产环境建议实现针对 Prometheus / OpenTelemetry 的适配器。
//
// 使用方式：
//
//	client, _ := ecerp.NewClient(key, secret, svcID,
//	    ecerp.WithMetrics(ecerp.NewLogMetrics(slog.Default())),
//	)
type LogMetrics struct {
	logger *slog.Logger
}

// NewLogMetrics 创建基于 slog 的 Metrics 实现
func NewLogMetrics(logger *slog.Logger) *LogMetrics {
	if logger == nil {
		logger = slog.Default()
	}
	return &LogMetrics{logger: logger}
}

func (m *LogMetrics) RecordRequest(method string, statusCode int, latency time.Duration, err error) {
	if err != nil {
		m.logger.Warn("ecerp metrics: request failed",
			slog.String("method", method),
			slog.Int("status_code", statusCode),
			slog.Duration("latency", latency),
			slog.String("error", err.Error()),
		)
		return
	}
	m.logger.Info("ecerp metrics: request ok",
		slog.String("method", method),
		slog.Int("status_code", statusCode),
		slog.Duration("latency", latency),
	)
}

func (m *LogMetrics) RecordRetry(method string, attempt int) {
	m.logger.Warn("ecerp metrics: retrying",
		slog.String("method", method),
		slog.Int("attempt", attempt),
	)
}

func (m *LogMetrics) RecordCircuitBreaker(state string) {
	m.logger.Warn("ecerp metrics: circuit breaker state changed",
		slog.String("state", state),
	)
}
