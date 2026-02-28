// Package product 提供易仓ERP产品相关API的封装
package product

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 产品服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建产品服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// Product 产品信息
type Product struct {
	ProductSKU    string   `json:"product_sku"`
	ProductName   string   `json:"product_name"`
	ProductNameEn string   `json:"product_name_en"`
	CategoryID    int      `json:"category_id"`
	BrandID       int      `json:"brand_id"`
	Weight        float64  `json:"weight"`
	Length        float64  `json:"length"`
	Width         float64  `json:"width"`
	Height        float64  `json:"height"`
	DeclaredValue float64  `json:"declared_value"`
	Status        int      `json:"status"`
	SaleStatus    int      `json:"sale_status"`
	ProductLevel  int      `json:"product_level"`
	ImageURL      string   `json:"image_url"`
	CreateTime    string   `json:"create_time"`
	UpdateTime    string   `json:"update_time"`
	Tags          []string `json:"tags,omitempty"`
}

// ProductTrouble 产品问题
type ProductTrouble struct {
	TroubleID   int    `json:"trouble_id"`
	ProductSKU  string `json:"product_sku"`
	TroubleType int    `json:"trouble_type"`
	Content     string `json:"content"`
	Status      int    `json:"status"`
	CreateTime  string `json:"create_time"`
}

// ListingAccess Listing权限
type ListingAccess struct {
	ListingID  int    `json:"listing_id"`
	ProductSKU string `json:"product_sku"`
	Platform   string `json:"platform"`
	AccountID  int    `json:"account_id"`
	Status     int    `json:"status"`
}

// ProductCustomsAttribute 产品海关属性
type ProductCustomsAttribute struct {
	ProductSKU    string  `json:"product_sku"`
	HsCode        string  `json:"hs_code"`
	DeclaredValue float64 `json:"product_declared_value"`
	OverseatypeEn string  `json:"pd_oversea_type_en"`
	Material      string  `json:"material"`
	MaterialEn    string  `json:"material_en"`
	Use           string  `json:"use"`
	UseEn         string  `json:"use_en"`
}

// ProductExpressAttribute 产品物流属性
type ProductExpressAttribute struct {
	ProductSKU    string  `json:"product_sku"`
	IsBattery     int     `json:"is_battery"`
	IsLiquid      int     `json:"is_liquid"`
	IsPowder      int     `json:"is_powder"`
	IsMagnetic    int     `json:"is_magnetic"`
	ExpressWeight float64 `json:"express_weight"`
}

// ProductBarcodeMap 产品条码
type ProductBarcodeMap struct {
	BarcodeID  int    `json:"barcode_id"`
	ProductSKU string `json:"product_sku"`
	Barcode    string `json:"barcode"`
	Type       string `json:"type"`
}

