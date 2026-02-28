// Package expense 提供易仓ERP费用相关API的封装
package expense

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 费用服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建费用服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// PaymentAccount 收款方支付方式
type PaymentAccount struct {
	AccountID   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	AccountType string `json:"account_type"`
}

// PaymentBank 银行卡信息
type PaymentBank struct {
	BankID     int    `json:"bank_id"`
	BankName   string `json:"bank_name"`
	BankNo     string `json:"bank_no"`
	CardHolder string `json:"card_holder"`
}

// PurchaseCompany 付款主体
type PurchaseCompany struct {
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}

// PurchasePayment 采购结算
type PurchasePayment struct {
	PaymentID   int     `json:"payment_id"`
	PaymentCode string  `json:"payment_code"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Status      int     `json:"status"`
	CreateTime  string  `json:"create_time"`
}

// OrderCostDetail 订单费用明细
type OrderCostDetail struct {
	OrderCode    string  `json:"order_code"`
	SKU          string  `json:"sku"`
	ProductName  string  `json:"product_name"`
	ProductCost  float64 `json:"product_cost"`
	ShippingCost float64 `json:"shipping_cost"`
	PlatformFee  float64 `json:"platform_fee"`
	TotalCost    float64 `json:"total_cost"`
}

// ServiceProviderAccount 服务商账单
type ServiceProviderAccount struct {
	AccountID    int     `json:"account_id"`
	ProviderCode string  `json:"provider_code"`
	ProviderName string  `json:"provider_name"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	Status       int     `json:"status"`
	CreateTime   string  `json:"create_time"`
}

// FnskuCost FNSKU产品成本
type FnskuCost struct {
	FNSKU       string  `json:"fnsku"`
	SKU         string  `json:"sku"`
	ProductName string  `json:"product_name"`
	Cost        float64 `json:"cost"`
}

