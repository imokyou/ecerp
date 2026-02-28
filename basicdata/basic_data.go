// Package basicdata 提供易仓ERP基础数据相关API的封装
package basicdata

import (
	"context"

	"github.com/imokyou/ecerp"
)

// Service 基础数据服务
type Service struct {
	C ecerp.Caller
}

// NewService 创建基础数据服务
func NewService(c ecerp.Caller) *Service {
	return &Service{C: c}
}

// ════════════════════════════════════════════
// 数据结构
// ════════════════════════════════════════════

// Warehouse 仓库信息
type Warehouse struct {
	WarehouseID   int    `json:"warehouse_id"`
	WarehouseCode string `json:"warehouse_code"`
	WarehouseName string `json:"warehouse_name"`
	WarehouseType int    `json:"warehouse_type"`
	Country       string `json:"country"`
	Status        int    `json:"status"`
}

// ShippingMethod 运输方式
type ShippingMethod struct {
	ShippingMethodID   int    `json:"shipping_method_id"`
	ShippingMethodCode string `json:"shipping_method_code"`
	ShippingMethodName string `json:"shipping_method_name"`
	ShippingMethodType int    `json:"shipping_method_type"`
	Status             int    `json:"status"`
}

// Country 国家信息
type Country struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	CountryEn   string `json:"country_en"`
}

// Currency 币种信息
type Currency struct {
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
}

// User 用户
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	RealName string `json:"real_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

// Organization 组织机构
type Organization struct {
	OrgID      int    `json:"org_id"`
	OrgName    string `json:"org_name"`
	ParentID   int    `json:"parent_id"`
	ParentName string `json:"parent_name"`
	Level      int    `json:"level"`
}

// BusinessLicense 营业执照
type BusinessLicense struct {
	BlID          int    `json:"bl_id"`
	CompanyName   string `json:"company_name"`
	LicenseNo     string `json:"license_no"`
	LegalPerson   string `json:"legal_person"`
	RegisterAddr  string `json:"register_addr"`
	BusinessScope string `json:"business_scope"`
	ValidFrom     string `json:"valid_from"`
	ValidTo       string `json:"valid_to"`
	Status        int    `json:"status"`
	CreateTime    string `json:"create_time"`
}

// TicketBusinessType 工单业务类型
type TicketBusinessType struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

// ShipAddressBook 物流地址
type ShipAddressBook struct {
	AddressID   int    `json:"address_id"`
	AddressName string `json:"address_name"`
	Contact     string `json:"contact"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Zipcode     string `json:"zipcode"`
}

// PdaRight PDA权限
type PdaRight struct {
	RightID   int    `json:"right_id"`
	RightName string `json:"right_name"`
	RightCode string `json:"right_code"`
	Status    int    `json:"status"`
}

// RolePermission 角色权限
type RolePermission struct {
	RoleID     int    `json:"role_id"`
	RoleName   string `json:"role_name"`
	Permission string `json:"permission"`
}

// ProductLevel 产品等级
type ProductLevel struct {
	LevelID   int    `json:"level_id"`
	LevelName string `json:"level_name"`
}

// UserCategory 自定义分类
type UserCategory struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     int    `json:"parent_id"`
}

// QcOption 质检项
type QcOption struct {
	OptionID   int    `json:"option_id"`
	OptionName string `json:"option_name"`
}

