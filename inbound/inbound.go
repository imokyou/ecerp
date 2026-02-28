// Package inbound 提供易仓ERP入库单相关API的封装
package inbound

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 入库单服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建入库单服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// ReceivingOrder 入库单
type ReceivingOrder struct {
	ReceivingCode string               `json:"receiving_code"`
	WarehouseCode string               `json:"warehouse_code"`
	PoCode        string               `json:"po_code"`
	ReceivingType int                  `json:"receiving_type"`
	Status        int                  `json:"status"`
	Note          string               `json:"note"`
	CreateTime    string               `json:"create_time"`
	ReceiveTime   string               `json:"receive_time"`
	Items         []ReceivingOrderItem `json:"items,omitempty"`
}

// ReceivingOrderItem 入库单明细
type ReceivingOrderItem struct {
	SKU          string  `json:"sku"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	ReceivedQty  int     `json:"received_qty"`
	QualifiedQty int     `json:"qualified_qty"`
	DefectiveQty int     `json:"defective_qty"`
	Cost         float64 `json:"cost"`
}

// PutAwayOrder 上架单
type PutAwayOrder struct {
	PutAwayCode   string `json:"put_away_code"`
	ReceivingCode string `json:"receiving_code"`
	WarehouseCode string `json:"warehouse_code"`
	PdStatus      int    `json:"pd_status"`
	PdType        int    `json:"pd_type"`
	CreateTime    string `json:"create_time"`
}

// QcOrder 质检单
type QcOrder struct {
	QcCode        string `json:"qc_code"`
	ReceivingCode string `json:"receiving_code"`
	SKU           string `json:"sku"`
	QualifiedQty  int    `json:"qualified_qty"`
	DefectiveQty  int    `json:"defective_qty"`
	Status        int    `json:"status"`
	CreateTime    string `json:"create_time"`
}

// BoxInfo 装箱信息
type BoxInfo struct {
	BoxNo  string        `json:"box_no"`
	Weight float64       `json:"weight"`
	Length float64       `json:"length"`
	Width  float64       `json:"width"`
	Height float64       `json:"height"`
	Items  []BoxInfoItem `json:"items,omitempty"`
}

// BoxInfoItem 装箱明细
type BoxInfoItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

// WarehouseCostItem 仓库成本项
type WarehouseCostItem struct {
	SKU           string  `json:"sku"`
	WarehouseCode string  `json:"warehouse_code"`
	Cost          float64 `json:"cost"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// SyncReceivingRequest 创建/编辑入库单请求
type SyncReceivingRequest struct {
	ReceivingCode string               `json:"receiving_code,omitempty"`
	WarehouseCode string               `json:"warehouse_code"`
	PoCode        string               `json:"po_code,omitempty"`
	ReceivingType int                  `json:"receiving_type,omitempty"`
	Note          string               `json:"note,omitempty"`
	Items         []ReceivingOrderItem `json:"items"`
}

// GetReceivingRequest 查询入库单信息请求
type GetReceivingRequest struct {
	PageRequest
	ReceivingCode string `json:"receiving_code,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
	PoCode        string `json:"po_code,omitempty"`
	Status        int    `json:"status,omitempty"`
	ReceivingType int    `json:"receiving_type,omitempty"`
	CreateTimeFor string `json:"create_time_for,omitempty"`
	CreateTimeTo  string `json:"create_time_to,omitempty"`
}

// GetReceivingDetailListRequest 获取入库单明细请求
type GetReceivingDetailListRequest struct {
	PageRequest
	ReceivingCode string `json:"receiving_code,omitempty"`
}

// PurchaseOrderReceivingRequest 入库单收货请求
type PurchaseOrderReceivingRequest struct {
	ReceivingCode string               `json:"receiving_code"`
	Items         []ReceivingOrderItem `json:"items"`
}

// ReceivingOrderQualityCheckRequest 入库单质检请求
type ReceivingOrderQualityCheckRequest struct {
	ReceivingCode string               `json:"receiving_code"`
	Items         []ReceivingOrderItem `json:"items"`
}

// GetPutAwayListRequest 获取上架单列表请求
type GetPutAwayListRequest struct {
	PageRequest
	PdStatus      int    `json:"pd_status,omitempty"`
	PdType        int    `json:"pd_type,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
	ReceivingType int    `json:"receiving_type,omitempty"`
}

// SearchQcOrdersRequest 查询质检单明细请求
type SearchQcOrdersRequest struct {
	PageRequest
	ReceivingCode string `json:"receiving_code,omitempty"`
	QcCode        string `json:"qc_code,omitempty"`
}

// OrderReturnWarehousingRequest 退件入库请求
type OrderReturnWarehousingRequest struct {
	OrderCode     string `json:"order_code"`
	WarehouseCode string `json:"warehouse_code"`
	Note          string `json:"note,omitempty"`
}

// PutAwayByQccodeRequest 根据质检信息操作上架请求
type PutAwayByQccodeRequest struct {
	QcCode string `json:"qc_code"`
}

// SyncConfirmReceivingRequest 强制完成入库单请求
type SyncConfirmReceivingRequest struct {
	ReceivingCode string `json:"receiving_code"`
}

// UploadReceivingBoxsInfoRequest 入库单上传装箱信息请求
type UploadReceivingBoxsInfoRequest struct {
	ReceivingCode string    `json:"receiving_code"`
	Boxes         []BoxInfo `json:"boxes"`
}

// BatchSetWarehouseCostRequest 批量更新入库单成本请求
type BatchSetWarehouseCostRequest struct {
	Items []WarehouseCostItem `json:"items"`
}

// ════════════════════════════════════════════
// 服务方法 (12个接口)
// ════════════════════════════════════════════

// SyncReceiving 创建/编辑入库单
func (s *Service) SyncReceiving(ctx context.Context, req *SyncReceivingRequest) error {
	return s.C.Do(ctx, "syncReceiving", req, nil)
}

// GetReceiving 查询入库单信息
func (s *Service) GetReceiving(ctx context.Context, req *GetReceivingRequest) ([]ReceivingOrder, error) {
	var result []ReceivingOrder
	err := s.C.Do(ctx, "getReceiving", req, &result)
	return result, err
}

// GetReceivingDetailList 获取入库单明细
func (s *Service) GetReceivingDetailList(ctx context.Context, req *GetReceivingDetailListRequest) ([]ReceivingOrderItem, error) {
	var result []ReceivingOrderItem
	err := s.C.Do(ctx, "getReceivingDetailList", req, &result)
	return result, err
}

// PurchaseOrderReceiving 入库单收货
func (s *Service) PurchaseOrderReceiving(ctx context.Context, req *PurchaseOrderReceivingRequest) error {
	return s.C.Do(ctx, "purchaseOrderReceiving", req, nil)
}

// ReceivingOrderQualityCheck 入库单质检
func (s *Service) ReceivingOrderQualityCheck(ctx context.Context, req *ReceivingOrderQualityCheckRequest) error {
	return s.C.Do(ctx, "receivingOrderQualityCheck", req, nil)
}

// GetPutAwayList 获取上架单列表
func (s *Service) GetPutAwayList(ctx context.Context, req *GetPutAwayListRequest) ([]PutAwayOrder, error) {
	var result []PutAwayOrder
	err := s.C.Do(ctx, "getPutAwayList", req, &result)
	return result, err
}

// SearchQcOrders 查询质检单明细
func (s *Service) SearchQcOrders(ctx context.Context, req *SearchQcOrdersRequest) ([]QcOrder, error) {
	var result []QcOrder
	err := s.C.Do(ctx, "searchQcOrders", req, &result)
	return result, err
}

// OrderReturnWarehousing 退件入库
func (s *Service) OrderReturnWarehousing(ctx context.Context, req *OrderReturnWarehousingRequest) error {
	return s.C.Do(ctx, "orderReturnWarehousing", req, nil)
}

// PutAwayByQccode 根据质检信息操作上架
func (s *Service) PutAwayByQccode(ctx context.Context, req *PutAwayByQccodeRequest) error {
	return s.C.Do(ctx, "putAwayByQccode", req, nil)
}

// SyncConfirmReceiving 强制完成入库单
func (s *Service) SyncConfirmReceiving(ctx context.Context, req *SyncConfirmReceivingRequest) error {
	return s.C.Do(ctx, "syncConfirmReceiving", req, nil)
}

// UploadReceivingBoxsInfo 入库单上传装箱信息
func (s *Service) UploadReceivingBoxsInfo(ctx context.Context, req *UploadReceivingBoxsInfoRequest) error {
	return s.C.Do(ctx, "uploadReceivingBoxsInfo", req, nil)
}

// BatchSetWarehouseCost 批量更新入库单成本
func (s *Service) BatchSetWarehouseCost(ctx context.Context, req *BatchSetWarehouseCostRequest) error {
	return s.C.Do(ctx, "batchSetWarehouseCost", req, nil)
}