// TrialFeeResult 费用试算结果
type TrialFeeResult struct {
	OrderCode   string  `json:"order_code"`
	ShippingFee float64 `json:"shipping_fee"`
	HandlingFee float64 `json:"handling_fee"`
	TotalFee    float64 `json:"total_fee"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// PaymentRecordsVerifyRequest 确认付款请求
type PaymentRecordsVerifyRequest struct {
	FprRefCode    string `json:"fpr_ref_code"`
	PaymentMethod int    `json:"payment_method,omitempty"`
	SbaID         int    `json:"sba_id,omitempty"`
	PcID          int    `json:"pc_id,omitempty"`
	BaID          int    `json:"ba_id,omitempty"`
}

// ApprovePurchasePaymentRequest 采购结算审核请求
type ApprovePurchasePaymentRequest struct {
	PaymentID int    `json:"payment_id"`
	Action    string `json:"action"`
	Note      string `json:"note,omitempty"`
}

// GetCostByFnskuRequest 获取FNSKU产品成本请求
type GetCostByFnskuRequest struct {
	FNSKU     string `json:"fnsku"`
	AccountID int    `json:"account_id,omitempty"`
}

// ResetOrderCostRequest 添加重新计费的订单请求
type ResetOrderCostRequest struct {
	OrderCodes []string `json:"order_codes"`
}

// GetOrderCostDetailRequest 订单费用和成本明细请求
type GetOrderCostDetailRequest struct {
	PageRequest
	OrderCode string `json:"order_code,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (14个接口)
// ════════════════════════════════════════════

// --- 付款单（出纳付款）---

// PaymentRecordsVerify 付款单（出纳付款）--确认付款
func (s *Service) PaymentRecordsVerify(ctx context.Context, req *PaymentRecordsVerifyRequest) error {
	return s.C.Do(ctx, "paymentRecordsVerify", req, nil)
}

// GetSuPaymentAccount 付款单（出纳付款）--收款方-支付方式
func (s *Service) GetSuPaymentAccount(ctx context.Context) ([]PaymentAccount, error) {
	var result []PaymentAccount
	err := s.C.Do(ctx, "getSuPaymentAccount", nil, &result)
	return result, err
}

// GetReceivePaymentBank 付款单（出纳付款）--收款方-银行卡
func (s *Service) GetReceivePaymentBank(ctx context.Context) ([]PaymentBank, error) {
	var result []PaymentBank
	err := s.C.Do(ctx, "getReceivePaymentBank", nil, &result)
	return result, err
}

// GetPurchaseCompany 付款单（出纳付款）--付款方-付款主体
func (s *Service) GetPurchaseCompany(ctx context.Context) ([]PurchaseCompany, error) {
	var result []PurchaseCompany
	err := s.C.Do(ctx, "getPurchaseCompany", nil, &result)
	return result, err
}

// GetPaymentBank 付款单（出纳付款）--付款方-银行卡
func (s *Service) GetPaymentBank(ctx context.Context) ([]PaymentBank, error) {
	var result []PaymentBank
	err := s.C.Do(ctx, "getPaymentBank", nil, &result)
	return result, err
}

// --- 采购结算 ---

// ApprovePurchasePayment 采购结算审核\复审
func (s *Service) ApprovePurchasePayment(ctx context.Context, req *ApprovePurchasePaymentRequest) error {
	return s.C.Do(ctx, "approvePurchasePayment", req, nil)
}

// GetPurchasePayment 采购结算（付款通知）列表
func (s *Service) GetPurchasePayment(ctx context.Context, req *PageRequest) ([]PurchasePayment, error) {
	var result []PurchasePayment
	err := s.C.Do(ctx, "getPurchasePayment", req, &result)
	return result, err
}

// --- 费用查询 ---

// GetCostByFnsku 根据FNSKU店铺，获取产品成本
func (s *Service) GetCostByFnsku(ctx context.Context, req *GetCostByFnskuRequest) (*FnskuCost, error) {
	var result FnskuCost
	err := s.C.Do(ctx, "getCostByFnsku", req, &result)
	return &result, err
}

// GetOrderCostDetailSku 订单费用和成本明细(按SKU)
func (s *Service) GetOrderCostDetailSku(ctx context.Context, req *GetOrderCostDetailRequest) ([]OrderCostDetail, error) {
	var result []OrderCostDetail
	err := s.C.Do(ctx, "getOrderCostDetailSku", req, &result)
	return result, err
}

// GetOrderCostDetail 订单费用和成本明细(按订单)
func (s *Service) GetOrderCostDetail(ctx context.Context, req *GetOrderCostDetailRequest) ([]OrderCostDetail, error) {
	var result []OrderCostDetail
	err := s.C.Do(ctx, "getOrderCostDetail", req, &result)
	return result, err
}

// --- 服务商/计费 ---

// GetServiceProviderAccount 服务商账单列表
func (s *Service) GetServiceProviderAccount(ctx context.Context, req *PageRequest) ([]ServiceProviderAccount, error) {
	var result []ServiceProviderAccount
	err := s.C.Do(ctx, "getServiceProviderAccount", req, &result)
	return result, err
}

// ResetOrderCost 添加重新计费的订单
func (s *Service) ResetOrderCost(ctx context.Context, req *ResetOrderCostRequest) error {
	return s.C.Do(ctx, "resetOrderCost", req, nil)
}

// --- 费用试算 ---

// BatchTrialFeeByAllApi 费用试算字段
func (s *Service) BatchTrialFeeByAllApi(ctx context.Context, req map[string]interface{}) ([]TrialFeeResult, error) {
	var result []TrialFeeResult
	err := s.C.Do(ctx, "batchTrialFeeByAllApi", req, &result)
	return result, err
}

// LoopBatchTrialFeeByAllApi 批量费用试算字段
func (s *Service) LoopBatchTrialFeeByAllApi(ctx context.Context, req map[string]interface{}) ([]TrialFeeResult, error) {
	var result []TrialFeeResult
	err := s.C.Do(ctx, "loopBatchTrialFeeByAllApi", req, &result)
	return result, err
}
