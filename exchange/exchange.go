// Package exchange 提供易仓ERP汇率相关API的封装
package exchange

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 汇率服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建汇率服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// CurrencyRate 币种汇率
type CurrencyRate struct {
	CurrencyCode string  `json:"currency_code"`
	CurrencyName string  `json:"currency_name"`
	Rate         float64 `json:"rate"`
	BaseCurrency string  `json:"base_currency"`
	UpdateTime   string  `json:"update_time"`
}

// AddCurrencyRateRequest 新增/编辑汇率请求
type AddCurrencyRateRequest struct {
	CurrencyCode string  `json:"currency_code"`
	Rate         float64 `json:"rate"`
	BaseCurrency string  `json:"base_currency,omitempty"`
}

// GetCurrencyList 获取汇率列表
func (s *Service) GetCurrencyList(ctx context.Context) ([]CurrencyRate, error) {
	var result []CurrencyRate
	err := s.C.Do(ctx, "getCurrencyList", nil, &result)
	return result, err
}

// EditCurrency 编辑汇率
func (s *Service) EditCurrency(ctx context.Context, req *AddCurrencyRateRequest) error {
	return s.C.Do(ctx, "editCurrency", req, nil)
}
