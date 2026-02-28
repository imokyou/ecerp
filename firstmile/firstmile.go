// Package firstmile 提供易仓ERP头程相关API的封装
package firstmile

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 头程服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建头程服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// TransferBatch 头程总单
type TransferBatch struct {
	BatchCode      string `json:"batch_code"`
	ShippingMethod string `json:"shipping_method"`
	WarehouseCode  string `json:"warehouse_code"`
	Status         int    `json:"status"`
	CreateTime     string `json:"create_time"`
}

// ShipmentInfo 总单详情
type ShipmentInfo struct {
	BatchCode      string `json:"batch_code"`
	ShippingMethod string `json:"shipping_method"`
	TrackingNumber string `json:"tracking_number"`
	Status         int    `json:"status"`
	Weight         string `json:"weight"`
}

// StockPlan 备货计划
type StockPlan struct {
	PlanID     int    `json:"plan_id"`
	PlanCode   string `json:"plan_code"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// FbaStockPlanData 货件创建数据
type FbaStockPlanData struct {
	ShipmentID string `json:"shipment_id"`
	SellerSKU  string `json:"seller_sku"`
	Quantity   int    `json:"quantity"`
	Status     int    `json:"status"`
}

// StpoOrder 海外仓头程单
type StpoOrder struct {
	StpoCode      string `json:"stpo_code"`
	WarehouseCode string `json:"warehouse_code"`
	Status        int    `json:"status"`
	CreateTime    string `json:"create_time"`
}

// CounterData 订柜管理数据
type CounterData struct {
	CounterID  int    `json:"counter_id"`
	CounterNo  string `json:"counter_no"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// DeliverBatch FBA发货管理数据
type DeliverBatch struct {
	BatchCode  string `json:"batch_code"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// DeliverOutboundBatch 头程发货单数据（FBA）
type DeliverOutboundBatch struct {
	BatchCode      string `json:"batch_code"`
	ShippingMethod string `json:"shipping_method"`
	Status         int    `json:"status"`
	CreateTime     string `json:"create_time"`
}

// FbaShipment FBA货件数据
type FbaShipment struct {
	ShipmentID   string `json:"shipment_id"`
	ShipmentName string `json:"shipment_name"`
	Status       string `json:"status"`
}

// TransferPackingItem 头程拣货单数据
type TransferPackingItem struct {
	PackingCode string `json:"packing_code"`
	BatchCode   string `json:"batch_code"`
	Status      int    `json:"status"`
}

// TransferProductPacking 装箱单管理
type TransferProductPacking struct {
	PackingCode string `json:"packing_code"`
	ProductSKU  string `json:"product_sku"`
	Quantity    int    `json:"quantity"`
}

// ShipBatch 头程出货单数据
type ShipBatch struct {
	BatchCode  string `json:"batch_code"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// TransferService 海外仓服务商
type TransferService struct {
	ServiceCode string `json:"service_code"`
	ServiceName string `json:"service_name"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// CreateShipmentBatchRequest 创建发货计划请求
type CreateShipmentBatchRequest struct {
	WarehouseCode  string `json:"warehouse_code"`
	ShippingMethod string `json:"shipping_method"`
}

// CreateStockPlanRequest 创建备货计划请求
type CreateStockPlanRequest struct {
	PlanCode string `json:"plan_code,omitempty"`
}

// CreateStockPlanShipmentRequest 创建fba货件请求
type CreateStockPlanShipmentRequest struct {
	PlanID     int    `json:"plan_id"`
	ShipmentID string `json:"shipment_id,omitempty"`
}

// CreateShipPlanRequest 头程创建请求
type CreateShipPlanRequest struct {
	WarehouseCode  string `json:"warehouse_code"`
	ShippingMethod string `json:"shipping_method,omitempty"`
}

// DeliveryTransferShipRequest 上传装箱发货请求
type DeliveryTransferShipRequest struct {
	BatchCode string `json:"batch_code"`
}

// UploadTransferFeeRequest 上传头程费用请求
type UploadTransferFeeRequest struct {
	BatchCode string  `json:"batch_code"`
	Fee       float64 `json:"fee"`
	Currency  string  `json:"currency,omitempty"`
}

// SerialDeliverImportRequest FBA发货管理导入SN请求
type SerialDeliverImportRequest struct {
	BatchCode string `json:"batch_code"`
	Items     []struct {
		SellerSKU    string `json:"seller_sku"`
		SerialNumber string `json:"serial_number"`
	} `json:"items"`
}

// ImportSnRequest 海外仓头程发货导入SN请求
type ImportSnRequest struct {
	StpoCode string `json:"stpo_code"`
	Items    []struct {
		ProductSKU   string `json:"product_sku"`
		SerialNumber string `json:"serial_number"`
	} `json:"items"`
}

// CreateAndEditServiceTransferPlanOrderRequest 创建/编辑头程计划单请求
type CreateAndEditServiceTransferPlanOrderRequest struct {
	PlanCode      string `json:"plan_code,omitempty"`
	WarehouseCode string `json:"warehouse_code"`
}

// ════════════════════════════════════════════
// 服务方法 (26个接口)
// ════════════════════════════════════════════

// --- 头程总单/计划 ---

// GetTransferBatchList 头程总单列表
func (s *Service) GetTransferBatchList(ctx context.Context, req *PageRequest) ([]TransferBatch, error) {
	var result []TransferBatch
	err := s.C.Do(ctx, "getTransferBatchList", req, &result)
	return result, err
}

// ShipmentAlwaysSingleInfo 总单详情
func (s *Service) ShipmentAlwaysSingleInfo(ctx context.Context, batchCode string) (*ShipmentInfo, error) {
	var result ShipmentInfo
	err := s.C.Do(ctx, "shipmentAlwaysSingleInfo", map[string]string{"batch_code": batchCode}, &result)
	return &result, err
}

// CreateShipmentBatch 创建发货计划
func (s *Service) CreateShipmentBatch(ctx context.Context, req *CreateShipmentBatchRequest) error {
	return s.C.Do(ctx, "createShipmentBatch", req, nil)
}

// CreateShipPlan 头程创建
func (s *Service) CreateShipPlan(ctx context.Context, req *CreateShipPlanRequest) error {
	return s.C.Do(ctx, "createShipPlan", req, nil)
}

// --- 备货/货件 ---

// CreateStockPlan 创建备货计划
func (s *Service) CreateStockPlan(ctx context.Context, req *CreateStockPlanRequest) error {
	return s.C.Do(ctx, "createStockPlan", req, nil)
}

// CreateStockPlanShipment 创建fba货件
func (s *Service) CreateStockPlanShipment(ctx context.Context, req *CreateStockPlanShipmentRequest) error {
	return s.C.Do(ctx, "createStockPlanShipment", req, nil)
}

// GetFbaStockPlan 获取货件创建数据
func (s *Service) GetFbaStockPlan(ctx context.Context, req *PageRequest) ([]FbaStockPlanData, error) {
	var result []FbaStockPlanData
	err := s.C.Do(ctx, "getFbaStockPlan", req, &result)
	return result, err
}

// GetFbaShipment 获取FBA货件数据
func (s *Service) GetFbaShipment(ctx context.Context, req *PageRequest) ([]FbaShipment, error) {
	var result []FbaShipment
	err := s.C.Do(ctx, "getFbaShipment", req, &result)
	return result, err
}

// --- FBA发货管理 ---

// GetDeliverBatch 获取FBA发货管理数据（待发货）
func (s *Service) GetDeliverBatch(ctx context.Context, req *PageRequest) ([]DeliverBatch, error) {
	var result []DeliverBatch
	err := s.C.Do(ctx, "getDeliverBatch", req, &result)
	return result, err
}

// GetDeliverOutboundBatch 获取头程发货单数据（FBA）
func (s *Service) GetDeliverOutboundBatch(ctx context.Context, req *PageRequest) ([]DeliverOutboundBatch, error) {
	var result []DeliverOutboundBatch
	err := s.C.Do(ctx, "getDeliverOutboundBatch", req, &result)
	return result, err
}

// VoidFbaDeliverOrder FBA发货管理-作废发货单
func (s *Service) VoidFbaDeliverOrder(ctx context.Context, batchCode string) error {
	return s.C.Do(ctx, "voidFbaDeliverOrder", map[string]string{"batch_code": batchCode}, nil)
}

// VoidShipment FBA备货-作废货件
func (s *Service) VoidShipment(ctx context.Context, shipmentID string) error {
	return s.C.Do(ctx, "voidShipment", map[string]string{"shipment_id": shipmentID}, nil)
}

// SerialDeliverImport FBA发货管理-导入SN
func (s *Service) SerialDeliverImport(ctx context.Context, req *SerialDeliverImportRequest) error {
	return s.C.Do(ctx, "serialDeliverImport", req, nil)
}

// UnMarkComplete fba发货单取消标记完成
func (s *Service) UnMarkComplete(ctx context.Context, doIDArr []int) error {
	return s.C.Do(ctx, "unMarkComplete", map[string]interface{}{"do_id_arr": doIDArr}, nil)
}

// --- 海外仓头程 ---

// GetStpoListNew 获取海外仓头程单(待发货)-新
func (s *Service) GetStpoListNew(ctx context.Context, req *PageRequest) ([]StpoOrder, error) {
	var result []StpoOrder
	err := s.C.Do(ctx, "getStpoListNew", req, &result)
	return result, err
}

// GetStpoList 获取海外仓头程单(待发货)---待下线
func (s *Service) GetStpoList(ctx context.Context, req *PageRequest) ([]StpoOrder, error) {
	var result []StpoOrder
	err := s.C.Do(ctx, "getStpoList", req, &result)
	return result, err
}

// ImportSn 海外仓头程发货-导入SN
func (s *Service) ImportSn(ctx context.Context, req *ImportSnRequest) error {
	return s.C.Do(ctx, "importSn", req, nil)
}

// GetTransferService 海外仓服务商代码
func (s *Service) GetTransferService(ctx context.Context) ([]TransferService, error) {
	var result []TransferService
	err := s.C.Do(ctx, "getTransferService", nil, &result)
	return result, err
}

// --- 订柜/装箱/拣货/出货 ---

// GetCounterList 获取头程订柜管理数据
func (s *Service) GetCounterList(ctx context.Context, req *PageRequest) ([]CounterData, error) {
	var result []CounterData
	err := s.C.Do(ctx, "getCounterList", req, &result)
	return result, err
}

// TransferProductPackingList 装箱单管理（Old）列表
func (s *Service) TransferProductPackingList(ctx context.Context, req *PageRequest) ([]TransferProductPacking, error) {
	var result []TransferProductPacking
	err := s.C.Do(ctx, "transferProductPackingList", req, &result)
	return result, err
}

// GetTransferPackingList 获取头程拣货单数据
func (s *Service) GetTransferPackingList(ctx context.Context, req *PageRequest) ([]TransferPackingItem, error) {
	var result []TransferPackingItem
	err := s.C.Do(ctx, "getTransferPackingList", req, &result)
	return result, err
}

// GetShipBatch 获取头程出货单数据
func (s *Service) GetShipBatch(ctx context.Context, req *PageRequest) ([]ShipBatch, error) {
	var result []ShipBatch
	err := s.C.Do(ctx, "getShipBatch", req, &result)
	return result, err
}

// TransferPackComplete 下架单回传（机器人定制）
func (s *Service) TransferPackComplete(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "transferPackComplete", req, nil)
}

// DeliveryTransferShip 上传装箱发货
func (s *Service) DeliveryTransferShip(ctx context.Context, req *DeliveryTransferShipRequest) error {
	return s.C.Do(ctx, "deliveryTransferShip", req, nil)
}

// --- 费用/计划单 ---

// UploadTransferFee 上传头程费用
func (s *Service) UploadTransferFee(ctx context.Context, req *UploadTransferFeeRequest) error {
	return s.C.Do(ctx, "uploadTransferFee", req, nil)
}

// CreateAndEditServiceTransferPlanOrder 创建/编辑头程计划单
func (s *Service) CreateAndEditServiceTransferPlanOrder(ctx context.Context, req *CreateAndEditServiceTransferPlanOrderRequest) error {
	return s.C.Do(ctx, "createAndEditServiceTransferPlanOrder", req, nil)
}
