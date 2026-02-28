package ecerp

import (
	"errors"
	"fmt"
)

// 常见错误码
const (
	CodeSuccess        = 200
	CodeBadRequest     = 400
	CodeUnauthorized   = 401
	CodeForbidden      = 403
	CodeNotFound       = 404
	CodeTooManyRequest = 429
	CodeServerError    = 500
)

// APIError 表示易仓ERP API返回的业务错误
type APIError struct {
	// Code 错误码
	Code int `json:"code"`

	// Message 错误描述
	Message string `json:"message"`
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("ecerp: api error code=%d message=%s", e.Code, e.Message)
}

// IsNotFound 判断是否为资源未找到错误
func (e *APIError) IsNotFound() bool {
	return e.Code == CodeNotFound
}

// IsAuthError 判断是否为认证错误
func (e *APIError) IsAuthError() bool {
	return e.Code == CodeUnauthorized || e.Code == CodeForbidden
}

// IsRateLimitError 判断是否为频率限制错误
func (e *APIError) IsRateLimitError() bool {
	return e.Code == CodeTooManyRequest
}

// IsServerError 判断是否为服务端错误
func (e *APIError) IsServerError() bool {
	return e.Code >= 500
}

// IsAPIError 从 error 链中提取 APIError。
// 支持 errors.As 解包 (如 fmt.Errorf %w 包装的错误)。
func IsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}
