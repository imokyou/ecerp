package provider_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/provider"
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

// PreShippingOrder 字段: OrderCode, WarehouseCode, Status...
func TestGetPreShippingOrder_Success(t *testing.T) {
	want := []provider.PreShippingOrder{{OrderCode: "PRE001", WarehouseCode: "WH01"}}
	svc := provider.NewService(okCaller(t, "getPreShippingOrder", want))
	got, err := svc.GetPreShippingOrder(context.Background(), &provider.GetPreShippingOrderRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// UploadTrackingNoRequest 字段查看源码:  OrderCode, TrackingNumber...
func TestUploadTrackingNo_Success(t *testing.T) {
	svc := provider.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "uploadTrackingNo" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.UploadTrackingNo(context.Background(), &provider.UploadTrackingNoRequest{OrderCode: "PRE001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

// WarehouseInfo 字段: WarehouseCode, WarehouseName
func TestGetWarehouseInfo_Success(t *testing.T) {
	want := []provider.WarehouseInfo{{WarehouseCode: "WH01", WarehouseName: "深圳仓"}}
	svc := provider.NewService(okCaller(t, "getWarehouseInfo", want))
	got, err := svc.GetWarehouseInfo(context.Background(), &provider.GetWarehouseInfoRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// InventoryInfo 字段: SKU, WarehouseCode, AvailableQty, TotalQty
func TestGetInventory_Success(t *testing.T) {
	want := []provider.InventoryInfo{{SKU: "SKU001", AvailableQty: 50}}
	svc := provider.NewService(okCaller(t, "getInventory", want))
	got, err := svc.GetInventory(context.Background(), &provider.GetInventoryRequest{WarehouseCode: "WH01"})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestOrderStatusModify_Success(t *testing.T) {
	svc := provider.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "orderStatusModify" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.OrderStatusModify(context.Background(), &provider.OrderStatusModifyRequest{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
