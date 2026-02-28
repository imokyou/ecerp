// Package inventory 提供易仓ERP库存相关API的封装
package inventory

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 库存服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建库存服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// ProductInventory 产品库存
type ProductInventory struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	AvailableQty  int    `json:"available_qty"`
	ReservedQty   int    `json:"reserved_qty"`
	TotalQty      int    `json:"total_qty"`
	DefectiveQty  int    `json:"defective_qty"`
}

// CargoRightsAdjustment 货权调整信息
type CargoRightsAdjustment struct {
	AcrID         int    `json:"acr_id"`
	AcrType       int    `json:"acr_type"`
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	AdjustQty     int    `json:"adjust_qty"`
	AcrdStatus    int    `json:"acrd_status"`
	CreateTime    string `json:"create_time"`
}

// InventoryBatch 批次库存
type InventoryBatch struct {
	BatchCode     string `json:"batch_code"`
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	Quantity      int    `json:"quantity"`
	CreateTime    string `json:"create_time"`
}

// InventoryBatchLog 库存批次日志
type InventoryBatchLog struct {
	LogID         int    `json:"log_id"`
	BatchCode     string `json:"batch_code"`
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	ChangeQty     int    `json:"change_qty"`
	ChangeType    string `json:"change_type"`
	CreateTime    string `json:"create_time"`
}

// LocationInventory 库位维度库存
type LocationInventory struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	LocationCode  string `json:"location_code"`
	Quantity      int    `json:"quantity"`
}

// InventoryStatistic 盘盈盘亏统计
type InventoryStatistic struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	ProfitQty     int    `json:"profit_qty"`
	LossQty       int    `json:"loss_qty"`
}

// FbaInventory FBA库存
type FbaInventory struct {
	SellerSKU      string `json:"seller_sku"`
	ASIN           string `json:"asin"`
	FulfillableQty int    `json:"fulfillable_qty"`
	InboundQty     int    `json:"inbound_qty"`
	ReservedQty    int    `json:"reserved_qty"`
}

// TransitBatch 在途批次
type TransitBatch struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	TransitQty    int    `json:"transit_qty"`
}

// AdjustmentInventory 批次库存调整记录
type AdjustmentInventory struct {
	AdjustID      int    `json:"adjust_id"`
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	AdjustQty     int    `json:"adjust_qty"`
	AdjustType    string `json:"adjust_type"`
	CreateTime    string `json:"create_time"`
}

// ProductLocation 产品库位
type ProductLocation struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	LocationCode  string `json:"location_code"`
	Quantity      int    `json:"quantity"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// GetAdjustingCargoRightsRequest 获取货权调整信息请求
type GetAdjustingCargoRightsRequest struct {
	PageRequest
	AcrType            string `json:"acr_type,omitempty"`
	CustomBusinessType string `json:"custom_business_type,omitempty"`
	AcrdStatus         string `json:"acrd_status,omitempty"`
}

// GetProductInventoryRequest 获取库存信息请求
type GetProductInventoryRequest struct {
	PageRequest
	ProductSKU    string `json:"product_sku,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
}

// AdjustInventoryBatchRequest 批次库存（调库存）请求
type AdjustInventoryBatchRequest struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	AdjustQty     int    `json:"adjust_qty"`
	Reason        string `json:"reason,omitempty"`
}

// MoveInventoryBatchRequest 批次库存（移货架）请求
type MoveInventoryBatchRequest struct {
	ProductSKU       string `json:"product_sku"`
	WarehouseCode    string `json:"warehouse_code"`
	FromLocationCode string `json:"from_location_code"`
	ToLocationCode   string `json:"to_location_code"`
	Quantity         int    `json:"quantity"`
}

// TakeStockForLocationRequest 库存盘点请求
type TakeStockForLocationRequest struct {
	WarehouseCode string `json:"warehouse_code"`
	LocationCode  string `json:"location_code"`
	ProductSKU    string `json:"product_sku"`
	Quantity      int    `json:"quantity"`
}

// ImportInventoryRequest 导入库存请求
type ImportInventoryRequest struct {
	Items []ImportInventoryItem `json:"items"`
}

// ImportInventoryItem 导入库存项
type ImportInventoryItem struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	Quantity      int    `json:"quantity"`
}

// SyncProductInventorySharedRequest 同步分销库存请求
type SyncProductInventorySharedRequest struct {
	ProductSKU    string `json:"product_sku"`
	WarehouseCode string `json:"warehouse_code"`
	Quantity      int    `json:"quantity"`
}

// ════════════════════════════════════════════
// 服务方法 (19个接口)
// ════════════════════════════════════════════

// --- 库存查询 ---

// GetAdjustingCargoRights 获取货权调整信息
func (s *Service) GetAdjustingCargoRights(ctx context.Context, req *GetAdjustingCargoRightsRequest) ([]CargoRightsAdjustment, error) {
	var result []CargoRightsAdjustment
	err := s.C.Do(ctx, "getAdjustingCargoRights", req, &result)
	return result, err
}

