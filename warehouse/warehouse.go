// Package warehouse 提供易仓ERP仓库相关API的封装
package warehouse

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 仓库服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建仓库服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// Warehouse 仓库信息
type Warehouse struct {
	WarehouseCode   string `json:"warehouse_code"`
	WarehouseName   string `json:"warehouse_name"`
	WarehouseType   int    `json:"warehouse_type"`
	WarehouseStatus int    `json:"warehouse_status"`
	Country         string `json:"country"`
	Address         string `json:"address"`
	Contact         string `json:"contact"`
	Phone           string `json:"phone"`
}

// ShippingMethod 运输方式
type ShippingMethod struct {
	ShippingMethodCode string `json:"shipping_method_code"`
	ShippingMethodName string `json:"shipping_method_name"`
	WarehouseCode      string `json:"warehouse_code"`
}

// WarehouseLocation 仓库库位
type WarehouseLocation struct {
	LocationCode  string `json:"location_code"`
	LocationName  string `json:"location_name"`
	WarehouseCode string `json:"warehouse_code"`
	AreaCode      string `json:"area_code"`
	LocationType  int    `json:"location_type"`
	Status        int    `json:"status"`
}

// WarehouseLocationType 库位类型
type WarehouseLocationType struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

// WarehouseArea 仓库分区
type WarehouseArea struct {
	AreaCode      string `json:"area_code"`
	AreaName      string `json:"area_name"`
	WarehouseCode string `json:"warehouse_code"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// GetWarehouseListRequest 获取仓库信息（所有仓库）请求
type GetWarehouseListRequest struct {
	PageRequest
	WarehouseCode   string `json:"warehouse_code,omitempty"`
	WarehouseType   int    `json:"warehouse_type,omitempty"`
	WarehouseStatus int    `json:"warehouse_status,omitempty"`
	WarehouseName   string `json:"warehouse_name,omitempty"`
}

// SyncWarehouseRequest 创建仓库请求
type SyncWarehouseRequest struct {
	WarehouseCode string `json:"warehouse_code"`
	WarehouseName string `json:"warehouse_name"`
	WarehouseType int    `json:"warehouse_type,omitempty"`
	Country       string `json:"country,omitempty"`
	Address       string `json:"address,omitempty"`
	Contact       string `json:"contact,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

// SyncWarehouseLocationRequest 创建库位请求
type SyncWarehouseLocationRequest struct {
	LocationCode  string `json:"location_code"`
	LocationName  string `json:"location_name,omitempty"`
	WarehouseCode string `json:"warehouse_code"`
	AreaCode      string `json:"area_code,omitempty"`
	LocationType  int    `json:"location_type,omitempty"`
}

// SyncWarehouseLocationTypeRequest 创建库位类型请求
type SyncWarehouseLocationTypeRequest struct {
	TypeName string `json:"type_name"`
}

// SyncWarehouseAreaRequest 创建仓库分区请求
type SyncWarehouseAreaRequest struct {
	AreaCode      string `json:"area_code"`
	AreaName      string `json:"area_name"`
	WarehouseCode string `json:"warehouse_code"`
}

// ════════════════════════════════════════════
// 服务方法 (11个接口)
// ════════════════════════════════════════════

// --- 仓库查询 ---

// GetWarehouseList 获取仓库信息（所有仓库）
func (s *Service) GetWarehouseList(ctx context.Context, req *GetWarehouseListRequest) ([]Warehouse, error) {
	var result []Warehouse
	err := s.C.Do(ctx, "getWarehouseList", req, &result)
	return result, err
}

// GetWarehouse 获取仓库信息
func (s *Service) GetWarehouse(ctx context.Context, warehouseCode string) (*Warehouse, error) {
	var result Warehouse
	err := s.C.Do(ctx, "getWarehouse", map[string]string{"warehouse_code": warehouseCode}, &result)
	return &result, err
}

// GetWarehouseShippingMethod 获取仓库信息(关联运输方式)
func (s *Service) GetWarehouseShippingMethod(ctx context.Context, warehouseCode string) ([]ShippingMethod, error) {
	var result []ShippingMethod
	err := s.C.Do(ctx, "getWarehouseShippingMethod", map[string]string{"warehouse_code": warehouseCode}, &result)
	return result, err
}

// GetShippingMethodsettings 获取仓库对应的运输方式
func (s *Service) GetShippingMethodsettings(ctx context.Context, warehouseCode string) ([]ShippingMethod, error) {
	var result []ShippingMethod
	err := s.C.Do(ctx, "getShippingMethodsettings", map[string]string{"warehouse_code": warehouseCode}, &result)
	return result, err
}

// --- 库位/分区查询 ---

// GetWarehouseLocation 获取仓库库位
func (s *Service) GetWarehouseLocation(ctx context.Context, warehouseCode string) ([]WarehouseLocation, error) {
	var result []WarehouseLocation
	err := s.C.Do(ctx, "getWarehouseLocation", map[string]string{"warehouse_code": warehouseCode}, &result)
	return result, err
}

// GetWarehouseLocationType 获取库位类型
func (s *Service) GetWarehouseLocationType(ctx context.Context) ([]WarehouseLocationType, error) {
	var result []WarehouseLocationType
	err := s.C.Do(ctx, "getWarehouseLocationType", nil, &result)
	return result, err
}

// GetWarehouseArea 获取仓库分区
func (s *Service) GetWarehouseArea(ctx context.Context, warehouseCode string) ([]WarehouseArea, error) {
	var result []WarehouseArea
	err := s.C.Do(ctx, "getWarehouseArea", map[string]string{"warehouse_code": warehouseCode}, &result)
	return result, err
}

// --- 创建/编辑 ---

// SyncWarehouse 创建仓库
func (s *Service) SyncWarehouse(ctx context.Context, req *SyncWarehouseRequest) error {
	return s.C.Do(ctx, "syncWarehouse", req, nil)
}

// SyncWarehouseLocation 创建库位
func (s *Service) SyncWarehouseLocation(ctx context.Context, req *SyncWarehouseLocationRequest) error {
	return s.C.Do(ctx, "syncWarehouseLocation", req, nil)
}

// SyncWarehouseLocationType 创建库位类型
func (s *Service) SyncWarehouseLocationType(ctx context.Context, req *SyncWarehouseLocationTypeRequest) error {
	return s.C.Do(ctx, "syncWarehouseLocationType", req, nil)
}

// SyncWarehouseArea 创建仓库分区
func (s *Service) SyncWarehouseArea(ctx context.Context, req *SyncWarehouseAreaRequest) error {
	return s.C.Do(ctx, "syncWarehouseArea", req, nil)
}
