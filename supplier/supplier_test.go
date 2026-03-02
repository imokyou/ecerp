package supplier_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/supplier"
)

type mockCaller struct {
	doFn func(context.Context, string, interface{}, interface{}) error
}

func (m *mockCaller) Do(ctx context.Context, method string, biz interface{}, result interface{}) error {
	if m.doFn != nil {
		return m.doFn(ctx, method, biz, result)
	}
	return nil
}

func okCaller(t *testing.T, wantMethod string, v interface{}) *mockCaller {
	t.Helper()
	return &mockCaller{doFn: func(_ context.Context, m string, _ interface{}, result interface{}) error {
		if m != wantMethod {
			t.Errorf("期望 method=%s, 收到 %s", wantMethod, m)
		}
		if result != nil && v != nil {
			b, _ := json.Marshal(v)
			return json.Unmarshal(b, result)
		}
		return nil
	}}
}

func TestGetSupplierList_Success(t *testing.T) {
	// Supplier 字段: SupplierCode, SupplierName, ContactName...
	want := []supplier.Supplier{{SupplierCode: "SUP001", SupplierName: "优质供应商"}}
	svc := supplier.NewService(okCaller(t, "getSupplierList", want))
	got, err := svc.GetSupplierList(context.Background(), &supplier.GetSupplierListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestSyncSupplierInfo_Success(t *testing.T) {
	svc := supplier.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "syncSupplierInfo" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncSupplierInfo(context.Background(), &supplier.SyncSupplierInfoRequest{SupplierCode: "SUP001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

// SupplierDetail 嵌入 Supplier（继承 SupplierCode/SupplierName）+ BankName, TaxID...
func TestGetSupplierInfo_Success(t *testing.T) {
	want := &supplier.SupplierDetail{Supplier: supplier.Supplier{SupplierCode: "SUP001", SupplierName: "优质供应商"}}
	svc := supplier.NewService(okCaller(t, "getSupplierInfo", want))
	got, err := svc.GetSupplierInfo(context.Background(), "SUP001")
	if err != nil || got == nil {
		t.Fatalf("err=%v", err)
	}
	if got.SupplierCode != "SUP001" {
		t.Errorf("SupplierCode 不符: %s", got.SupplierCode)
	}
}

// SupplierProduct 字段: SupplierCode, SKU, SupplierSKU...
func TestGetSupplierProductList_Success(t *testing.T) {
	want := []supplier.SupplierProduct{{SupplierCode: "SUP001", SKU: "SKU001"}}
	svc := supplier.NewService(okCaller(t, "getSupplierProductList", want))
	got, err := svc.GetSupplierProductList(context.Background(), &supplier.GetSupplierProductListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// Carrier 字段: CarrierID, CarrierCode, CarrierName
func TestGetCarrier_Success(t *testing.T) {
	want := []supplier.Carrier{{CarrierCode: "DHL", CarrierName: "DHL快递"}}
	svc := supplier.NewService(okCaller(t, "getCarrier", want))
	got, err := svc.GetCarrier(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
