// Package purchase 提供易仓ERP采购单相关API的封装
package purchase

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 采购单服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建采购单服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// PurchaseOrder 采购单
type PurchaseOrder struct {
	PurchaseCode  string              `json:"purchase_code"`
	WarehouseCode string              `json:"warehouse_code"`
	SupplierCode  string              `json:"supplier_code"`
	SupplierName  string              `json:"supplier_name"`
	Currency      string              `json:"currency"`
	TotalAmount   float64             `json:"total_amount"`
	Status        int                 `json:"status"`
	CreateTime    string              `json:"create_time"`
	AuditTime     string              `json:"audit_time"`
	Note          string              `json:"note"`
	Items         []PurchaseOrderItem `json:"items,omitempty"`
}

// PurchaseOrderItem 采购单明细
type PurchaseOrderItem struct {
	SKU       string  `json:"sku"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Amount    float64 `json:"amount"`
}

// PurchaseRequestOrder 申购单
type PurchaseRequestOrder struct {
	RequestCode   string `json:"request_code"`
	SKU           string `json:"sku"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	WarehouseCode string `json:"warehouse_code"`
	Status        int    `json:"status"`
	ProType       int    `json:"pro_type"`
	CreateTime    string `json:"create_time"`
}

// PurchasePlan 采购计划
type PurchasePlan struct {
	PlanID        int    `json:"plan_id"`
	SKU           string `json:"sku"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	WarehouseCode string `json:"warehouse_code"`
	Status        int    `json:"status"`
	CreateTime    string `json:"create_time"`
}

// PurchaseOrderFile 采购单附件
type PurchaseOrderFile struct {
	FileID       int    `json:"file_id"`
	FileName     string `json:"file_name"`
	FileURL      string `json:"file_url"`
	PurchaseCode string `json:"purchase_code"`
	CreateTime   string `json:"create_time"`
}

// PurchaseChange 采购单变更
type PurchaseChange struct {
	ChangeID     int    `json:"change_id"`
	PurchaseCode string `json:"purchase_code"`
	ChangeType   string `json:"change_type"`
	BeforeValue  string `json:"before_value"`
	AfterValue   string `json:"after_value"`
	Operator     string `json:"operator"`
	CreateTime   string `json:"create_time"`
}

// PurchaseTrackStatus 采购/财务跟单状态
type PurchaseTrackStatus struct {
	StatusID   int    `json:"status_id"`
	StatusName string `json:"status_name"`
}

// AbnormalReceipt 收货异常/QC异常
type AbnormalReceipt struct {
	RecordID     int    `json:"record_id"`
	PurchaseCode string `json:"purchase_code"`
	SKU          string `json:"sku"`
	Type         string `json:"type"`
	Quantity     int    `json:"quantity"`
	Status       int    `json:"status"`
	Note         string `json:"note"`
	CreateTime   string `json:"create_time"`
}

// PurchaseEta 采购单预计到货时间
type PurchaseEta struct {
	PurchaseCode string `json:"purchase_code"`
	EtaDate      string `json:"eta_date"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// SyncPurchaseOrderRequest 创建/编辑采购单请求
type SyncPurchaseOrderRequest struct {
	PurchaseCode  string              `json:"purchase_code,omitempty"` // 空则新建
	WarehouseCode string              `json:"warehouse_code"`
	SupplierCode  string              `json:"supplier_code"`
	Currency      string              `json:"currency,omitempty"`
	Note          string              `json:"note,omitempty"`
	Items         []PurchaseOrderItem `json:"items"`
}

// GetPurchaseOrdersRequest 查询采购单信息请求
type GetPurchaseOrdersRequest struct {
	PageRequest
	PurchaseCode  string `json:"purchase_code,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
	SupplierCode  string `json:"supplier_code,omitempty"`
	Status        int    `json:"status,omitempty"`
	SKU           string `json:"sku,omitempty"`
	DateType      string `json:"date_type,omitempty"`
	DateFor       string `json:"date_for,omitempty"`
	DateTo        string `json:"date_to,omitempty"`
	IsUrgent      int    `json:"is_urgent,omitempty"`
}

// GetPurchaseRequestOrdersRequest 获取申购单请求
type GetPurchaseRequestOrdersRequest struct {
	PageRequest
	ProType  int    `json:"pro_type,omitempty"`
	OrgID    int    `json:"org_id,omitempty"`
	CodeType string `json:"code_type,omitempty"`
	Code     string `json:"code,omitempty"`
	SKUType  string `json:"sku_type,omitempty"`
	SKU      string `json:"sku,omitempty"`
	DateType string `json:"date_type,omitempty"`
	DateFor  string `json:"date_for,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	UserType string `json:"user_type,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
}

// GetPurchasePlanRequest 获取采购计划请求
type GetPurchasePlanRequest struct {
	PageRequest
	SKU           string `json:"sku,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
}

// HandlingExceptionRequest 处理质检/收货异常请求
type HandlingExceptionRequest struct {
	RecordID int    `json:"record_id"`
	Action   string `json:"action,omitempty"`
	Note     string `json:"note,omitempty"`
}

// GetPurchaseChangeRequest 采购单变更列表请求
type GetPurchaseChangeRequest struct {
	PageRequest
	PurchaseCode string `json:"purchase_code,omitempty"`
}

// VerifyPurchaseRequest 审核采购单请求
type VerifyPurchaseRequest struct {
	PurchaseCode string `json:"purchase_code"`
}

// GetPurchaseEtaRequest 查询采购单预计到货时间请求
type GetPurchaseEtaRequest struct {
	PurchaseCode string `json:"purchase_code,omitempty"`
}

// ProductMatching1688Request 1688商品匹配请求
type ProductMatching1688Request struct {
	SKU string `json:"sku"`
}

// SyncBatchPurchaseOrdersRequest 批量导入采购单请求
type SyncBatchPurchaseOrdersRequest struct {
	Orders []SyncPurchaseOrderRequest `json:"orders"`
}

// SyncPurchaseTrackingNoteRequest 采购单跟踪备注请求
type SyncPurchaseTrackingNoteRequest struct {
	PurchaseCode string `json:"purchase_code"`
	Note         string `json:"note"`
}

// GetAbnormalReceiptListRequest 查询收货/QC异常处理请求
type GetAbnormalReceiptListRequest struct {
	PageRequest
	PurchaseCode string `json:"purchase_code,omitempty"`
	SKU          string `json:"sku,omitempty"`
	Status       int    `json:"status,omitempty"`
}

// RevocationPurchaseRequest 撤销采购单请求
type RevocationPurchaseRequest struct {
	PurchaseCode string `json:"purchase_code"`
}

// ════════════════════════════════════════════
// 服务方法 (20个接口)
// ════════════════════════════════════════════

// SyncPurchaseOrders 创建/编辑采购单
func (s *Service) SyncPurchaseOrders(ctx context.Context, req *SyncPurchaseOrderRequest) error {
	return s.C.Do(ctx, "syncPurchaseOrders", req, nil)
}

// GetPurchaseOrders 查询采购单信息
func (s *Service) GetPurchaseOrders(ctx context.Context, req *GetPurchaseOrdersRequest) ([]PurchaseOrder, error) {
	var result []PurchaseOrder
	err := s.C.Do(ctx, "getPurchaseOrders", req, &result)
	return result, err
}

// GetPurchaseRequestOrdersNew 获取申购单（新）
func (s *Service) GetPurchaseRequestOrdersNew(ctx context.Context, req *GetPurchaseRequestOrdersRequest) ([]PurchaseRequestOrder, error) {
	var result []PurchaseRequestOrder
	err := s.C.Do(ctx, "getPurchaseRequestOrdersNew", req, &result)
	return result, err
}

// GetPurchaseRequestOrders 获取申购单（旧）
func (s *Service) GetPurchaseRequestOrders(ctx context.Context, req *GetPurchaseRequestOrdersRequest) ([]PurchaseRequestOrder, error) {
	var result []PurchaseRequestOrder
	err := s.C.Do(ctx, "getPurchaseRequestOrders", req, &result)
	return result, err
}

// GetPurchasePlan 获取采购计划
func (s *Service) GetPurchasePlan(ctx context.Context, req *GetPurchasePlanRequest) ([]PurchasePlan, error) {
	var result []PurchasePlan
	err := s.C.Do(ctx, "getPurchasePlan", req, &result)
	return result, err
}

// HandlingQcExceptions 处理质检异常
func (s *Service) HandlingQcExceptions(ctx context.Context, req *HandlingExceptionRequest) error {
	return s.C.Do(ctx, "handlingQcExceptions", req, nil)
}

// PurchaseForceCompletion 采购单强制完成
func (s *Service) PurchaseForceCompletion(ctx context.Context, purchaseCode string) error {
	return s.C.Do(ctx, "purchaseForceCompletion", map[string]string{"purchase_code": purchaseCode}, nil)
}

// HandlingReceivingExceptions 处理收货异常
func (s *Service) HandlingReceivingExceptions(ctx context.Context, req *HandlingExceptionRequest) error {
	return s.C.Do(ctx, "handlingReceivingExceptions", req, nil)
}

// GetPurchaseOrderFiles 获取采购单附件
func (s *Service) GetPurchaseOrderFiles(ctx context.Context, purchaseCode string) ([]PurchaseOrderFile, error) {
	var result []PurchaseOrderFile
	err := s.C.Do(ctx, "getPurchaseOrderFiles", map[string]string{"purchase_code": purchaseCode}, &result)
	return result, err
}

// GetPurchaseChange 采购单变更列表
func (s *Service) GetPurchaseChange(ctx context.Context, req *GetPurchaseChangeRequest) ([]PurchaseChange, error) {
	var result []PurchaseChange
	err := s.C.Do(ctx, "getPurchaseChange", req, &result)
	return result, err
}

// VerifyPurchase 审核采购单
func (s *Service) VerifyPurchase(ctx context.Context, req *VerifyPurchaseRequest) error {
	return s.C.Do(ctx, "verifyPurchase", req, nil)
}

// GetPurchaseOrdersDateEta 查询采购单预计到货时间
func (s *Service) GetPurchaseOrdersDateEta(ctx context.Context, req *GetPurchaseEtaRequest) ([]PurchaseEta, error) {
	var result []PurchaseEta
	err := s.C.Do(ctx, "getPurchaseOrdersDateEta", req, &result)
	return result, err
}

// ProductMatching1688 1688商品匹配
func (s *Service) ProductMatching1688(ctx context.Context, req *ProductMatching1688Request) error {
	return s.C.Do(ctx, "productMatching1688", req, nil)
}

// SyncBatchPurchaseOrders 批量导入采购单
func (s *Service) SyncBatchPurchaseOrders(ctx context.Context, req *SyncBatchPurchaseOrdersRequest) error {
	return s.C.Do(ctx, "syncBatchPurchaseOrders", req, nil)
}

// SyncPurchaseTrackingNote 采购单跟踪备注
func (s *Service) SyncPurchaseTrackingNote(ctx context.Context, req *SyncPurchaseTrackingNoteRequest) error {
	return s.C.Do(ctx, "syncPurchaseTrackingNote", req, nil)
}

// AbnormalReceiptList 查询收货异常处理
func (s *Service) AbnormalReceiptList(ctx context.Context, req *GetAbnormalReceiptListRequest) ([]AbnormalReceipt, error) {
	var result []AbnormalReceipt
	err := s.C.Do(ctx, "abnormalReceiptList", req, &result)
	return result, err
}

// QcReceiptList 查询QC异常处理
func (s *Service) QcReceiptList(ctx context.Context, req *GetAbnormalReceiptListRequest) ([]AbnormalReceipt, error) {
	var result []AbnormalReceipt
	err := s.C.Do(ctx, "qcReceiptList", req, &result)
	return result, err
}

// RevocationPurchase 撤销采购单
func (s *Service) RevocationPurchase(ctx context.Context, req *RevocationPurchaseRequest) error {
	return s.C.Do(ctx, "revocationPurchase", req, nil)
}

// GetFinanceTrackStatus 获取财务跟单状态
func (s *Service) GetFinanceTrackStatus(ctx context.Context) ([]PurchaseTrackStatus, error) {
	var result []PurchaseTrackStatus
	err := s.C.Do(ctx, "getFinanceTrackStatus", nil, &result)
	return result, err
}

// GetPurchaseTrackStatus 获取采购跟单状态
func (s *Service) GetPurchaseTrackStatus(ctx context.Context) ([]PurchaseTrackStatus, error) {
	var result []PurchaseTrackStatus
	err := s.C.Do(ctx, "getPurchaseTrackStatus", nil, &result)
	return result, err
}
