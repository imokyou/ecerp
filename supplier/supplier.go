// Package supplier 提供易仓ERP供应商相关API的封装
package supplier

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 供应商服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建供应商服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// Supplier 供应商信息
type Supplier struct {
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	ContactName  string `json:"contact_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Country      string `json:"country"`
	Currency     string `json:"currency"`
	PaymentTerms string `json:"payment_terms"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// SupplierDetail 供应商详情
type SupplierDetail struct {
	Supplier
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	TaxID       string `json:"tax_id"`
	Note        string `json:"note"`
}

// SupplierProduct 供应商产品
type SupplierProduct struct {
	SupplierCode  string  `json:"supplier_code"`
	SKU           string  `json:"sku"`
	SupplierSKU   string  `json:"supplier_sku"`
	ProductName   string  `json:"product_name"`
	PurchasePrice float64 `json:"purchase_price"`
	Currency      string  `json:"currency"`
	LeadTime      int     `json:"lead_time"`
	MinOrderQty   int     `json:"min_order_qty"`
}

// SupplierKpi 供应商KPI
type SupplierKpi struct {
	SupplierCode   string  `json:"supplier_code"`
	SupplierName   string  `json:"supplier_name"`
	DeliveryRate   float64 `json:"delivery_rate"`
	QualityRate    float64 `json:"quality_rate"`
	OnTimeRate     float64 `json:"on_time_rate"`
	Score          float64 `json:"score"`
	EvaluationTime string  `json:"evaluation_time"`
}

// Carrier 承运商
type Carrier struct {
	CarrierID   int    `json:"carrier_id"`
	CarrierCode string `json:"carrier_code"`
	CarrierName string `json:"carrier_name"`
	Status      int    `json:"status"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// GetSupplierListRequest 获取供应商列表请求
type GetSupplierListRequest struct {
	PageRequest
	SupplierCode string `json:"supplier_code,omitempty"`
	SupplierName string `json:"supplier_name,omitempty"`
	Status       int    `json:"status,omitempty"`
}

// SyncSupplierInfoRequest 建立/编辑供应商请求
type SyncSupplierInfoRequest struct {
	SupplierCode string `json:"supplier_code,omitempty"`
	SupplierName string `json:"supplier_name"`
	ContactName  string `json:"contact_name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Address      string `json:"address,omitempty"`
	Country      string `json:"country,omitempty"`
	Currency     string `json:"currency,omitempty"`
	PaymentTerms string `json:"payment_terms,omitempty"`
	BankName     string `json:"bank_name,omitempty"`
	BankAccount  string `json:"bank_account,omitempty"`
	TaxID        string `json:"tax_id,omitempty"`
	Note         string `json:"note,omitempty"`
}

// GetSupplierProductListRequest 获取供应商产品列表请求
type GetSupplierProductListRequest struct {
	PageRequest
	SupplierCode string `json:"supplier_code,omitempty"`
	SKU          string `json:"sku,omitempty"`
}

// SyncSupplierProductRequest 建立/编辑供应商产品请求
type SyncSupplierProductRequest struct {
	SupplierCode  string  `json:"supplier_code"`
	SKU           string  `json:"sku"`
	SupplierSKU   string  `json:"supplier_sku,omitempty"`
	PurchasePrice float64 `json:"purchase_price,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	LeadTime      int     `json:"lead_time,omitempty"`
	MinOrderQty   int     `json:"min_order_qty,omitempty"`
}

// GetSupplierKpiListRequest 供应商KPI列表请求
type GetSupplierKpiListRequest struct {
	PageRequest
	SupplierCode string `json:"supplier_code,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (7个接口)
// ════════════════════════════════════════════

// GetSupplierList 获取供应商列表
func (s *Service) GetSupplierList(ctx context.Context, req *GetSupplierListRequest) ([]Supplier, error) {
	var result []Supplier
	err := s.C.Do(ctx, "getSupplierList", req, &result)
	return result, err
}

// SyncSupplierInfo 建立/编辑供应商
func (s *Service) SyncSupplierInfo(ctx context.Context, req *SyncSupplierInfoRequest) error {
	return s.C.Do(ctx, "syncSupplierInfo", req, nil)
}

// GetSupplierProductList 获取供应商产品列表
func (s *Service) GetSupplierProductList(ctx context.Context, req *GetSupplierProductListRequest) ([]SupplierProduct, error) {
	var result []SupplierProduct
	err := s.C.Do(ctx, "getSupplierProductList", req, &result)
	return result, err
}

// GetSupplierInfo 获取指定供应商信息
func (s *Service) GetSupplierInfo(ctx context.Context, supplierCode string) (*SupplierDetail, error) {
	var result SupplierDetail
	err := s.C.Do(ctx, "getSupplierInfo", map[string]string{"supplier_code": supplierCode}, &result)
	return &result, err
}

// SyncSupplierProduct 建立/编辑供应商产品
func (s *Service) SyncSupplierProduct(ctx context.Context, req *SyncSupplierProductRequest) error {
	return s.C.Do(ctx, "syncSupplierProduct", req, nil)
}

// GetSupplierKpiList 供应商KPI列表信息
func (s *Service) GetSupplierKpiList(ctx context.Context, req *GetSupplierKpiListRequest) ([]SupplierKpi, error) {
	var result []SupplierKpi
	err := s.C.Do(ctx, "getSupplierKpiList", req, &result)
	return result, err
}

// GetCarrier 获取承运商数据
func (s *Service) GetCarrier(ctx context.Context) ([]Carrier, error) {
	var result []Carrier
	err := s.C.Do(ctx, "getCarrier", nil, &result)
	return result, err
}