// ProductBox 产品箱规
type ProductBox struct {
	BoxID      int     `json:"box_id"`
	ProductSKU string  `json:"product_sku"`
	Length     float64 `json:"length"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	Quantity   int     `json:"quantity"`
}

// ProductParent 产品款式
type ProductParent struct {
	ParentID   int    `json:"parent_id"`
	ParentSKU  string `json:"parent_sku"`
	ParentName string `json:"parent_name"`
}

// ProductCategory 产品品类
type ProductCategory struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     int    `json:"parent_id"`
	Level        int    `json:"level"`
}

// ProductBrand 产品品牌
type ProductBrand struct {
	BrandID   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

// ProductColor 产品颜色
type ProductColor struct {
	ColorID   int    `json:"color_id"`
	ColorName string `json:"color_name"`
}

// ProductSize 产品尺寸
type ProductSize struct {
	SizeID   int    `json:"size_id"`
	SizeName string `json:"size_name"`
}

// ProductLevel 产品等级
type ProductLevel struct {
	LevelID   int    `json:"level_id"`
	LevelName string `json:"level_name"`
}

// SaleStatus 销售状态
type SaleStatus struct {
	StatusID   int    `json:"status_id"`
	StatusName string `json:"status_name"`
}

// SupplierInfo 供应商信息
type SupplierInfo struct {
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	Contact      string `json:"contact"`
}

// SkuRelation SKU关系映射
type SkuRelation struct {
	SKU        string `json:"sku"`
	RelatedSKU string `json:"related_sku"`
	Type       string `json:"type"`
}

// EnquiryItem 询价管理
type EnquiryItem struct {
	EnquiryID  int    `json:"enquiry_id"`
	ProductSKU string `json:"product_sku"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// WarehouseCost 产品成本
type WarehouseCost struct {
	ProductSKU    string  `json:"product_sku"`
	WarehouseCode string  `json:"warehouse_code"`
	Cost          float64 `json:"cost"`
}

// UserCategory 自定义分类属性
type UserCategory struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// ProductTort 产品侵权
type ProductTort struct {
	TortID   int    `json:"tort_id"`
	TortName string `json:"tort_name"`
}

// QcTemplate 质检模板
type QcTemplate struct {
	TemplateID   int    `json:"template_id"`
	TemplateName string `json:"template_name"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// SyncProductRequest 创建、编辑产品请求
type SyncProductRequest struct {
	ProductSKU    string   `json:"product_sku"`
	ProductName   string   `json:"product_name"`
	ProductNameEn string   `json:"product_name_en,omitempty"`
	CategoryID    int      `json:"category_id,omitempty"`
	BrandID       int      `json:"brand_id,omitempty"`
	Weight        float64  `json:"weight,omitempty"`
	Length        float64  `json:"length,omitempty"`
	Width         float64  `json:"width,omitempty"`
	Height        float64  `json:"height,omitempty"`
	DeclaredValue float64  `json:"declared_value,omitempty"`
	ImageURL      string   `json:"image_url,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

// SyncBatchProductRequest 批量创建、编辑产品请求
type SyncBatchProductRequest struct {
	Products []SyncProductRequest `json:"products"`
}

// SyncBatchProductAttrRequest 批量更新产品属性请求
type SyncBatchProductAttrRequest struct {
	ProductSKUArr []string `json:"product_sku_arr"`
	HsCode        string   `json:"hs_code,omitempty"`
	DeclaredValue float64  `json:"product_declared_value,omitempty"`
	Material      string   `json:"material,omitempty"`
	MaterialEn    string   `json:"material_en,omitempty"`
	Use           string   `json:"use,omitempty"`
	UseEn         string   `json:"use_en,omitempty"`
}

// GetProductListRequest 获取产品列表请求
type GetProductListRequest struct {
	PageRequest
	ProductSKU string `json:"product_sku,omitempty"`
	Status     int    `json:"status,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (46个接口)
// ════════════════════════════════════════════

// --- 产品 CRUD ---

// SyncProduct 创建、编辑产品
func (s *Service) SyncProduct(ctx context.Context, req *SyncProductRequest) error {
	return s.C.Do(ctx, "syncProduct", req, nil)
}

// SyncBatchProduct 批量创建、编辑产品
func (s *Service) SyncBatchProduct(ctx context.Context, req *SyncBatchProductRequest) error {
	return s.C.Do(ctx, "syncBatchProduct", req, nil)
}

// SyncBatchProductAttr 批量更新产品属性
func (s *Service) SyncBatchProductAttr(ctx context.Context, req *SyncBatchProductAttrRequest) error {
	return s.C.Do(ctx, "syncBatchProductAttr", req, nil)
}

// GetProductBySku 获取单个产品信息
func (s *Service) GetProductBySku(ctx context.Context, sku string) (*Product, error) {
	var result Product
	err := s.C.Do(ctx, "getProductBySku", map[string]string{"product_sku": sku}, &result)
	return &result, err
}

// GetWmsProductList 获取产品列表
func (s *Service) GetWmsProductList(ctx context.Context, req *GetProductListRequest) ([]Product, error) {
	var result []Product
	err := s.C.Do(ctx, "getWmsProductList", req, &result)
	return result, err
}

// UpdateProductForWms 更新产品（WMS）
func (s *Service) UpdateProductForWms(ctx context.Context, req *SyncProductRequest) error {
	return s.C.Do(ctx, "updateProductForWms", req, nil)
}

// UpdateProductImages 批量更新产品主图
func (s *Service) UpdateProductImages(ctx context.Context, items []map[string]string) error {
	return s.C.Do(ctx, "updateProductImages", items, nil)
}

// UpdateElProduct 更新易链产品状态 (beta)
func (s *Service) UpdateElProduct(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "updateElProduct", req, nil)
}

// --- 销售状态与等级 ---

// GetSaleStatus 获取产品销售状态
func (s *Service) GetSaleStatus(ctx context.Context) ([]SaleStatus, error) {
	var result []SaleStatus
	err := s.C.Do(ctx, "getSaleStatus", nil, &result)
	return result, err
}

// --- SKU关系映射 ---

// GetSkuRelation 查询SKU关系映射
func (s *Service) GetSkuRelation(ctx context.Context, sku string) ([]SkuRelation, error) {
	var result []SkuRelation
	err := s.C.Do(ctx, "getSkuRelation", map[string]string{"sku": sku}, &result)
	return result, err
}

// ModifySkuRelation 编辑SKU关系映射
func (s *Service) ModifySkuRelation(ctx context.Context, req *SkuRelation) error {
	return s.C.Do(ctx, "modifySkuRelation", req, nil)
}

// --- 产品问题 ---

// AddProductTrouble 创建产品问题
func (s *Service) AddProductTrouble(ctx context.Context, req *ProductTrouble) error {
	return s.C.Do(ctx, "addProductTrouble", req, nil)
}

// ProductTroubleDel 删除产品问题
func (s *Service) ProductTroubleDel(ctx context.Context, troubleID int) error {
	return s.C.Do(ctx, "productTroubleDel", map[string]int{"trouble_id": troubleID}, nil)
}

// SyncProductTrouble 产品问题同步 (注: 实际方法名 getProductTroubleType 在基础数据中也有)
func (s *Service) GetProductTroubleType(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getProductTroubleType", nil, &result)
	return result, err
}

// GetProductTrouble 产品问题明细
func (s *Service) GetProductTrouble(ctx context.Context, req *PageRequest) ([]ProductTrouble, error) {
	var result []ProductTrouble
	err := s.C.Do(ctx, "getProductTrouble", req, &result)
	return result, err
}

// --- Listing和海关 ---

// GetListingAccess Listing权限列表
func (s *Service) GetListingAccess(ctx context.Context, req *PageRequest) ([]ListingAccess, error) {
	var result []ListingAccess
	err := s.C.Do(ctx, "getListingAccess", req, &result)
	return result, err
}

// GetProductCustomsAttribute 获取产品海关属性
func (s *Service) GetProductCustomsAttribute(ctx context.Context, sku string) (*ProductCustomsAttribute, error) {
	var result ProductCustomsAttribute
	err := s.C.Do(ctx, "getProductCustomsAttribute", map[string]string{"product_sku": sku}, &result)
	return &result, err
}

// --- 产品款式 ---

// SyncProductParent 创建、编辑产品款式
func (s *Service) SyncProductParent(ctx context.Context, req *ProductParent) error {
	return s.C.Do(ctx, "syncProductParent", req, nil)
}

// GetProductParent 获取产品款式数据
func (s *Service) GetProductParent(ctx context.Context) ([]ProductParent, error) {
	var result []ProductParent
	err := s.C.Do(ctx, "getProductParent", nil, &result)
	return result, err
}

// --- 物流属性 ---

// SyncProductExpressAttribute 创建/编辑产品物流属性
func (s *Service) SyncProductExpressAttribute(ctx context.Context, req *ProductExpressAttribute) error {
	return s.C.Do(ctx, "syncProductExpressAttribute", req, nil)
}

// GetProductExpressAttribute 获取产品物流属性信息
func (s *Service) GetProductExpressAttribute(ctx context.Context, sku string) (*ProductExpressAttribute, error) {
	var result ProductExpressAttribute
	err := s.C.Do(ctx, "getProductExpressAttribute", map[string]string{"product_sku": sku}, &result)
	return &result, err
}

// --- 颜色/尺寸 ---

// SyncColour 创建/编辑产品颜色
func (s *Service) SyncColour(ctx context.Context, req *ProductColor) error {
	return s.C.Do(ctx, "syncColour", req, nil)
}

// SyncSize 创建/编辑产品尺寸
func (s *Service) SyncSize(ctx context.Context, req *ProductSize) error {
	return s.C.Do(ctx, "syncSize", req, nil)
}

// GetProductColor 获取产品颜色数据
func (s *Service) GetProductColor(ctx context.Context) ([]ProductColor, error) {
	var result []ProductColor
	err := s.C.Do(ctx, "getProductColor", nil, &result)
	return result, err
}

// GetProductSize 获取产品尺寸数据
func (s *Service) GetProductSize(ctx context.Context) ([]ProductSize, error) {
	var result []ProductSize
	err := s.C.Do(ctx, "getProductSize", nil, &result)
	return result, err
}

// --- 条码管理 ---

// GetProductBarcodeMapList 获取产品条码管理
func (s *Service) GetProductBarcodeMapList(ctx context.Context, req *PageRequest) ([]ProductBarcodeMap, error) {
	var result []ProductBarcodeMap
	err := s.C.Do(ctx, "getProductBarcodeMapList", req, &result)
	return result, err
}

// SyncProductBarCodeMap 新增、编辑产品条码
func (s *Service) SyncProductBarCodeMap(ctx context.Context, req *ProductBarcodeMap) error {
	return s.C.Do(ctx, "syncProductBarCodeMap", req, nil)
}

// BatchAddProductBarCodeMap 批量添加产品条码
func (s *Service) BatchAddProductBarCodeMap(ctx context.Context, items []ProductBarcodeMap) error {
	return s.C.Do(ctx, "batchAddProductBarCodeMap", items, nil)
}

// --- 箱规管理 ---

// SyncProductBoxes 新增、编辑产品箱规
func (s *Service) SyncProductBoxes(ctx context.Context, req *ProductBox) error {
	return s.C.Do(ctx, "syncProductBoxes", req, nil)
}

// BatchAddProductBoxes 批量添加产品箱规
func (s *Service) BatchAddProductBoxes(ctx context.Context, items []ProductBox) error {
	return s.C.Do(ctx, "batchAddProductBoxes", items, nil)
}

// BatchEditProductBoxes 批量编辑产品箱规
func (s *Service) BatchEditProductBoxes(ctx context.Context, items []ProductBox) error {
	return s.C.Do(ctx, "batchEditProductBoxes", items, nil)
}

// --- 海外仓组合产品 ---

// SyncOverseasWarehouseCombinationProduct 创建、编辑海外仓组合产品
func (s *Service) SyncOverseasWarehouseCombinationProduct(ctx context.Context, req map[string]interface{}) error {
	return s.C.Do(ctx, "syncOverseasWarehouseCombinationProduct", req, nil)
}

// --- 品类管理 ---

// CategotyList 获取产品品类列表 (注意: 文档拼写为 categoty)
func (s *Service) CategotyList(ctx context.Context, req *PageRequest) ([]ProductCategory, error) {
	var result []ProductCategory
	err := s.C.Do(ctx, "categotyList", req, &result) // 注意: 文档原始拼写为 categoty (缺少 'r')
	return result, err
}

// GetProductCategory 获取产品品类列表（新）
func (s *Service) GetProductCategory(ctx context.Context, req *PageRequest) ([]ProductCategory, error) {
	var result []ProductCategory
	err := s.C.Do(ctx, "getProductCategory", req, &result)
	return result, err
}

// EditCategory 创建、编辑产品品类
func (s *Service) EditCategory(ctx context.Context, req *ProductCategory) error {
	return s.C.Do(ctx, "editCategory", req, nil)
}

// --- 产品等级/品牌 ---

// GetProductLevel 获取产品等级数据
func (s *Service) GetProductLevel(ctx context.Context) ([]ProductLevel, error) {
	var result []ProductLevel
	err := s.C.Do(ctx, "getProductLevel", nil, &result)
	return result, err
}

// GetBrand 获取产品品牌
func (s *Service) GetBrand(ctx context.Context) ([]ProductBrand, error) {
	var result []ProductBrand
	err := s.C.Do(ctx, "getBrand", nil, &result)
	return result, err
}

// --- 供应商/用户/分类 ---

// GetAllSupplier 获取供应商代码和联系方式
func (s *Service) GetAllSupplier(ctx context.Context) ([]SupplierInfo, error) {
	var result []SupplierInfo
	err := s.C.Do(ctx, "getAllSupplier", nil, &result)
	return result, err
}

// GetProductUserCategory 获取自定义产品分类
func (s *Service) GetProductUserCategory(ctx context.Context) ([]UserCategory, error) {
	var result []UserCategory
	err := s.C.Do(ctx, "getProductUserCategory", nil, &result)
	return result, err
}

// GetPuc 获取自定义分类属性
func (s *Service) GetPuc(ctx context.Context) ([]UserCategory, error) {
	var result []UserCategory
	err := s.C.Do(ctx, "getPuc", nil, &result)
	return result, err
}

// GetUserAll 获取用户部门/小组数据
func (s *Service) GetUserAll(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "getUserAll", nil, &result)
	return result, err
}

// GetProductTortAll 获取产品侵权列表
func (s *Service) GetProductTortAll(ctx context.Context) ([]ProductTort, error) {
	var result []ProductTort
	err := s.C.Do(ctx, "getProductTortAll", nil, &result)
	return result, err
}

// GetQcTemplateAll 获取质检模板列表
func (s *Service) GetQcTemplateAll(ctx context.Context) ([]QcTemplate, error) {
	var result []QcTemplate
	err := s.C.Do(ctx, "getQcTemplateAll", nil, &result)
	return result, err
}

// GetProductCategoryBase 获取产品品类基础信息
func (s *Service) GetProductCategoryBase(ctx context.Context) ([]ProductCategory, error) {
	var result []ProductCategory
	err := s.C.Do(ctx, "getProductCategoryBase", nil, &result)
	return result, err
}

// --- 询价/成本 ---

// EnquiryList 询价管理列表
func (s *Service) EnquiryList(ctx context.Context, req *PageRequest) ([]EnquiryItem, error) {
	var result []EnquiryItem
	err := s.C.Do(ctx, "enquiryList", req, &result)
	return result, err
}

// GetWarehouseCostList 获取产品成本列表
func (s *Service) GetWarehouseCostList(ctx context.Context, req *PageRequest) ([]WarehouseCost, error) {
	var result []WarehouseCost
	err := s.C.Do(ctx, "getWarehouseCostList", req, &result)
	return result, err
}
