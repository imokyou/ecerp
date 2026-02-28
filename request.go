package ecerp

import "encoding/json"

// Request 统一请求结构
type Request struct {
	AppKey          string `json:"app_key"`
	ServiceID       string `json:"service_id"`
	InterfaceMethod string `json:"interface_method"`
	Timestamp       string `json:"timestamp"`
	NonceStr        string `json:"nonce_str"`
	Charset         string `json:"charset"`
	Version         string `json:"version"`
	SignType        string `json:"sign_type"`
	BizContent      string `json:"biz_content"`
	Sign            string `json:"sign"`
}

// Response 统一响应结构
type Response struct {
	// Code 响应码，200 表示成功
	Code int `json:"code"`

	// Message 响应消息
	Message string `json:"message"`

	// Data 业务数据（原始 JSON）
	Data json.RawMessage `json:"data"`
}

// PageRequest 通用分页请求参数
type PageRequest struct {
	// Page 页码，从 1 开始
	Page int `json:"page,omitempty"`

	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
}

// PageResponse 通用分页响应
type PageResponse[T any] struct {
	// TotalCount 总记录数
	TotalCount int `json:"totalCount"`

	// PageSize 每页条数
	PageSize int `json:"pageSize"`

	// Page 当前页码
	Page int `json:"page"`

	// List 数据列表
	List []T `json:"list"`
}
