package inventory_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/inventory"
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

// ProductInventory 字段: ProductSKU, WarehouseCode, AvailableQty, TotalQty...
func TestGetProductInventory_Success(t *testing.T) {
	want := []inventory.ProductInventory{{ProductSKU: "SKU001", AvailableQty: 100}}
	svc := inventory.NewService(okCaller(t, "getProductInventory", want))
	got, err := svc.GetProductInventory(context.Background(), &inventory.GetProductInventoryRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].AvailableQty != 100 {
		t.Errorf("AvailableQty 不符: %d", got[0].AvailableQty)
	}
}

func TestGetProductInventoryNew_Success(t *testing.T) {
	want := []inventory.ProductInventory{{ProductSKU: "SKU002", AvailableQty: 50}}
	svc := inventory.NewService(okCaller(t, "getProductInventoryNew", want))
	got, err := svc.GetProductInventoryNew(context.Background(), &inventory.GetProductInventoryRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// FbaInventory 字段参见 inventory.go
func TestGetFbaInventory_Success(t *testing.T) {
	want := []inventory.FbaInventory{{SellerSKU: "FBA-SKU"}}
	svc := inventory.NewService(okCaller(t, "getFbaInventory", want))
	got, err := svc.GetFbaInventory(context.Background(), &inventory.PageRequest{Page: 1, PageSize: 10})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// InventoryBatch 有 ProductSKU + Quantity
func TestGetInventoryBatch_Success(t *testing.T) {
	want := []inventory.InventoryBatch{{BatchCode: "BATCH001", ProductSKU: "SKU001", Quantity: 30}}
	svc := inventory.NewService(okCaller(t, "getInventoryBatch", want))
	got, err := svc.GetInventoryBatch(context.Background(), &inventory.GetProductInventoryRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestAdjustInventoryBatch_Success(t *testing.T) {
	called := false
	svc := inventory.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "adjustInventoryBatch" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.AdjustInventoryBatch(context.Background(), &inventory.AdjustInventoryBatchRequest{})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestMoveInventoryBatch_Success(t *testing.T) {
	svc := inventory.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "moveInventoryBatch" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.MoveInventoryBatch(context.Background(), &inventory.MoveInventoryBatchRequest{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

// LocationInventory 有 Quantity
func TestGetProductInventoryForLocation_Success(t *testing.T) {
	want := []inventory.LocationInventory{{LocationCode: "LOC-A1", ProductSKU: "SKU001", Quantity: 5}}
	svc := inventory.NewService(okCaller(t, "getProductInventoryForLocation", want))
	got, err := svc.GetProductInventoryForLocation(context.Background(), &inventory.GetProductInventoryRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
