// Package fbasta 提供易仓ERP FBASTA (FBA发货助手) 相关API的封装
package fbasta

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service FBASTA服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建FBASTA服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// InboundPlan 工作流程
type InboundPlan struct {
	InboundPlanID string `json:"inbound_plan_id"`
	Status        string `json:"status"`
	CreateTime    string `json:"create_time"`
}

// PlacementOption 分仓选项
type PlacementOption struct {
	PlacementOptionID string `json:"placement_option_id"`
	Status            string `json:"status"`
	FulfillmentCenter string `json:"fulfillment_center"`
}

// PackingInformation 包装组信息
type PackingInformation struct {
	PackingGroupID string `json:"packing_group_id"`
	Status         string `json:"status"`
}

// TransportationOption 运输选项
type TransportationOption struct {
	TransportationOptionID string  `json:"transportation_option_id"`
	CarrierName            string  `json:"carrier_name"`
	Price                  float64 `json:"price"`
	Currency               string  `json:"currency"`
}

// DeliveryWindowOption 交付窗口选项
type DeliveryWindowOption struct {
	DeliveryWindowID string `json:"delivery_window_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}

// Operation 异步任务
type Operation struct {
	OperationID string `json:"operation_id"`
	Status      string `json:"status"`
	Result      string `json:"result"`
}

// Label 面单
type Label struct {
	LabelURL    string `json:"label_url"`
	LabelFormat string `json:"label_format"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// CreateInboundPlanRequest 创建工作流程请求
type CreateInboundPlanRequest struct {
	DestinationMarketplaces []string               `json:"destination_marketplaces"`
	Items                   []InboundPlanItem      `json:"items"`
	SourceAddress           map[string]interface{} `json:"source_address"`
}

// InboundPlanItem 工作流程项
type InboundPlanItem struct {
	SellerSKU string `json:"seller_sku"`
	Quantity  int    `json:"quantity"`
	ASIN      string `json:"asin,omitempty"`
}

// DraftPackingInformationRequest 保存装箱请求
type DraftPackingInformationRequest struct {
	FipID          string       `json:"fip_id"`
	FippiID        string       `json:"fippi_id"`
	BoxPackingType string       `json:"box_packing_type,omitempty"`
	Boxes          []PackingBox `json:"boxes"`
}

// PackingBox 装箱信息
type PackingBox struct {
	SellerSKU   string `json:"seller_sku"`
	Quantity    int    `json:"quantity"`
	BoxQuantity int    `json:"box_quantity"`
}

// ConfirmDeliveryWindowAndTransportationOptionsRequest 确认运输选项和交货窗口请求
type ConfirmDeliveryWindowAndTransportationOptionsRequest struct {
	InboundPlanID        string `json:"inbound_plan_id"`
	ShipmentID           string `json:"shipment_id"`
	TransportationOption string `json:"transportation_option_id"`
	DeliveryWindow       string `json:"delivery_window_id,omitempty"`
}

// UpdateShipmentTrackingDetailsRequest 更新跟踪信息请求
type UpdateShipmentTrackingDetailsRequest struct {
	InboundPlanID  string `json:"inbound_plan_id"`
	ShipmentID     string `json:"shipment_id"`
	TrackingNumber string `json:"tracking_number"`
	CarrierName    string `json:"carrier_name,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (20个接口)
// ════════════════════════════════════════════

// --- 工作流程 (3个) ---

// CreateInboundPlan 创建工作流程
func (s *Service) CreateInboundPlan(ctx context.Context, req *CreateInboundPlanRequest) error {
	return s.C.Do(ctx, "createInboundPlan", req, nil)
}

// GetInboundPlan 查询工作流程
func (s *Service) GetInboundPlan(ctx context.Context, inboundPlanID string) (*InboundPlan, error) {
	var result InboundPlan
	err := s.C.Do(ctx, "getInboundPlan", map[string]string{"inbound_plan_id": inboundPlanID}, &result)
	return &result, err
}

// CancelInboundPlan 取消工作流程
func (s *Service) CancelInboundPlan(ctx context.Context, inboundPlanID string) error {
	return s.C.Do(ctx, "cancelInboundPlan", map[string]string{"inbound_plan_id": inboundPlanID}, nil)
}

// --- 分仓选项 (4个) ---

// GeneratePlacementOptions 生成分仓选项
func (s *Service) GeneratePlacementOptions(ctx context.Context, inboundPlanID string) error {
	return s.C.Do(ctx, "generatePlacementOptions", map[string]string{"inbound_plan_id": inboundPlanID}, nil)
}

// GetPlacementOptions 查询分仓选项
func (s *Service) GetPlacementOptions(ctx context.Context, inboundPlanID string) ([]PlacementOption, error) {
	var result []PlacementOption
	err := s.C.Do(ctx, "getPlacementOptions", map[string]string{"inbound_plan_id": inboundPlanID}, &result)
	return result, err
}

// ConfirmPlacementOption 确认分仓选项
func (s *Service) ConfirmPlacementOption(ctx context.Context, inboundPlanID, placementOptionID string) error {
	return s.C.Do(ctx, "confirmPlacementOption", map[string]string{
		"inbound_plan_id":     inboundPlanID,
		"placement_option_id": placementOptionID,
	}, nil)
}

// CancelPlacementOption 取消生成分仓锁定
func (s *Service) CancelPlacementOption(ctx context.Context, inboundPlanID string) error {
	return s.C.Do(ctx, "cancelPlacementOption", map[string]string{"inbound_plan_id": inboundPlanID}, nil)
}

// --- 包装/装箱 (4个) ---

// GetPackingInformation 查询包装组信息
func (s *Service) GetPackingInformation(ctx context.Context, inboundPlanID string) ([]PackingInformation, error) {
	var result []PackingInformation
	err := s.C.Do(ctx, "getPackingInformation", map[string]string{"inbound_plan_id": inboundPlanID}, &result)
	return result, err
}

// SetPackingInformation 确认装箱
func (s *Service) SetPackingInformation(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "setPackingInformation", req, nil)
}

// DraftPackingInformation 保存装箱
func (s *Service) DraftPackingInformation(ctx context.Context, req *DraftPackingInformationRequest) error {
	return s.C.Do(ctx, "draftPackingInformation", req, nil)
}

// --- 运输选项 (4个) ---

// GenerateTransportationOptions 生成运输选项
func (s *Service) GenerateTransportationOptions(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "generateTransportationOptions", req, nil)
}

// GetTransportationOptions 查询运输选项
func (s *Service) GetTransportationOptions(ctx context.Context, inboundPlanID, shipmentID string) ([]TransportationOption, error) {
	var result []TransportationOption
	err := s.C.Do(ctx, "getTransportationOptions", map[string]string{
		"inbound_plan_id": inboundPlanID,
		"shipment_id":     shipmentID,
	}, &result)
	return result, err
}

// GenerateTransportationOptionsNew 货件预览 - 生成运输选项
func (s *Service) GenerateTransportationOptionsNew(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "generateTransportationOptionsNew", req, nil)
}

// GetTransportationOptionsNew 货件预览 - 查询运输选项
func (s *Service) GetTransportationOptionsNew(ctx context.Context, inboundPlanID, shipmentID string) ([]TransportationOption, error) {
	var result []TransportationOption
	err := s.C.Do(ctx, "getTransportationOptionsNew", map[string]string{
		"inbound_plan_id": inboundPlanID,
		"shipment_id":     shipmentID,
	}, &result)
	return result, err
}

// --- 交付窗口 (2个) ---

// GenerateDeliveryWindowOptions 生成交付窗口
func (s *Service) GenerateDeliveryWindowOptions(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "generateDeliveryWindowOptions", req, nil)
}

// GetDeliveryWindowOptions 查询交付窗口
func (s *Service) GetDeliveryWindowOptions(ctx context.Context, inboundPlanID, shipmentID string) ([]DeliveryWindowOption, error) {
	var result []DeliveryWindowOption
	err := s.C.Do(ctx, "getDeliveryWindowOptions", map[string]string{
		"inbound_plan_id": inboundPlanID,
		"shipment_id":     shipmentID,
	}, &result)
	return result, err
}

// --- 确认/跟踪/面单 (3个) ---

// ConfirmDeliveryWindowAndTransportationOptions 确认运输选项和交货窗口
func (s *Service) ConfirmDeliveryWindowAndTransportationOptions(ctx context.Context, req *ConfirmDeliveryWindowAndTransportationOptionsRequest) error {
	return s.C.Do(ctx, "confirmDeliveryWindowAndTransportationOptions", req, nil)
}

// UpdateShipmentTrackingDetails 更新跟踪信息
func (s *Service) UpdateShipmentTrackingDetails(ctx context.Context, req *UpdateShipmentTrackingDetailsRequest) error {
	return s.C.Do(ctx, "updateShipmentTrackingDetails", req, nil)
}

// GetLabels 获取面单
func (s *Service) GetLabels(ctx context.Context, inboundPlanID, shipmentID string) (*Label, error) {
	var result Label
	err := s.C.Do(ctx, "getLabels", map[string]string{
		"inbound_plan_id": inboundPlanID,
		"shipment_id":     shipmentID,
	}, &result)
	return &result, err
}

// --- 异步任务 (1个) ---

// GetOperation 查询异步任务进度
func (s *Service) GetOperation(ctx context.Context, operationID string) (*Operation, error) {
	var result Operation
	err := s.C.Do(ctx, "getOperation", map[string]string{"operation_id": operationID}, &result)
	return &result, err
}
