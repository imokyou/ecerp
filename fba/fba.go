// Package fba 提供易仓ERP FBA相关API的封装
package fba

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service FBA服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建FBA服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// Return FBA退件
type Return struct {
	ReturnID      int    `json:"return_id"`
	OrderCode     string `json:"order_code"`
	FNSKU         string `json:"fnsku"`
	ASIN          string `json:"asin"`
	SKU           string `json:"sku"`
	Quantity      int    `json:"quantity"`
	Reason        string `json:"reason"`
	Status        string `json:"status"`
	FulfillCenter string `json:"fulfill_center"`
	CreateTime    string `json:"create_time"`
}

// Reimbursement FBA赔偿
type Reimbursement struct {
	ReimbursementID string  `json:"reimbursement_id"`
	CaseID          string  `json:"case_id"`
	OrderCode       string  `json:"order_code"`
	FNSKU           string  `json:"fnsku"`
	ASIN            string  `json:"asin"`
	SKU             string  `json:"sku"`
	Quantity        int     `json:"quantity"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Reason          string  `json:"reason"`
	CreateTime      string  `json:"create_time"`
}

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// GetFbaReturnRequest 查询FBA退件请求
type GetFbaReturnRequest struct {
	PageRequest
	OrderCode     string `json:"order_code,omitempty"`
	SKU           string `json:"sku,omitempty"`
	CreateTimeFor string `json:"create_time_for,omitempty"`
	CreateTimeTo  string `json:"create_time_to,omitempty"`
}

// GetFbaReimbursementRequest 查询FBA赔偿请求
type GetFbaReimbursementRequest struct {
	PageRequest
	ReimbursementID string `json:"reimbursement_id,omitempty"`
	CaseID          string `json:"case_id,omitempty"`
	CreateTimeFor   string `json:"create_time_for,omitempty"`
	CreateTimeTo    string `json:"create_time_to,omitempty"`
}

// GetFbaReturn 查询FBA退件
func (s *Service) GetFbaReturn(ctx context.Context, req *GetFbaReturnRequest) ([]Return, error) {
	var result []Return
	err := s.C.Do(ctx, "getFbaReturn", req, &result)
	return result, err
}

// GetFbaReimbursement 查询FBA赔偿
func (s *Service) GetFbaReimbursement(ctx context.Context, req *GetFbaReimbursementRequest) ([]Reimbursement, error) {
	var result []Reimbursement
	err := s.C.Do(ctx, "getFbaReimbursement", req, &result)
	return result, err
}
