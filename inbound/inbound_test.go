package inbound_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/inbound"
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

// inbound 包实际使用 SyncReceiving / GetReceivingList 等
func TestSyncReceiving_Success(t *testing.T) {
	called := false
	svc := inbound.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "syncReceiving" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncReceiving(context.Background(), &inbound.SyncReceivingRequest{WarehouseCode: "WH01"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

// ReceivingOrder 字段: ReceivingCode, WarehouseCode, PoCode, Status
func TestGetReceivingList_Success(t *testing.T) {
	want := []inbound.ReceivingOrder{{ReceivingCode: "RCV001", WarehouseCode: "WH01"}}
	svc := inbound.NewService(okCaller(t, "getReceiving", want))
	got, err := svc.GetReceiving(context.Background(), &inbound.GetReceivingRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].ReceivingCode != "RCV001" {
		t.Errorf("ReceivingCode 不符")
	}
}

// PutAwayOrder 字段: PutAwayCode, ReceivingCode, WarehouseCode
func TestGetPutAwayList_Success(t *testing.T) {
	want := []inbound.PutAwayOrder{{PutAwayCode: "PTA001", WarehouseCode: "WH01"}}
	svc := inbound.NewService(okCaller(t, "getPutAwayList", want))
	got, err := svc.GetPutAwayList(context.Background(), &inbound.GetPutAwayListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
