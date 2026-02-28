// Package packing 提供易仓ERP组包拣货相关API的封装
package packing

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 组包拣货服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建组包拣货服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// PackageOrder 组包订单
type PackageOrder struct {
	OrderCode      string `json:"order_code"`
	PackageCode    string `json:"package_code"`
	WarehouseCode  string `json:"warehouse_code"`
	ShippingMethod string `json:"shipping_method"`
	TrackingNumber string `json:"tracking_number"`
	Weight         string `json:"weight"`
	Status         int    `json:"status"`
	CreateTime     string `json:"create_time"`
}

// PackageOrderLabel 组包面单
type PackageOrderLabel struct {
	OrderCode string `json:"order_code"`
	LabelURL  string `json:"label_url"`
	LabelType string `json:"label_type"`
}

// FastPackageOrderRequest 快捷组包揽件请求
type FastPackageOrderRequest struct {
	OrderCode      string `json:"order_code"`
	WarehouseCode  string `json:"warehouse_code,omitempty"`
	ShippingMethod string `json:"shipping_method,omitempty"`
	Weight         string `json:"weight,omitempty"`
}

// GetPackageOrderConditionRequest 获取组包条件请求
type GetPackageOrderConditionRequest struct {
	OrderCode string `json:"order_code"`
}

// GetPackageOrderLabelRequest 获取组包面单请求
type GetPackageOrderLabelRequest struct {
	OrderCode string `json:"order_code"`
}

// CreatePackageOrderNew 快捷组包揽件
func (s *Service) CreatePackageOrderNew(ctx context.Context, req *FastPackageOrderRequest) error {
	return s.C.Do(ctx, "createPackageOrderNew", req, nil)
}

// GetPackageOrderCondition 获取组包条件
func (s *Service) GetPackageOrderCondition(ctx context.Context, req *GetPackageOrderConditionRequest) (*PackageOrder, error) {
	var result PackageOrder
	err := s.C.Do(ctx, "getPackageOrderCondition", req, &result)
	return &result, err
}

// GetPackageOrderLabel 获取组包面单
func (s *Service) GetPackageOrderLabel(ctx context.Context, req *GetPackageOrderLabelRequest) (*PackageOrderLabel, error) {
	var result PackageOrderLabel
	err := s.C.Do(ctx, "getPackageOrderLabel", req, &result)
	return &result, err
}
