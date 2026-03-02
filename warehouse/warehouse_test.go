package warehouse_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/warehouse"
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

func TestGetWarehouseList_Success(t *testing.T) {
	want := []warehouse.Warehouse{{WarehouseCode: "WH01", WarehouseName: "深圳仓"}}
	svc := warehouse.NewService(okCaller(t, "getWarehouseList", want))
	got, err := svc.GetWarehouseList(context.Background(), &warehouse.GetWarehouseListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].WarehouseCode != "WH01" {
		t.Errorf("WarehouseCode 不符")
	}
}

func TestGetWarehouse_Success(t *testing.T) {
	want := &warehouse.Warehouse{WarehouseCode: "WH01", WarehouseName: "广州仓"}
	svc := warehouse.NewService(okCaller(t, "getWarehouse", want))
	got, err := svc.GetWarehouse(context.Background(), "WH01")
	if err != nil || got == nil {
		t.Fatalf("err=%v", err)
	}
}

// ShippingMethod 实际字段：ShippingMethodCode, ShippingMethodName, WarehouseCode
func TestGetWarehouseShippingMethod_Success(t *testing.T) {
	want := []warehouse.ShippingMethod{{ShippingMethodCode: "DHL"}}
	svc := warehouse.NewService(okCaller(t, "getWarehouseShippingMethod", want))
	got, err := svc.GetWarehouseShippingMethod(context.Background(), "WH01")
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetWarehouseLocation_Success(t *testing.T) {
	want := []warehouse.WarehouseLocation{{LocationCode: "A01-01-01", WarehouseCode: "WH01"}}
	svc := warehouse.NewService(okCaller(t, "getWarehouseLocation", want))
	got, err := svc.GetWarehouseLocation(context.Background(), "WH01")
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// WarehouseLocationType 实际字段：TypeID, TypeName
func TestGetWarehouseLocationType_Success(t *testing.T) {
	want := []warehouse.WarehouseLocationType{{TypeID: 1, TypeName: "普通库位"}}
	svc := warehouse.NewService(okCaller(t, "getWarehouseLocationType", want))
	got, err := svc.GetWarehouseLocationType(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestSyncWarehouse_Success(t *testing.T) {
	called := false
	svc := warehouse.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "syncWarehouse" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncWarehouse(context.Background(), &warehouse.SyncWarehouseRequest{WarehouseCode: "WH-NEW"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestSyncWarehouseLocation_Success(t *testing.T) {
	svc := warehouse.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "syncWarehouseLocation" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncWarehouseLocation(context.Background(), &warehouse.SyncWarehouseLocationRequest{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