// GetProductInventoryNew 获取库存信息（新）
func (s *Service) GetProductInventoryNew(ctx context.Context, req *GetProductInventoryRequest) ([]ProductInventory, error) {
	var result []ProductInventory
	err := s.C.Do(ctx, "getProductInventoryNew", req, &result)
	return result, err
}

// GetProductInventory 获取库存信息
func (s *Service) GetProductInventory(ctx context.Context, req *GetProductInventoryRequest) ([]ProductInventory, error) {
	var result []ProductInventory
	err := s.C.Do(ctx, "getProductInventory", req, &result)
	return result, err
}

// GetProductInventoryTeam 获取团队库存信息
func (s *Service) GetProductInventoryTeam(ctx context.Context, req *GetProductInventoryRequest) ([]ProductInventory, error) {
	var result []ProductInventory
	err := s.C.Do(ctx, "getProductInventoryTeam", req, &result)
	return result, err
}

// GetProductInventoryForLocation 查询库存（库位维度显示）
func (s *Service) GetProductInventoryForLocation(ctx context.Context, req *GetProductInventoryRequest) ([]LocationInventory, error) {
	var result []LocationInventory
	err := s.C.Do(ctx, "getProductInventoryForLocation", req, &result)
	return result, err
}

// GetFbaInventory 查询FBA库存
func (s *Service) GetFbaInventory(ctx context.Context, req *PageRequest) ([]FbaInventory, error) {
	var result []FbaInventory
	err := s.C.Do(ctx, "getFbaInventory", req, &result)
	return result, err
}

// --- 批次库存 ---

// GetInventoryBatch 获取批次库存
func (s *Service) GetInventoryBatch(ctx context.Context, req *GetProductInventoryRequest) ([]InventoryBatch, error) {
	var result []InventoryBatch
	err := s.C.Do(ctx, "getInventoryBatch", req, &result)
	return result, err
}

// AdjustInventoryBatch 批次库存（调库存）
func (s *Service) AdjustInventoryBatch(ctx context.Context, req *AdjustInventoryBatchRequest) error {
	return s.C.Do(ctx, "adjustInventoryBatch", req, nil)
}

// MoveInventoryBatch 批次库存（移货架）
func (s *Service) MoveInventoryBatch(ctx context.Context, req *MoveInventoryBatchRequest) error {
	return s.C.Do(ctx, "moveInventoryBatch", req, nil)
}

// GetInventoryBatchLog 库存批次日志
func (s *Service) GetInventoryBatchLog(ctx context.Context, req *GetProductInventoryRequest) ([]InventoryBatchLog, error) {
	var result []InventoryBatchLog
	err := s.C.Do(ctx, "getInventoryBatchLog", req, &result)
	return result, err
}

// GetAdjustmentInventoryList 获取批次库存调整记录
func (s *Service) GetAdjustmentInventoryList(ctx context.Context, req *PageRequest) ([]AdjustmentInventory, error) {
	var result []AdjustmentInventory
	err := s.C.Do(ctx, "getAdjustmentInventoryList", req, &result)
	return result, err
}

// GetTransitBatchNumber 获取产品批次在途数量
func (s *Service) GetTransitBatchNumber(ctx context.Context, req *GetProductInventoryRequest) ([]TransitBatch, error) {
	var result []TransitBatch
	err := s.C.Do(ctx, "getTransitBatchNumber", req, &result)
	return result, err
}

// --- 盘点/统计 ---

// TakeStockForLocation 库存盘点（根据第三方盘点结果调整库存）
func (s *Service) TakeStockForLocation(ctx context.Context, req *TakeStockForLocationRequest) error {
	return s.C.Do(ctx, "takeStockForLocation", req, nil)
}

// InventoryStatistics 盘盈盘亏统计
func (s *Service) InventoryStatistics(ctx context.Context, req *PageRequest) ([]InventoryStatistic, error) {
	var result []InventoryStatistic
	err := s.C.Do(ctx, "inventoryStatistics", req, &result)
	return result, err
}

// --- 导入/同步 ---

// ImportInventory 导入库存
func (s *Service) ImportInventory(ctx context.Context, req *ImportInventoryRequest) error {
	return s.C.Do(ctx, "importInventory", req, nil)
}

// SyncProductInventoryShared 同步分销库存
func (s *Service) SyncProductInventoryShared(ctx context.Context, req *SyncProductInventorySharedRequest) error {
	return s.C.Do(ctx, "syncProductInventoryShared", req, nil)
}

// SyncProductInventorySharedBatch 批量同步分销库存
func (s *Service) SyncProductInventorySharedBatch(ctx context.Context, items []SyncProductInventorySharedRequest) error {
	return s.C.Do(ctx, "syncProductInventorySharedBatch", items, nil)
}

// --- 其他 ---

// GetProductInventoryMaxAge 获取产品最大库龄（分仓库）
func (s *Service) GetProductInventoryMaxAge(ctx context.Context, req *GetProductInventoryRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getProductInventoryMaxAge", req, &result)
	return result, err
}

// GetProductLocation 获取产品库位
func (s *Service) GetProductLocation(ctx context.Context, req *GetProductInventoryRequest) ([]ProductLocation, error) {
	var result []ProductLocation
	err := s.C.Do(ctx, "getProductLocation", req, &result)
	return result, err
}
