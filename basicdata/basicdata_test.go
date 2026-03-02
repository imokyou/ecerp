package basicdata_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/basicdata"
)

// ─── mock ───────────────────────────────────────

type mockCaller struct {
	doFn func(ctx context.Context, method string, biz interface{}, result interface{}) error
}

func (m *mockCaller) Do(ctx context.Context, method string, biz interface{}, result interface{}) error {
	if m.doFn != nil {
		return m.doFn(ctx, method, biz, result)
	}
	return nil
}

func successCaller(t *testing.T, wantMethod string, v interface{}) *mockCaller {
	t.Helper()
	return &mockCaller{doFn: func(_ context.Context, m string, _ interface{}, result interface{}) error {
		if m != wantMethod {
			t.Errorf("期望 method=%s, 收到 %s", wantMethod, m)
		}
		if result == nil || v == nil {
			return nil
		}
		b, _ := json.Marshal(v)
		return json.Unmarshal(b, result)
	}}
}

func errCaller(err error) *mockCaller {
	return &mockCaller{doFn: func(_ context.Context, _ string, _ interface{}, _ interface{}) error { return err }}
}

// ─── GetAllWarehouse ────────────────────────────

func TestGetAllWarehouse_Success(t *testing.T) {
	want := []basicdata.Warehouse{{WarehouseCode: "WH01", WarehouseName: "深圳仓"}}
	svc := basicdata.NewService(successCaller(t, "getWarehouses", want))
	got, err := svc.GetAllWarehouse(context.Background())
	if err != nil {
		t.Fatalf("不应返回错误: %v", err)
	}
	if len(got) != 1 || got[0].WarehouseCode != "WH01" {
		t.Errorf("数据不符: %+v", got)
	}
}

func TestGetAllWarehouse_Error(t *testing.T) {
	svc := basicdata.NewService(errCaller(&ecerp.APIError{Code: 401, Message: "unauth"}))
	_, err := svc.GetAllWarehouse(context.Background())
	if err == nil {
		t.Fatal("期望返回错误")
	}
	var apiErr *ecerp.APIError
	if !errors.As(err, &apiErr) || apiErr.Code != 401 {
		t.Errorf("期望 APIError code=401")
	}
}

// ─── GetWarehouseForOrder ───────────────────────

func TestGetWarehouseForOrder_Success(t *testing.T) {
	want := []basicdata.Warehouse{{WarehouseCode: "WH02"}}
	svc := basicdata.NewService(successCaller(t, "getWarehouseForOrder", want))
	got, err := svc.GetWarehouseForOrder(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// ─── GetShippingMethod ──────────────────────────

func TestGetShippingMethod_Success(t *testing.T) {
	want := []basicdata.ShippingMethod{{ShippingMethodCode: "DHL", ShippingMethodName: "DHL Express"}}
	svc := basicdata.NewService(successCaller(t, "getShippingMethod", want))
	got, err := svc.GetShippingMethod(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].ShippingMethodCode != "DHL" {
		t.Errorf("ShippingMethodCode 不符")
	}
}

func TestGetShippingMethodForOrder_Success(t *testing.T) {
	want := []basicdata.ShippingMethod{{ShippingMethodCode: "UPS"}}
	svc := basicdata.NewService(successCaller(t, "getShippingMethodForOrder", want))
	got, err := svc.GetShippingMethodForOrder(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// ─── GetCountry / GetCurrency ───────────────────

func TestGetCountry_Success(t *testing.T) {
	want := []basicdata.Country{{CountryCode: "US", CountryName: "美国"}}
	svc := basicdata.NewService(successCaller(t, "getCountry", want))
	got, err := svc.GetCountry(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetCurrency_Success(t *testing.T) {
	want := []basicdata.Currency{{CurrencyCode: "USD", CurrencyName: "美元"}}
	svc := basicdata.NewService(successCaller(t, "getCurrency", want))
	got, err := svc.GetCurrency(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// ─── GetUser / Organization ─────────────────────

func TestGetUser_Success(t *testing.T) {
	want := []basicdata.User{{UserID: 1, UserName: "admin"}}
	svc := basicdata.NewService(successCaller(t, "getUser", want))
	got, err := svc.GetUser(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetUserOrganization_Success(t *testing.T) {
	want := []basicdata.Organization{{OrgID: 10, OrgName: "研发部"}}
	svc := basicdata.NewService(successCaller(t, "getUserOrganization", want))
	got, err := svc.GetUserOrganization(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// ─── BusinessLicense ────────────────────────────

func TestGetBusinessLicenseList_Success(t *testing.T) {
	want := []basicdata.BusinessLicense{{BlID: 1, CompanyName: "测试公司"}}
	svc := basicdata.NewService(successCaller(t, "getBusinessLicenseList", want))
	got, err := svc.GetBusinessLicenseList(context.Background(), &basicdata.GetBusinessLicenseListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestBusinessLicenseBindAccount_Success(t *testing.T) {
	called := false
	svc := basicdata.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "businessLicenseBindAccount" {
			t.Errorf("期望 method=businessLicenseBindAccount, 收到 %s", m)
		}
		return nil
	}})
	err := svc.BusinessLicenseBindAccount(context.Background(), &basicdata.BusinessLicenseBindAccountRequest{BlID: 1, AccountIDs: []int{10, 11}})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestDeleteBusinessLicense_Success(t *testing.T) {
	svc := basicdata.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "deleteBusinessLicense" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	if err := svc.DeleteBusinessLicense(context.Background(), 99); err != nil {
		t.Fatalf("err=%v", err)
	}
}

// ─── 其他基础数据 ────────────────────────────────

func TestGetSupplier_Success(t *testing.T) {
	want := []basicdata.Supplier{{SupplierCode: "SUP01", SupplierName: "优质供应商"}}
	svc := basicdata.NewService(successCaller(t, "getSupplier", want))
	got, err := svc.GetSupplier(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetProductsLevel_Success(t *testing.T) {
	want := []basicdata.ProductLevel{{LevelID: 1, LevelName: "A级"}}
	svc := basicdata.NewService(successCaller(t, "getProductsLevel", want))
	got, err := svc.GetProductsLevel(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetPackage_Success(t *testing.T) {
	want := []basicdata.PackageMaterial{{PackageName: "标准纸箱"}}
	svc := basicdata.NewService(successCaller(t, "getPackage", want))
	got, err := svc.GetPackage(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestCreateProductForWms_Success(t *testing.T) {
	svc := basicdata.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "createProductFowWms" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateProductForWms(context.Background(), &basicdata.CreateProductForWmsRequest{SKU: "SKU-001", ProductName: "测试产品"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
