// Package compass 提供易仓ERP数据罗盘相关API的封装
package compass

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 数据罗盘服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建数据罗盘服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// OrderStatisticsRequest 订单统计请求
type OrderStatisticsRequest struct {
	PageRequest
	AccountID int    `json:"account_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// ProductSaleRequest 产品销售请求
type ProductSaleRequest struct {
	PageRequest
	AccountID int    `json:"account_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (4个接口)
// ════════════════════════════════════════════

// GetOrderStatisticsV2 订单统计V2
func (s *Service) GetOrderStatisticsV2(ctx context.Context, req *OrderStatisticsRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getOrderStatisticsV2", req, &result)
	return result, err
}

// GetOrderStatistics 订单统计
func (s *Service) GetOrderStatistics(ctx context.Context, req *OrderStatisticsRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getOrderStatistics", req, &result)
	return result, err
}

// GetProductSaleSummary 产品维度-产品销量情况汇总
func (s *Service) GetProductSaleSummary(ctx context.Context, req *ProductSaleRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getProductSaleSummary", req, &result)
	return result, err
}

// GetProductSale 产品维度-产品销售情况
func (s *Service) GetProductSale(ctx context.Context, req *ProductSaleRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getProductSale", req, &result)
	return result, err
}
