// Package transfer 提供易仓ERP调拨单相关API的封装
package transfer

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 调拨单服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建调拨单服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// TransferOrder 调拨单
type TransferOrder struct {
	TransferCode      string              `json:"transfer_code"`
	FromWarehouseCode string              `json:"from_warehouse_code"`
	ToWarehouseCode   string              `json:"to_warehouse_code"`
	ShippingMethod    string              `json:"shipping_method"`
	Status            int                 `json:"status"`
	Note              string              `json:"note"`
	CreateTime        string              `json:"create_time"`
	Items             []TransferOrderItem `json:"items,omitempty"`
}

// TransferOrderItem 调拨单明细
type TransferOrderItem struct {
	SKU         string `json:"sku"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// CreateTransferOrderRequest 创建调拨单请求
type CreateTransferOrderRequest struct {
	FromWarehouseCode string              `json:"from_warehouse_code"`
	ToWarehouseCode   string              `json:"to_warehouse_code"`
	ShippingMethod    string              `json:"shipping_method,omitempty"`
	Note              string              `json:"note,omitempty"`
	Items             []TransferOrderItem `json:"items"`
}

// GetTransferOrderListRequest 获取调拨单列表请求
type GetTransferOrderListRequest struct {
	PageRequest
	TransferCode      string `json:"transfer_code,omitempty"`
	FromWarehouseCode string `json:"from_warehouse_code,omitempty"`
	ToWarehouseCode   string `json:"to_warehouse_code,omitempty"`
	Status            int    `json:"status,omitempty"`
	CreateTimeFor     string `json:"create_time_for,omitempty"`
	CreateTimeTo      string `json:"create_time_to,omitempty"`
}

// EditTransferOrderRequest 编辑调拨单请求
type EditTransferOrderRequest struct {
	TransferCode      string              `json:"transfer_code"`
	FromWarehouseCode string              `json:"from_warehouse_code,omitempty"`
	ToWarehouseCode   string              `json:"to_warehouse_code,omitempty"`
	ShippingMethod    string              `json:"shipping_method,omitempty"`
	Note              string              `json:"note,omitempty"`
	Items             []TransferOrderItem `json:"items,omitempty"`
}

// CreateTransferOrder 创建调拨单
func (s *Service) CreateTransferOrder(ctx context.Context, req *CreateTransferOrderRequest) error {
	return s.C.Do(ctx, "createTranferOrder", req, nil) // 注意：文档中为 Tranfer (缺少's')
}

// GetTransferOrderList 获取调拨单列表
func (s *Service) GetTransferOrderList(ctx context.Context, req *GetTransferOrderListRequest) ([]TransferOrder, error) {
	var result []TransferOrder
	err := s.C.Do(ctx, "getTransferOrders", req, &result)
	return result, err
}

// EditTransferOrder 编辑调拨单
func (s *Service) EditTransferOrder(ctx context.Context, req *EditTransferOrderRequest) error {
	return s.C.Do(ctx, "editTranferOrder", req, nil) // 注意：文档中为 Tranfer (缺少's')
}
