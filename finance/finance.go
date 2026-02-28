// Package finance 提供易仓ERP财务相关API的封装
package finance

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 财务服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建财务服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// FinancialReportRequest 财务报告请求
type FinancialReportRequest struct {
	PageRequest
	UserAccount    string `json:"user_account,omitempty"`
	UnitCurrency   string `json:"unit_currency,omitempty"`
	TimeZoneType   int    `json:"time_zone_type,omitempty"`
	TimeType       int    `json:"time_type,omitempty"`
	AmazonOrderID  string `json:"amazon_order_id,omitempty"`
	SettlementType string `json:"settlement_type,omitempty"`
	StartDate      string `json:"start_date,omitempty"`
	EndDate        string `json:"end_date,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (20个接口)
// ════════════════════════════════════════════

// --- 财务利润报表 ---

// GetFinancialOrderReportDetail 获取财务订单维度利润详情
func (s *Service) GetFinancialOrderReportDetail(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFinancialOrderReportDetail", req, &result)
	return result, err
}

// GetFinancialOrderReportList 获取财务订单维度利润列表
func (s *Service) GetFinancialOrderReportList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFinancialOrderReportList", req, &result)
	return result, err
}

// GetFinancialSellerSKUReportList 获取财务SellerSKU维度利润列表
func (s *Service) GetFinancialSellerSKUReportList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFinancialSellerSKUReportList", req, &result)
	return result, err
}

// --- 付款 ---

// GetFeeTransferRecords 付款异常（收款通知）-列表
func (s *Service) GetFeeTransferRecords(ctx context.Context, req *PageRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFeeTransferRecords", req, &result)
	return result, err
}

// GetFeePaymentRecords 付款单（出纳付款）列表
func (s *Service) GetFeePaymentRecords(ctx context.Context, req *PageRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFeePaymentRecords", req, &result)
	return result, err
}

// --- 库存分类账 ---

// GetFbaLedgerDetailList 获取库存分类账明细列表
func (s *Service) GetFbaLedgerDetailList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFbaLedgerDetailList", req, &result)
	return result, err
}

// GetFbaLedgerSummaryList 按月/按天获取库存分类账列表
func (s *Service) GetFbaLedgerSummaryList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFbaLedgerSummaryList", req, &result)
	return result, err
}

// --- 金蝶集成 ---

// GetFcsProductList 金蝶集成-获取产品配对信息
func (s *Service) GetFcsProductList(ctx context.Context, req *PageRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFcsProductList", req, &result)
	return result, err
}

// --- FBA报告 ---

// GetFbaInventoryReport 获取FBA进销存月报列表
func (s *Service) GetFbaInventoryReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFbaInventoryReport", req, &result)
	return result, err
}

// GetDailyInventoryHistoryReport 获取FBA每日库存报告列表
func (s *Service) GetDailyInventoryHistoryReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getDailyInventoryHistoryReport", req, &result)
	return result, err
}

// GetAdjustmentsReport 获取FBA调整报告列表
func (s *Service) GetAdjustmentsReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getAdjustmentsReport", req, &result)
	return result, err
}

// GetSalesReport 获取FBA销售报告列表
func (s *Service) GetSalesReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getSalesReport", req, &result)
	return result, err
}

// GetReturnsReport 获取FBA退货报告列表
func (s *Service) GetReturnsReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getReturnsReport", req, &result)
	return result, err
}

// GetRemovalOrderDetailReport 获取移除订单列表
func (s *Service) GetRemovalOrderDetailReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getRemovalOrderDetailReport", req, &result)
	return result, err
}

// GetInventoryEventDetailReport 获取库存事件详情列表
func (s *Service) GetInventoryEventDetailReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getInventoryEventDetailReport", req, &result)
	return result, err
}

// GetMonthlyReport 获取FBA月末库存报告列表
func (s *Service) GetMonthlyReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getMonthlyReport", req, &result)
	return result, err
}

// GetRemovalShipmentReport 获取移除发货报告
func (s *Service) GetRemovalShipmentReport(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getRemovalShipmentReport", req, &result)
	return result, err
}

// --- 店铺/交易 ---

// GetFinancialEventGroupList 获取店铺划款记录
func (s *Service) GetFinancialEventGroupList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getFinancialEventGroupList", req, &result)
	return result, err
}

// GetTransactionReportDetailList 获取transaction交易明细
func (s *Service) GetTransactionReportDetailList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getTransactionReportDetailList", req, &result)
	return result, err
}

// GetInventoryReceiptsList 获取FBA到货报告列表
func (s *Service) GetInventoryReceiptsList(ctx context.Context, req *FinancialReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getInventoryReceiptsList", req, &result)
	return result, err
}
