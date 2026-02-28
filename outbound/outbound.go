// Package outbound 提供易仓ERP出库单相关API的封装
package outbound

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 出库单服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建出库单服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// CeiveOrder 出库单
type CeiveOrder struct {
	CeiveCode      string           `json:"ceive_code"`
	WarehouseCode  string           `json:"warehouse_code"`
	ShippingMethod string           `json:"shipping_method"`
	TrackingNumber string           `json:"tracking_number"`
	ActionType     string           `json:"action_type"`
	Status         int              `json:"status"`
	Note           string           `json:"note"`
	CreateTime     string           `json:"create_time"`
	ShipTime       string           `json:"ship_time"`
	Items          []CeiveOrderItem `json:"items,omitempty"`
}

// CeiveOrderItem 出库单明细
type CeiveOrderItem struct {
	ProductSKU  string `json:"product_sku"`
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	LcCode      string `json:"lc_code"`
}

// DeliveryDetail 出库明细
type DeliveryDetail struct {
	CeiveCode     string `json:"ceive_code"`
	ProductSKU    string `json:"product_sku"`
	ProductName   string `json:"product_name"`
	Qty           int    `json:"qty"`
	WarehouseCode string `json:"warehouse_code"`
	CreateTime    string `json:"create_time"`
}

// ShippingBoxNumber 渠道框号
type ShippingBoxNumber struct {
	OrderCode string `json:"order_code"`
	BoxNumber string `json:"box_number"`
	Channel   string `json:"channel"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// SaveCeiveRequest 创建/编辑出库单请求
type SaveCeiveRequest struct {
	CeiveCode     string           `json:"ceive_code,omitempty"`
	WarehouseCode string           `json:"warehouse_code"`
	ActionType    string           `json:"action_type"`
	Note          string           `json:"note,omitempty"`
	CeiveProduct  []CeiveOrderItem `json:"ceive_product"`
}

// GetCeiveUseListRequest 获取出库单列表请求
type GetCeiveUseListRequest struct {
	PageRequest
	CeiveCode     string `json:"ceive_code,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
	Status        int    `json:"status,omitempty"`
	CreateTimeFor string `json:"create_time_for,omitempty"`
	CreateTimeTo  string `json:"create_time_to,omitempty"`
}

// GetShippingBoxNumberRequest 获取渠道框号请求
type GetShippingBoxNumberRequest struct {
	OrderCode string `json:"order_code"`
}

// GetDeliveryDetailListRequest 出库明细请求
type GetDeliveryDetailListRequest struct {
	PageRequest
	CeiveCode     string `json:"ceive_code,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
}

// NewProductStorageRequest 新产品入库请求
type NewProductStorageRequest struct {
	WarehouseCode string           `json:"warehouse_code"`
	Items         []CeiveOrderItem `json:"items"`
}

// DefectiveGoodsShelvesRequest 次品上架请求
type DefectiveGoodsShelvesRequest struct {
	WarehouseCode string           `json:"warehouse_code"`
	Items         []CeiveOrderItem `json:"items"`
}

// ════════════════════════════════════════════
// 服务方法 (6个接口)
// ════════════════════════════════════════════

// SaveCeive 创建/编辑出库单
func (s *Service) SaveCeive(ctx context.Context, req *SaveCeiveRequest) error {
	return s.C.Do(ctx, "saveCeive", req, nil)
}

// GetCeiveUseList 获取出库单列表
func (s *Service) GetCeiveUseList(ctx context.Context, req *GetCeiveUseListRequest) ([]CeiveOrder, error) {
	var result []CeiveOrder
	err := s.C.Do(ctx, "getCeiveUseList", req, &result)
	return result, err
}

// GetShippingBoxNumber 获取出库订单的渠道框号
func (s *Service) GetShippingBoxNumber(ctx context.Context, req *GetShippingBoxNumberRequest) ([]ShippingBoxNumber, error) {
	var result []ShippingBoxNumber
	err := s.C.Do(ctx, "getShippingBoxNumber", req, &result)
	return result, err
}

// GetDeliveryDetailList 出库明细
func (s *Service) GetDeliveryDetailList(ctx context.Context, req *GetDeliveryDetailListRequest) ([]DeliveryDetail, error) {
	var result []DeliveryDetail
	err := s.C.Do(ctx, "getDeliveryDetailList", req, &result)
	return result, err
}

// NewProductStorage 新产品入库
func (s *Service) NewProductStorage(ctx context.Context, req *NewProductStorageRequest) error {
	return s.C.Do(ctx, "newProductStorage", req, nil)
}

// DefectiveGoodsShelves 次品上架
func (s *Service) DefectiveGoodsShelves(ctx context.Context, req *DefectiveGoodsShelvesRequest) error {
	return s.C.Do(ctx, "defectiveGoodsShelves", req, nil)
}
