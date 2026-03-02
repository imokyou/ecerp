package firstmile_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/firstmile"
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

// firstmile 包实际有 TransferBatch / StockPlan / StpoOrder 等
// TransferBatch 字段: BatchCode, ShippingMethod, WarehouseCode, Status
func TestCreateShipmentBatch_Success(t *testing.T) {
	called := false
	svc := firstmile.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "createShipmentBatch" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateShipmentBatch(context.Background(), &firstmile.CreateShipmentBatchRequest{WarehouseCode: "WH01"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

// StockPlan 字段: PlanID, PlanCode, Status
func TestCreateStockPlan_Success(t *testing.T) {
	svc := firstmile.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "createStockPlan" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateStockPlan(context.Background(), &firstmile.CreateStockPlanRequest{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestGetTransferBatchList_Success(t *testing.T) {
	want := []firstmile.TransferBatch{{BatchCode: "TB001", WarehouseCode: "WH01"}}
	svc := firstmile.NewService(okCaller(t, "getTransferBatchList", want))
	got, err := svc.GetTransferBatchList(context.Background(), &firstmile.PageRequest{Page: 1, PageSize: 10})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].BatchCode != "TB001" {
		t.Errorf("BatchCode 不符")
	}
}
