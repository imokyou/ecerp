// Package amazonads 提供易仓ERP亚马逊广告相关API的封装
package amazonads

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 亚马逊广告服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建亚马逊广告服务
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

// AdStore 广告店铺
type AdStore struct {
	AccountID   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	Platform    string `json:"platform"`
	Site        string `json:"site"`
}

// TaskStatus 任务状态
type TaskStatus struct {
	TaskID     string `json:"task_id"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// ReportRequest 报告请求
type ReportRequest struct {
	AccountID int    `json:"account_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	TaskID    string `json:"task_id,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (6个接口)
// ════════════════════════════════════════════

// ReimbursementDownload 广告-获取赔偿原始报告
func (s *Service) ReimbursementDownload(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "ReimbursementDownload", req, &result)
	return result, err
}

// ReportDownload 广告-获取广告原始报告
func (s *Service) ReportDownload(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "ReportDownload", req, &result)
	return result, err
}

// AdInvoice 广告-发票下载
func (s *Service) AdInvoice(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "AdInvoice", req, &result)
	return result, err
}

// GetAdvertisingDetail 广告-获取广告明细数据
func (s *Service) GetAdvertisingDetail(ctx context.Context, req *ReportRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "GetAdvertisingDetail", req, &result)
	return result, err
}

// GetTasksStatus 广告-获取任务列表状态
func (s *Service) GetTasksStatus(ctx context.Context, req *ReportRequest) ([]TaskStatus, error) {
	var result []TaskStatus
	err := s.C.Do(ctx, "GetTasksStatus", req, &result)
	return result, err
}

// GetAuthAdStoreSiteList 广告-获取有广告数据的店铺列表
func (s *Service) GetAuthAdStoreSiteList(ctx context.Context) ([]AdStore, error) {
	var result []AdStore
	err := s.C.Do(ctx, "GetAuthAdStoreSiteList", nil, &result)
	return result, err
}
