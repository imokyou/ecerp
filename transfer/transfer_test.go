package transfer_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/transfer"
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

func TestCreateTransferOrder_Success(t *testing.T) {
	called := false
	svc := transfer.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "createTranferOrder" { // 注意: 源码 typo "Tranfer"
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateTransferOrder(context.Background(), &transfer.CreateTransferOrderRequest{FromWarehouseCode: "WH01", ToWarehouseCode: "WH02"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

// 实际方法名 "getTransferOrders"（源码内部 API 名）
func TestGetTransferOrderList_Success(t *testing.T) {
	want := []transfer.TransferOrder{{TransferCode: "TR001", FromWarehouseCode: "WH01"}}
	svc := transfer.NewService(okCaller(t, "getTransferOrders", want))
	got, err := svc.GetTransferOrderList(context.Background(), &transfer.GetTransferOrderListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].TransferCode != "TR001" {
		t.Errorf("TransferCode 不符")
	}
}

func TestEditTransferOrder_Success(t *testing.T) {
	svc := transfer.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "editTranferOrder" { // 注意: 源码 typo "Tranfer"
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.EditTransferOrder(context.Background(), &transfer.EditTransferOrderRequest{TransferCode: "TR001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
