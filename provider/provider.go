// Package provider 提供易仓ERP服务商相关API的封装
package provider

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 服务商API服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建服务商服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// PreShippingOrder 待发货订单
type PreShippingOrder struct {
	OrderCode      string            `json:"order_code"`
	ReferenceCode  string            `json:"reference_code"`
	ShippingMethod string            `json:"shipping_method"`
	WarehouseCode  string            `json:"warehouse_code"`
	CountryCode    string            `json:"country_code"`
	BuyerName      string            `json:"buyer_name"`
	BuyerPhone     string            `json:"buyer_phone"`
	Street1        string            `json:"street1"`
	Street2        string            `json:"street2"`
	City           string            `json:"city"`
	Province       string            `json:"province"`
	Zipcode        string            `json:"zipcode"`
	Weight         float64           `json:"weight"`
	Status         int               `json:"status"`
	CreateTime     string            `json:"create_time"`
	Items          []PreShippingItem `json:"items,omitempty"`
}

// PreShippingItem 订单商品明细
type PreShippingItem struct {
	SKU         string  `json:"sku"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Weight      float64 `json:"weight"`
}

// WarehouseInfo 仓库信息（待下线）
type WarehouseInfo struct {
	WarehouseID   int    `json:"warehouse_id"`
	WarehouseCode string `json:"warehouse_code"`
	WarehouseName string `json:"warehouse_name"`
	Country       string `json:"country"`
	Address       string `json:"address"`
}

// InventoryInfo 库存信息（待下线）
type InventoryInfo struct {
	SKU           string `json:"sku"`
	ProductName   string `json:"product_name"`
	WarehouseCode string `json:"warehouse_code"`
	AvailableQty  int    `json:"available_qty"`
	TotalQty      int    `json:"total_qty"`
}

// OrderLabel 跟踪号及标签
type OrderLabel struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	LabelURL       string `json:"label_url"`
}

// OrderStatus 订单状态（待下线）
type OrderStatus struct {
	OrderCode string `json:"order_code"`
	Status    int    `json:"status"`
	StatusMsg string `json:"status_msg"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// GetPreShippingOrderRequest 获取订单请求
type GetPreShippingOrderRequest struct {
	PageRequest
	OrderCode     string `json:"order_code,omitempty"`
	Status        int    `json:"status,omitempty"`
	CreateTimeFor string `json:"create_time_for,omitempty"`
	CreateTimeTo  string `json:"create_time_to,omitempty"`
}

// UploadTrackingNoRequest 上传跟踪号及标签请求
type UploadTrackingNoRequest struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	ShippingMethod string `json:"shipping_method,omitempty"`
	LabelURL       string `json:"label_url,omitempty"`
	Weight         string `json:"weight,omitempty"`
	ShipFee        string `json:"ship_fee,omitempty"`
}

// OrderStatusModifyRequest 订单出库/取消请求
type OrderStatusModifyRequest struct {
	OrderCode string `json:"order_code"`
	Status    int    `json:"status"`
	Note      string `json:"note,omitempty"`
}

// UpdateOrdersFeeRequest 更新订单尾程费用请求
type UpdateOrdersFeeRequest struct {
	OrderCode string  `json:"order_code"`
	Fee       float64 `json:"fee"`
	Currency  string  `json:"currency,omitempty"`
}

// GetWarehouseInfoRequest 获取仓库信息请求（待下线）
type GetWarehouseInfoRequest struct {
	WarehouseCodeList []string `json:"warehouse_code_list,omitempty"`
	WarehouseDescLike string   `json:"warehouse_desc_like,omitempty"`
}

// GetInventoryRequest 查询库存请求（待下线）
type GetInventoryRequest struct {
	PageRequest
	WarehouseCode string `json:"warehouse_code,omitempty"`
	SKU           string `json:"sku,omitempty"`
}

// PrintOrderLabelRequest 获取跟踪号及标签请求（待下线）
type PrintOrderLabelRequest struct {
	OrderCode string `json:"order_code"`
}

// GetOrderStatusRequest 获取订单状态请求（待下线）
type GetOrderStatusRequest struct {
	OrderCodes []string `json:"order_codes"`
}

// ════════════════════════════════════════════
// 服务方法 (8个接口)
// ════════════════════════════════════════════

// GetPreShippingOrder 获取订单
func (s *Service) GetPreShippingOrder(ctx context.Context, req *GetPreShippingOrderRequest) ([]PreShippingOrder, error) {
	var result []PreShippingOrder
	err := s.C.Do(ctx, "getPreShippingOrder", req, &result)
	return result, err
}

// UploadTrackingNo 上传跟踪号及标签
func (s *Service) UploadTrackingNo(ctx context.Context, req *UploadTrackingNoRequest) error {
	return s.C.Do(ctx, "uploadTrackingNo", req, nil)
}

// OrderStatusModify 订单出库/取消
func (s *Service) OrderStatusModify(ctx context.Context, req *OrderStatusModifyRequest) error {
	return s.C.Do(ctx, "orderStatusModify", req, nil)
}

// UpdateOrdersFee 更新订单尾程费用
func (s *Service) UpdateOrdersFee(ctx context.Context, req *UpdateOrdersFeeRequest) error {
	return s.C.Do(ctx, "updateOrdersFee", req, nil)
}

// GetWarehouseInfo 获取仓库信息（待下线）
func (s *Service) GetWarehouseInfo(ctx context.Context, req *GetWarehouseInfoRequest) ([]WarehouseInfo, error) {
	var result []WarehouseInfo
	err := s.C.Do(ctx, "getWarehouseInfo", req, &result)
	return result, err
}

// GetInventory 查询库存（待下线）
func (s *Service) GetInventory(ctx context.Context, req *GetInventoryRequest) ([]InventoryInfo, error) {
	var result []InventoryInfo
	err := s.C.Do(ctx, "getInventory", req, &result)
	return result, err
}

// PrintOrderLabel 获取跟踪号及标签（待下线）
func (s *Service) PrintOrderLabel(ctx context.Context, req *PrintOrderLabelRequest) (*OrderLabel, error) {
	var result OrderLabel
	err := s.C.Do(ctx, "printOrderLabel", req, &result)
	return &result, err
}

// GetOrderStatus 获取订单状态（待下线）
func (s *Service) GetOrderStatus(ctx context.Context, req *GetOrderStatusRequest) ([]OrderStatus, error) {
	var result []OrderStatus
	err := s.C.Do(ctx, "getOrderStatus", req, &result)
	return result, err
}
