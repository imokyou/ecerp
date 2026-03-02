// Package order 提供易仓ERP订单相关API的封装
package order

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 订单服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建订单服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// Order 订单信息
type Order struct {
	OrderCode      string      `json:"order_code"`
	ReferenceCode  string      `json:"reference_code"`
	PlatformCode   string      `json:"platform_code"`
	WarehouseCode  string      `json:"warehouse_code"`
	ShippingMethod string      `json:"shipping_method"`
	TrackingNumber string      `json:"tracking_number"`
	CountryCode    string      `json:"country_code"`
	Currency       string      `json:"currency"`
	Amount         float64     `json:"amount"`
	Status         int         `json:"status"`
	BuyerName      string      `json:"buyer_name"`
	BuyerEmail     string      `json:"buyer_email"`
	BuyerPhone     string      `json:"buyer_phone"`
	Street1        string      `json:"street1"`
	Street2        string      `json:"street2"`
	City           string      `json:"city"`
	Province       string      `json:"province"`
	Zipcode        string      `json:"zipcode"`
	CreateTime     string      `json:"create_time"`
	PayTime        string      `json:"pay_time"`
	Items          []OrderItem `json:"items"`
}

// OrderItem 订单明细
type OrderItem struct {
	SKU           string  `json:"sku"`
	ProductName   string  `json:"product_name"`
	Quantity      int     `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	TotalPrice    float64 `json:"total_price"`
	WarehouseCode string  `json:"warehouse_code"`
}

// ShipOrder 发货单
type ShipOrder struct {
	OrderCode      string `json:"order_code"`
	ShippingMethod string `json:"shipping_method"`
	TrackingNumber string `json:"tracking_number"`
	WarehouseCode  string `json:"warehouse_code"`
	Weight         string `json:"weight"`
	Status         int    `json:"status"`
	ShipDate       string `json:"ship_date"`
}

// ShipTransfer 仓头程发货单
type ShipTransfer struct {
	TransferCode   string `json:"transfer_code"`
	ShippingMethod string `json:"shipping_method"`
	TrackingNumber string `json:"tracking_number"`
	WarehouseCode  string `json:"warehouse_code"`
	Status         int    `json:"status"`
	CreateTime     string `json:"create_time"`
}

// Ticket 工单
type Ticket struct {
	TicketID   int    `json:"ticket_id"`
	OrderCode  string `json:"order_code"`
	Type       string `json:"type"`
	Status     int    `json:"status"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

// ReturnOrder 退件单
type ReturnOrder struct {
	ReturnCode     string `json:"return_code"`
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	Status         int    `json:"status"`
	Reason         string `json:"reason"`
	CreateTime     string `json:"create_time"`
}

// RmaOrder RMA退款单
type RmaOrder struct {
	RmaCode    string  `json:"rma_code"`
	OrderCode  string  `json:"order_code"`
	Amount     float64 `json:"amount"`
	Reason     string  `json:"reason"`
	Status     int     `json:"status"`
	CreateTime string  `json:"create_time"`
}

// RmaReason RMA原因
type RmaReason struct {
	ReasonID   int    `json:"reason_id"`
	ReasonName string `json:"reason_name"`
}

// OrderEventLog 订单流程日志
type OrderEventLog struct {
	LogID      int    `json:"log_id"`
	OrderCode  string `json:"order_code"`
	Event      string `json:"event"`
	Operator   string `json:"operator"`
	CreateTime string `json:"create_time"`
}

// OrderRelation 关联订单
type OrderRelation struct {
	OrderCode    string `json:"order_code"`
	RelatedCode  string `json:"related_code"`
	RelationType string `json:"relation_type"`
}

// SettlementReport 结算报告
type SettlementReport struct {
	ReportID   int     `json:"report_id"`
	OrderCode  string  `json:"order_code"`
	Platform   string  `json:"platform"`
	Amount     float64 `json:"amount"`
	Fee        float64 `json:"fee"`
	Currency   string  `json:"currency"`
	CreateTime string  `json:"create_time"`
}

// InterceptOrder 拦截单
type InterceptOrder struct {
	InterceptCode string `json:"intercept_code"`
	OrderCode     string `json:"order_code"`
	Status        int    `json:"status"`
	Reason        string `json:"reason"`
	CreateTime    string `json:"create_time"`
}

// PickingData 拣货单详情
type PickingData struct {
	PickCode  string `json:"pick_code"`
	OrderCode string `json:"order_code"`
	SKU       string `json:"sku"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`
}

// ShipBatchData 出货总单
type ShipBatchData struct {
	BatchCode      string `json:"batch_code"`
	ShippingMethod string `json:"shipping_method"`
	OrderCount     int    `json:"order_count"`
	TotalWeight    string `json:"total_weight"`
	CreateTime     string `json:"create_time"`
}

// OrderPackage 订单包裹
type OrderPackage struct {
	PackageCode    string `json:"package_code"`
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	Weight         string `json:"weight"`
}

// OrderPackageImage 订单打包图片
type OrderPackageImage struct {
	OrderCode string `json:"order_code"`
	ImageURL  string `json:"image_url"`
}

// StockOrder 缺货订单
type StockOrder struct {
	OrderCode string `json:"order_code"`
	SKU       string `json:"sku"`
	ShortQty  int    `json:"short_qty"`
	Status    int    `json:"status"`
}

// NotFeeOrder 未出账订单
type NotFeeOrder struct {
	OrderCode      string `json:"order_code"`
	ShippingMethod string `json:"shipping_method"`
	WarehouseCode  string `json:"warehouse_code"`
	Status         int    `json:"status"`
}

// FbaDelivery 转FBA配送订单
type FbaDelivery struct {
	OrderCode  string `json:"order_code"`
	ShipmentID string `json:"shipment_id"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// Trail 轨迹
type Trail struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	Event          string `json:"event"`
	Location       string `json:"location"`
	EventTime      string `json:"event_time"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// SyncOrderRequest 建立/更新订单请求
type SyncOrderRequest struct {
	OrderCode      string      `json:"order_code"`
	ReferenceCode  string      `json:"reference_code,omitempty"`
	PlatformCode   string      `json:"platform_code,omitempty"`
	WarehouseCode  string      `json:"warehouse_code"`
	ShippingMethod string      `json:"shipping_method"`
	CountryCode    string      `json:"country_code"`
	Currency       string      `json:"currency,omitempty"`
	Amount         float64     `json:"amount,omitempty"`
	BuyerName      string      `json:"buyer_name"`
	BuyerEmail     string      `json:"buyer_email,omitempty"`
	BuyerPhone     string      `json:"buyer_phone,omitempty"`
	Street1        string      `json:"street1"`
	Street2        string      `json:"street2,omitempty"`
	City           string      `json:"city"`
	Province       string      `json:"province,omitempty"`
	Zipcode        string      `json:"zipcode,omitempty"`
	Items          []OrderItem `json:"items"`
}

// GetOrderListRequest 获取订单列表请求
type GetOrderListRequest struct {
	PageRequest
	OrderCode     string `json:"order_code,omitempty"`
	ReferenceCode string `json:"reference_code,omitempty"`
	Status        int    `json:"status,omitempty"`
	CreateTimeFor string `json:"create_time_for,omitempty"`
	CreateTimeTo  string `json:"create_time_to,omitempty"`
}

// GetShipOrderListRequest 查询销售出库单列表请求
type GetShipOrderListRequest struct {
	PageRequest
	OrderCode   string `json:"order_code,omitempty"`
	ShipDateFor string `json:"ship_date_for,omitempty"`
	ShipDateTo  string `json:"ship_date_to,omitempty"`
}

// GetShipTransferListRequest 查询仓头程发货单请求
type GetShipTransferListRequest struct {
	PageRequest
	TransferCode  string `json:"transfer_code,omitempty"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
}

// HandleTicketRequest 处理工单请求
type HandleTicketRequest struct {
	TicketID int    `json:"ticket_id"`
	Content  string `json:"content,omitempty"`
}

// GetTicketListRequest 获取工单列表请求
type GetTicketListRequest struct {
	PageRequest
	OrderCode string `json:"order_code,omitempty"`
	Status    int    `json:"status,omitempty"`
}

// CreateReturnOrderRequest 创建仓库退件请求
type CreateReturnOrderRequest struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Reason         string `json:"reason,omitempty"`
}

// EditReturnOrderRequest 修改仓库退件请求
type EditReturnOrderRequest struct {
	ReturnCode     string `json:"return_code"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Reason         string `json:"reason,omitempty"`
}

// CreateRmaOrderRequest 创建RMA退款请求
type CreateRmaOrderRequest struct {
	OrderCode string  `json:"order_code"`
	Amount    float64 `json:"amount,omitempty"`
	Reason    string  `json:"reason,omitempty"`
}

// UploadTrackingNoListRequest 批量上传跟踪号请求
type UploadTrackingNoListRequest struct {
	List []TrackingNoItem `json:"list"`
}

// TrackingNoItem 跟踪号项
type TrackingNoItem struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	ShippingMethod string `json:"shipping_method,omitempty"`
}

// CancelOrderRequest 取消订单请求
type CancelOrderRequest struct {
	OrderCode string `json:"order_code"`
	Reason    string `json:"reason,omitempty"`
}

// MarkOrderRequest 标记订单请求
type MarkOrderRequest struct {
	OrderCode string `json:"order_code"`
	MarkType  string `json:"mark_type"`
	MarkValue string `json:"mark_value,omitempty"`
}

// OrderVerifyRequest 审核订单请求
type OrderVerifyRequest struct {
	OrderCode string `json:"order_code"`
}

// GetLabelByCodeRequest 通过跟踪号获取面单请求
type GetLabelByCodeRequest struct {
	TrackingNumber string `json:"tracking_number"`
}

// AddTrailRequest 新增或更新轨迹请求
type AddTrailRequest struct {
	OrderCode      string `json:"order_code"`
	TrackingNumber string `json:"tracking_number"`
	Event          string `json:"event"`
	Location       string `json:"location,omitempty"`
	EventTime      string `json:"event_time,omitempty"`
}

// CreateFbaOrderRequest 创建转FBA订单请求
type CreateFbaOrderRequest struct {
	OrderCode     string `json:"order_code"`
	WarehouseCode string `json:"warehouse_code,omitempty"`
	ShipmentID    string `json:"shipment_id,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (40个接口)
// ════════════════════════════════════════════

// SyncOrder 建立订单/更新订单
func (s *Service) SyncOrder(ctx context.Context, req *SyncOrderRequest) error {
	return s.C.Do(ctx, "syncOrder", req, nil)
}

// GetOrderList 获取订单列表
func (s *Service) GetOrderList(ctx context.Context, req *GetOrderListRequest) ([]Order, error) {
	var result []Order
	err := s.C.Do(ctx, "getOrderList", req, &result)
	return result, err
}

// GetOrderListLite 获取订单列表精简接口
func (s *Service) GetOrderListLite(ctx context.Context, req *GetOrderListRequest) ([]Order, error) {
	var result []Order
	err := s.C.Do(ctx, "getOrderListLite", req, &result)
	return result, err
}

// CancelOrder 取消订单
func (s *Service) CancelOrder(ctx context.Context, req *CancelOrderRequest) error {
	return s.C.Do(ctx, "cancelOrder", req, nil)
}

// OrderVerify 审核订单
func (s *Service) OrderVerify(ctx context.Context, req *OrderVerifyRequest) error {
	return s.C.Do(ctx, "orderVerify", req, nil)
}

// MarkOrder 标记订单
func (s *Service) MarkOrder(ctx context.Context, req *MarkOrderRequest) error {
	return s.C.Do(ctx, "markOrder", req, nil)
}

// GetOrderRelation 获取关联订单
func (s *Service) GetOrderRelation(ctx context.Context, orderCode string) ([]OrderRelation, error) {
	var result []OrderRelation
	err := s.C.Do(ctx, "getOrderRelation", map[string]string{"order_code": orderCode}, &result)
	return result, err
}

// GetOrderEventLog 获取订单流程日志
func (s *Service) GetOrderEventLog(ctx context.Context, orderCode string) ([]OrderEventLog, error) {
	var result []OrderEventLog
	err := s.C.Do(ctx, "getOrderEventLog", map[string]string{"order_code": orderCode}, &result)
	return result, err
}

// GetShipOrderList 查询销售出库单列表
func (s *Service) GetShipOrderList(ctx context.Context, req *GetShipOrderListRequest) ([]ShipOrder, error) {
	var result []ShipOrder
	err := s.C.Do(ctx, "getShipOrderList", req, &result)
	return result, err
}

// GetShipTransferList 查询仓头程发货单列表
func (s *Service) GetShipTransferList(ctx context.Context, req *GetShipTransferListRequest) ([]ShipTransfer, error) {
	var result []ShipTransfer
	err := s.C.Do(ctx, "getShipTransferList", req, &result)
	return result, err
}

// GetLabelByCode 通过跟踪号获取面单
func (s *Service) GetLabelByCode(ctx context.Context, req *GetLabelByCodeRequest) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.C.Do(ctx, "getLabelByCode", req, &result)
	return result, err
}

// HandleTicket 处理工单
func (s *Service) HandleTicket(ctx context.Context, req *HandleTicketRequest) error {
	return s.C.Do(ctx, "handleTicket", req, nil)
}

// GetTicketList 获取工单管理列表
func (s *Service) GetTicketList(ctx context.Context, req *GetTicketListRequest) ([]Ticket, error) {
	var result []Ticket
	err := s.C.Do(ctx, "getTicketList", req, &result)
	return result, err
}

// RmaReturnMarkArrival 仓库退件标记到货
func (s *Service) RmaReturnMarkArrival(ctx context.Context, returnCode string) error {
	return s.C.Do(ctx, "rmaReturnMarkArrival", map[string]string{"return_code": returnCode}, nil)
}

// UploadTrackingNoList 批量上传跟踪号
func (s *Service) UploadTrackingNoList(ctx context.Context, req *UploadTrackingNoListRequest) error {
	return s.C.Do(ctx, "uploadTrackingNoList", req, nil)
}

// GetPickingData 拣货单详情（播种墙）
func (s *Service) GetPickingData(ctx context.Context, pickCode string) ([]PickingData, error) {
	var result []PickingData
	err := s.C.Do(ctx, "getPickingData", map[string]string{"pick_code": pickCode}, &result)
	return result, err
}

// SubmitPack 分拣完成（播种墙）
func (s *Service) SubmitPack(ctx context.Context, pickCode string) error {
	return s.C.Do(ctx, "submitPack", map[string]string{"pick_code": pickCode}, nil)
}

// PickingComplete 下架单拣货回传
func (s *Service) PickingComplete(ctx context.Context, pickCode string) error {
	return s.C.Do(ctx, "pickingComplete", map[string]string{"pick_code": pickCode}, nil)
}

// EditReturnOrder 修改仓库退件
func (s *Service) EditReturnOrder(ctx context.Context, req *EditReturnOrderRequest) error {
	return s.C.Do(ctx, "editReturnOrder", req, nil)
}

// CreateReturnOrder 创建仓库退件
func (s *Service) CreateReturnOrder(ctx context.Context, req *CreateReturnOrderRequest) error {
	return s.C.Do(ctx, "createReturnOrder", req, nil)
}

// GetShipBatchDataForOwms 出货总单
func (s *Service) GetShipBatchDataForOwms(ctx context.Context, req *GetShipOrderListRequest) ([]ShipBatchData, error) {
	var result []ShipBatchData
	err := s.C.Do(ctx, "getShipBatchDataForOwms", req, &result)
	return result, err
}

// GetNotFeeOrders 查询服务商未出账列表
func (s *Service) GetNotFeeOrders(ctx context.Context, req *GetOrderListRequest) ([]NotFeeOrder, error) {
	var result []NotFeeOrder
	err := s.C.Do(ctx, "getNotFeeOrders", req, &result)
	return result, err
}

// CreateRmaOrder 创建RMA退款
func (s *Service) CreateRmaOrder(ctx context.Context, req *CreateRmaOrderRequest) error {
	return s.C.Do(ctx, "createRmaOrder", req, nil)
}

// GetRmaReason 获取RMA原因
func (s *Service) GetRmaReason(ctx context.Context) ([]RmaReason, error) {
	var result []RmaReason
	err := s.C.Do(ctx, "getRmaReason", nil, &result)
	return result, err
}

// UpdateRmaReturnTrackingNo 更新退件单跟踪号
func (s *Service) UpdateRmaReturnTrackingNo(ctx context.Context, returnCode, trackingNo string) error {
	return s.C.Do(ctx, "updateRmaReturnTrackingNo", map[string]string{
		"return_code":     returnCode,
		"tracking_number": trackingNo,
	}, nil)
}

// GetOrders 查询仓配订单信息
func (s *Service) GetOrders(ctx context.Context, req *GetOrderListRequest) ([]Order, error) {
	var result []Order
	err := s.C.Do(ctx, "getOrders", req, &result)
	return result, err
}

// GetAmazonOriginalReportList 获取结算报告原始数据列表
func (s *Service) GetAmazonOriginalReportList(ctx context.Context, req *GetOrderListRequest) ([]SettlementReport, error) {
	var result []SettlementReport
	err := s.C.Do(ctx, "getamazonOriginalReportList", req, &result)
	return result, err
}

// GetAmazonReportList 获取结算报告列表
func (s *Service) GetAmazonReportList(ctx context.Context, req *GetOrderListRequest) ([]SettlementReport, error) {
	var result []SettlementReport
	err := s.C.Do(ctx, "getAmazonReportList", req, &result)
	return result, err
}

// GetRmaRefaList 退件重发列表
func (s *Service) GetRmaRefaList(ctx context.Context, req *GetOrderListRequest) ([]RmaOrder, error) {
	var result []RmaOrder
	err := s.C.Do(ctx, "getRmaRefaList", req, &result)
	return result, err
}

// GetStockOrderList 获取缺货订单列表
func (s *Service) GetStockOrderList(ctx context.Context, req *GetOrderListRequest) ([]StockOrder, error) {
	var result []StockOrder
	err := s.C.Do(ctx, "getStockOrderList", req, &result)
	return result, err
}

// AddTrail 新增或更新轨迹
func (s *Service) AddTrail(ctx context.Context, req *AddTrailRequest) error {
	return s.C.Do(ctx, "addTrail", req, nil)
}

// GetRmaRefundList 退款订单列表
func (s *Service) GetRmaRefundList(ctx context.Context, req *GetOrderListRequest) ([]RmaOrder, error) {
	var result []RmaOrder
	err := s.C.Do(ctx, "getRmaRefundList", req, &result)
	return result, err
}

// GetRmaReturnList 退件列表
func (s *Service) GetRmaReturnList(ctx context.Context, req *GetOrderListRequest) ([]ReturnOrder, error) {
	var result []ReturnOrder
	err := s.C.Do(ctx, "getRmaReturnList", req, &result)
	return result, err
}

// GetInterceptorList 查询拦截单列表
func (s *Service) GetInterceptorList(ctx context.Context, req *GetOrderListRequest) ([]InterceptOrder, error) {
	var result []InterceptOrder
	err := s.C.Do(ctx, "getInterceptorList", req, &result)
	return result, err
}

// GetOrderPackageImage 获取订单打包图片
func (s *Service) GetOrderPackageImage(ctx context.Context, orderCode string) ([]OrderPackageImage, error) {
	var result []OrderPackageImage
	err := s.C.Do(ctx, "getOrderPackageImage", map[string]string{"order_code": orderCode}, &result)
	return result, err
}

// GetPickList 下架订单查询
func (s *Service) GetPickList(ctx context.Context, req *GetOrderListRequest) ([]PickingData, error) {
	var result []PickingData
	err := s.C.Do(ctx, "getPickList", req, &result)
	return result, err
}

// GetOrderPackageList 订单拆分包裹
func (s *Service) GetOrderPackageList(ctx context.Context, orderCode string) ([]OrderPackage, error) {
	var result []OrderPackage
	err := s.C.Do(ctx, "getOrderPackageList", map[string]string{"order_code": orderCode}, &result)
	return result, err
}

// GetOrderInfo 获取仓库订单详情
func (s *Service) GetOrderInfo(ctx context.Context, orderCode string) (*Order, error) {
	var result Order
	err := s.C.Do(ctx, "getOrderInfo", map[string]string{"order_code": orderCode}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateFbaOrder 创建转FBA订单
func (s *Service) CreateFbaOrder(ctx context.Context, req *CreateFbaOrderRequest) error {
	return s.C.Do(ctx, "createFbaOrder", req, nil)
}

// GetTransferFbaDelivery 转FBA配送订单查询
func (s *Service) GetTransferFbaDelivery(ctx context.Context, req *GetOrderListRequest) ([]FbaDelivery, error) {
	var result []FbaDelivery
	err := s.C.Do(ctx, "getTransferFbaDelivery", req, &result)
	return result, err
}