// PackageMaterial 包材
type PackageMaterial struct {
	PackageID   int     `json:"package_id"`
	PackageName string  `json:"package_name"`
	Length      float64 `json:"length"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Weight      float64 `json:"weight"`
}

// Supplier 供应商
type Supplier struct {
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	Contact      string `json:"contact"`
	Phone        string `json:"phone"`
	Status       int    `json:"status"`
}

// ProductParent 产品款式
type ProductParent struct {
	ParentID   int    `json:"parent_id"`
	ParentSKU  string `json:"parent_sku"`
	ParentName string `json:"parent_name"`
}

// WarehouseShipping 仓库运输方式映射
type WarehouseShipping struct {
	WarehouseCode      string `json:"warehouse_code"`
	WarehouseName      string `json:"warehouse_name"`
	ShippingMethodCode string `json:"shipping_method_code"`
	ShippingMethodName string `json:"shipping_method_name"`
}

// ProductTroubleType 产品问题类型
type ProductTroubleType struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

// ProductCategoryBase 产品品类
type ProductCategoryBase struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     int    `json:"parent_id"`
	Level        int    `json:"level"`
}

// PaymentAccount 支付账号
type PaymentAccount struct {
	AccountID   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	AccountType string `json:"account_type"`
	BankName    string `json:"bank_name"`
	AccountNo   string `json:"account_no"`
}

// PurchaseShipper 采购承运商
type PurchaseShipper struct {
	ShipperID   int    `json:"shipper_id"`
	ShipperName string `json:"shipper_name"`
	ShipperCode string `json:"shipper_code"`
}

// ════════════════════════════════════════════
// 请求参数
// ════════════════════════════════════════════

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// BusinessLicenseBindAccountRequest 营业执照绑定店铺请求
type BusinessLicenseBindAccountRequest struct {
	BlID       int   `json:"bl_id"`
	AccountIDs []int `json:"account_ids"`
}

// BusinessLicenseBindShippingMethodRequest 营业执照绑定运输方式请求
type BusinessLicenseBindShippingMethodRequest struct {
	BlID              int   `json:"bl_id"`
	ShippingMethodIDs []int `json:"shipping_method_ids"`
}

// BusinessLicenseCancelBindRequest 营业执照取消绑定请求
type BusinessLicenseCancelBindRequest struct {
	BlID int    `json:"bl_id"`
	Type string `json:"type"` // account or shipping_method
	IDs  []int  `json:"ids"`
}

// EditBusinessLicenseRequest 新增/编辑营业执照请求
type EditBusinessLicenseRequest struct {
	BlID          int    `json:"bl_id,omitempty"` // 0 为新增
	CompanyName   string `json:"company_name"`
	LicenseNo     string `json:"license_no"`
	LegalPerson   string `json:"legal_person,omitempty"`
	RegisterAddr  string `json:"register_addr,omitempty"`
	BusinessScope string `json:"business_scope,omitempty"`
	ValidFrom     string `json:"valid_from,omitempty"`
	ValidTo       string `json:"valid_to,omitempty"`
}

// GetBusinessLicenseListRequest 营业执照列表请求
type GetBusinessLicenseListRequest struct {
	PageRequest
	CompanyName string `json:"company_name,omitempty"`
	LicenseNo   string `json:"license_no,omitempty"`
}

// CreateProductForWmsRequest 创建产品(WMS)请求
type CreateProductForWmsRequest struct {
	SKU         string  `json:"sku"`
	ProductName string  `json:"product_name"`
	Weight      float64 `json:"weight,omitempty"`
	Length      float64 `json:"length,omitempty"`
	Width       float64 `json:"width,omitempty"`
	Height      float64 `json:"height,omitempty"`
}

// ════════════════════════════════════════════
// 服务方法 (34个接口)
// ════════════════════════════════════════════

// --- 仓库、运输、国家、币种 ---

// GetWarehouses 获取所有仓库
func (s *Service) GetAllWarehouse(ctx context.Context) ([]Warehouse, error) {
	var result []Warehouse
	err := s.C.Do(ctx, "getWarehouses", nil, &result)
	return result, err
}

// GetWarehouseForOrder 获取所有仓库（订单相关）
func (s *Service) GetWarehouseForOrder(ctx context.Context) ([]Warehouse, error) {
	var result []Warehouse
	err := s.C.Do(ctx, "getWarehouseForOrder", nil, &result)
	return result, err
}

// GetShippingMethod 获取运输方式
func (s *Service) GetShippingMethod(ctx context.Context) ([]ShippingMethod, error) {
	var result []ShippingMethod
	err := s.C.Do(ctx, "getShippingMethod", nil, &result)
	return result, err
}

// GetShippingMethodForOrder 获取运输方式（订单相关）
func (s *Service) GetShippingMethodForOrder(ctx context.Context) ([]ShippingMethod, error) {
	var result []ShippingMethod
	err := s.C.Do(ctx, "getShippingMethodForOrder", nil, &result)
	return result, err
}

// GetWarehouseShippingForOrder 获取仓库与运输方式映射（订单相关）
func (s *Service) GetWarehouseShippingForOrder(ctx context.Context) ([]WarehouseShipping, error) {
	var result []WarehouseShipping
	err := s.C.Do(ctx, "getWarehouseShippingForOrder", nil, &result)
	return result, err
}

// GetCountry 获取国家列表
func (s *Service) GetCountry(ctx context.Context) ([]Country, error) {
	var result []Country
	err := s.C.Do(ctx, "getCountry", nil, &result)
	return result, err
}

// GetCurrency 获取币种
func (s *Service) GetCurrency(ctx context.Context) ([]Currency, error) {
	var result []Currency
	err := s.C.Do(ctx, "getCurrency", nil, &result)
	return result, err
}

// --- 用户、组织 ---

// GetUser 获取用户列表
func (s *Service) GetUser(ctx context.Context) ([]User, error) {
	var result []User
	err := s.C.Do(ctx, "getUser", nil, &result)
	return result, err
}

// GetUserOrganization 获取组织机构信息
func (s *Service) GetUserOrganization(ctx context.Context) ([]Organization, error) {
	var result []Organization
	err := s.C.Do(ctx, "getUserOrganization", nil, &result)
	return result, err
}

// GetUserOrganizationAll 获取机构组织（全部）
func (s *Service) GetUserOrganizationAll(ctx context.Context) ([]Organization, error) {
	var result []Organization
	err := s.C.Do(ctx, "getUserOrganizationAll", nil, &result)
	return result, err
}

// GetUserOrganizationData 获取用户部门列表
func (s *Service) GetUserOrganizationData(ctx context.Context) ([]Organization, error) {
	var result []Organization
	err := s.C.Do(ctx, "getUserOrganizationData", nil, &result)
	return result, err
}

// --- 营业执照管理 (8个) ---

// BusinessLicenseBindAccount 营业执照-绑定店铺
func (s *Service) BusinessLicenseBindAccount(ctx context.Context, req *BusinessLicenseBindAccountRequest) error {
	return s.C.Do(ctx, "businessLicenseBindAccount", req, nil)
}

// BusinessLicenseBindShippingMethod 营业执照-绑定运输方式
func (s *Service) BusinessLicenseBindShippingMethod(ctx context.Context, req *BusinessLicenseBindShippingMethodRequest) error {
	return s.C.Do(ctx, "businessLicenseBindShippingMethod", req, nil)
}

// BusinessLicenseCancelBind 营业执照-取消绑定店铺或运输方式
func (s *Service) BusinessLicenseCancelBind(ctx context.Context, req *BusinessLicenseCancelBindRequest) error {
	return s.C.Do(ctx, "businessLicenseCancelBind", req, nil)
}

// BusinessLicenseShippingMethod 营业执照-已绑定运输方式
func (s *Service) BusinessLicenseShippingMethod(ctx context.Context, blID int) ([]ShippingMethod, error) {
	var result []ShippingMethod
	err := s.C.Do(ctx, "businessLicenseShippingMethod", map[string]int{"bl_id": blID}, &result)
	return result, err
}

// BusinessLicenseAccount 营业执照-已绑定店铺
func (s *Service) BusinessLicenseAccount(ctx context.Context, blID int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.C.Do(ctx, "businessLicenseAccount", map[string]int{"bl_id": blID}, &result)
	return result, err
}

// DeleteBusinessLicense 营业执照-删除
func (s *Service) DeleteBusinessLicense(ctx context.Context, blID int) error {
	return s.C.Do(ctx, "deleteBusinessLicense", map[string]int{"bl_id": blID}, nil)
}

// EditBusinessLicense 营业执照-新增和编辑
func (s *Service) EditBusinessLicense(ctx context.Context, req *EditBusinessLicenseRequest) error {
	return s.C.Do(ctx, "editBusinessLicense", req, nil)
}

// GetBusinessLicenseList 营业执照-列表
func (s *Service) GetBusinessLicenseList(ctx context.Context, req *GetBusinessLicenseListRequest) ([]BusinessLicense, error) {
	var result []BusinessLicense
	err := s.C.Do(ctx, "getBusinessLicenseList", req, &result)
	return result, err
}

// --- 其他基础数据 ---

// GetTicketBusinessType 获取工单管理业务类型
func (s *Service) GetTicketBusinessType(ctx context.Context) ([]TicketBusinessType, error) {
	var result []TicketBusinessType
	err := s.C.Do(ctx, "getTicketBusinessType", nil, &result)
	return result, err
}

// GetShipAddressBooks 获取物流地址
func (s *Service) GetShipAddressBooks(ctx context.Context) ([]ShipAddressBook, error) {
	var result []ShipAddressBook
	err := s.C.Do(ctx, "getShipAddressBooks", nil, &result)
	return result, err
}

// GetPdaRightControlAll 获取PDA权限数据
func (s *Service) GetPdaRightControlAll(ctx context.Context) ([]PdaRight, error) {
	var result []PdaRight
	err := s.C.Do(ctx, "getPdaRightControlAll", nil, &result)
	return result, err
}

// GetRolePermissionAll 获取角色业务权限基础数据
func (s *Service) GetRolePermissionAll(ctx context.Context) ([]RolePermission, error) {
	var result []RolePermission
	err := s.C.Do(ctx, "getRolePermissionAll", nil, &result)
	return result, err
}

// GetProductsLevel 获取产品等级
func (s *Service) GetProductsLevel(ctx context.Context) ([]ProductLevel, error) {
	var result []ProductLevel
	err := s.C.Do(ctx, "getProductsLevel", nil, &result)
	return result, err
}

// GetUserCategory 获取自定义分类
func (s *Service) GetUserCategory(ctx context.Context) ([]UserCategory, error) {
	var result []UserCategory
	err := s.C.Do(ctx, "getUserCategory", nil, &result)
	return result, err
}

// GetQcOption 获取质检项
func (s *Service) GetQcOption(ctx context.Context) ([]QcOption, error) {
	var result []QcOption
	err := s.C.Do(ctx, "getQcOption", nil, &result)
	return result, err
}

// GetPackage 获取包材
func (s *Service) GetPackage(ctx context.Context) ([]PackageMaterial, error) {
	var result []PackageMaterial
	err := s.C.Do(ctx, "getPackage", nil, &result)
	return result, err
}

// GetSupplier 获取供应商
func (s *Service) GetSupplier(ctx context.Context) ([]Supplier, error) {
	var result []Supplier
	err := s.C.Do(ctx, "getSupplier", nil, &result)
	return result, err
}

// GetProductsParent 获取产品款式
func (s *Service) GetProductsParent(ctx context.Context) ([]ProductParent, error) {
	var result []ProductParent
	err := s.C.Do(ctx, "getProductsParent", nil, &result)
	return result, err
}

// GetProductTroubleType 获取产品问题类型
func (s *Service) GetProductTroubleType(ctx context.Context) ([]ProductTroubleType, error) {
	var result []ProductTroubleType
	err := s.C.Do(ctx, "getProductTroubleType", nil, &result)
	return result, err
}

// GetProductCategoryBase 获取产品品类
func (s *Service) GetProductCategoryBase(ctx context.Context) ([]ProductCategoryBase, error) {
	var result []ProductCategoryBase
	err := s.C.Do(ctx, "getProductCategoryBase", nil, &result)
	return result, err
}

// GetPaymentAccount 获取支付账号信息
func (s *Service) GetPaymentAccount(ctx context.Context) ([]PaymentAccount, error) {
	var result []PaymentAccount
	err := s.C.Do(ctx, "getPaymentAccount", nil, &result)
	return result, err
}

// GetPurchaseShipper 获取采购承运商
func (s *Service) GetPurchaseShipper(ctx context.Context) ([]PurchaseShipper, error) {
	var result []PurchaseShipper
	err := s.C.Do(ctx, "getPurchaseShipper", nil, &result)
	return result, err
}

// CreateProductForWms 创建产品（WMS）
func (s *Service) CreateProductForWms(ctx context.Context, req *CreateProductForWmsRequest) error {
	return s.C.Do(ctx, "createProductFowWms", req, nil)
}
