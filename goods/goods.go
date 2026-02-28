// Package goods 提供易仓ERP商品相关API的封装
package goods

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 商品服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建商品服务
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

// ReportRequest 报告请求
type ReportRequest struct {
	PageRequest
	AccountID int    `json:"account_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (34个接口)
// ════════════════════════════════════════════

// --- 移除订单 ---

// NewRemovalOrderList 移除订单列表new
func (s *Service) NewRemovalOrderList(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "NewRemovalOrderList", req, &result)
	return result, err
}

// RemovalOrderList 移除订单列表
func (s *Service) RemovalOrderList(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "RemovalOrderList", req, &result)
	return result, err
}

// RemovalShipmentDetail 移除订单明细接口
func (s *Service) RemovalShipmentDetail(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "RemovalShipmentDetail", req, &result)
	return result, err
}

// --- 结算/回款 ---

// AmazonSettlementReport 回款明细
func (s *Service) AmazonSettlementReport(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonSettlementReport", req, &result)
	return result, err
}

// AmazonSettlementReportDataFlatFile 结算报告V2版
func (s *Service) AmazonSettlementReportDataFlatFile(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonSettlementReportDataFlatFile", req, &result)
	return result, err
}

// --- FBA仓储费报告 ---

// AmazonFbaStorageFeeChargesNew FBA月仓租费报告-页面版
func (s *Service) AmazonFbaStorageFeeChargesNew(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaStorageFeeChargesNew", req, &result)
	return result, err
}

// AmazonFbaStorageFeeCharges FBA月仓租费报告
func (s *Service) AmazonFbaStorageFeeCharges(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaStorageFeeCharges", req, &result)
	return result, err
}

// AmazonFbaFulfillmentLongtermStorageFeeCharges FBA长期仓租费报告
func (s *Service) AmazonFbaFulfillmentLongtermStorageFeeCharges(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaFulfillmentLongtermStorageFeeCharges", req, &result)
	return result, err
}

// --- 索赔 ---

// FbaReimbursementProposal 买家退款未退货索赔（内测版）
func (s *Service) FbaReimbursementProposal(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "FbaReimbursementProposal", req, &result)
	return result, err
}

// FbaReimbursementProposalDetail 买家退款未退货索赔-索赔记录（内测版）
func (s *Service) FbaReimbursementProposalDetail(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "FbaReimbursementProposalDetail", req, &result)
	return result, err
}

// AmazonClaimProposal 库内丢失货损索赔（内侧版）
func (s *Service) AmazonClaimProposal(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonClaimProposal", req, &result)
	return result, err
}

// AmazonClaimProposalDetail 库内丢失货损索赔详情（内测版）
func (s *Service) AmazonClaimProposalDetail(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonClaimProposalDetail", req, &result)
	return result, err
}

// --- Listing ---

// GetWalmartListing 获取Walmart平台listing列表
func (s *Service) GetWalmartListing(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getWalmartListing", req, &result)
	return result, err
}

// GetWayfairListing 获取wayfair平台listing列表
func (s *Service) GetWayfairListing(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getWayfairListing", req, &result)
	return result, err
}

// AmazonListing 亚马逊listing列表
func (s *Service) AmazonListing(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonListing", req, &result)
	return result, err
}

// ListingPerformance listing表现接口
func (s *Service) ListingPerformance(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "ListingPerformance", req, &result)
	return result, err
}

// ListingSummaryOriginal listing表现-日维度接口
func (s *Service) ListingSummaryOriginal(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "ListingSummaryOriginal", req, &result)
	return result, err
}

// GetItemList 获取在线商品列表
func (s *Service) GetItemList(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getItemList", req, &result)
	return result, err
}

// --- Amazon报告 ---

// AmazonFbaShipmentReplacementData 换货报告
func (s *Service) AmazonFbaShipmentReplacementData(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaShipmentReplacementData", req, &result)
	return result, err
}

// AmazonFbaShipmentPromotionData 亚马逊物流促销报告
func (s *Service) AmazonFbaShipmentPromotionData(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaShipmentPromotionData", req, &result)
	return result, err
}

// AmazonSellerPerformanceReport 店铺绩效 - ODR报告V2版
func (s *Service) AmazonSellerPerformanceReport(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonSellerPerformanceReport", req, &result)
	return result, err
}

// AmazonSalesAndTrafficReport 获取销售和流量报告
func (s *Service) AmazonSalesAndTrafficReport(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonSalesAndTrafficReport", req, &result)
	return result, err
}

// AmazonFbaReimbursement FBA赔偿报告
func (s *Service) AmazonFbaReimbursement(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaReimbursement", req, &result)
	return result, err
}

// AmazonFbaEstimatedFbaFees FBA预计费用报告
func (s *Service) AmazonFbaEstimatedFbaFees(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaEstimatedFbaFees", req, &result)
	return result, err
}

// --- FBA库存 ---

// AmazonFbaFulfillmentInventory FBA健康库存报告
func (s *Service) AmazonFbaFulfillmentInventory(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaFulfillmentInventory", req, &result)
	return result, err
}

// FbaInventory FBA库存接口
func (s *Service) FbaInventory(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "FbaInventory", req, &result)
	return result, err
}

// AmazonReservedInventory 预留库存报告
func (s *Service) AmazonReservedInventory(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonReservedInventory", req, &result)
	return result, err
}

// AmazonFbaMyiAllInventory 全量库存
func (s *Service) AmazonFbaMyiAllInventory(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaMyiAllInventory", req, &result)
	return result, err
}

// AmazonInventoryCapacity 库容绩效-接口
func (s *Service) AmazonInventoryCapacity(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonInventoryCapacity", req, &result)
	return result, err
}

// --- FBA退货/订单 ---

// AmazonFbaFulfillmentCustomerReturnsData FBA退货、退款
func (s *Service) AmazonFbaFulfillmentCustomerReturnsData(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonFbaFulfillmentCustomerReturnsData", req, &result)
	return result, err
}

// AmazonOrderDetail Amazon原始订单明细接口
func (s *Service) AmazonOrderDetail(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonOrderDetail", req, &result)
	return result, err
}

// AmazonOrderOriginal Amazon原始订单接口
func (s *Service) AmazonOrderOriginal(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonOrderOriginal", req, &result)
	return result, err
}

// --- 评价/Review ---

// Feedback Feedback-查询接口
func (s *Service) Feedback(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "Feedback", req, &result)
	return result, err
}

// AmazonReview 亚马逊review评语接口
func (s *Service) AmazonReview(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AmazonReview", req, &result)
	return result, err
}
