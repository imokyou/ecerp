package outbound_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/outbound"
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

// outbound 包实际使用 CeiveOrder / SaveCeive / GetCeiveUseList 等
func TestGetCeiveUseList_Success(t *testing.T) {
	want := []outbound.CeiveOrder{{CeiveCode: "OB001", WarehouseCode: "WH01"}}
	svc := outbound.NewService(okCaller(t, "getCeiveUseList", want))
	got, err := svc.GetCeiveUseList(context.Background(), &outbound.GetCeiveUseListRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].CeiveCode != "OB001" {
		t.Errorf("CeiveCode 不符")
	}
}

func TestSaveCeive_Success(t *testing.T) {
	called := false
	svc := outbound.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "saveCeive" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SaveCeive(context.Background(), &outbound.SaveCeiveRequest{WarehouseCode: "WH01", ActionType: "out"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestGetShippingBoxNumber_Success(t *testing.T) {
	want := []outbound.ShippingBoxNumber{{OrderCode: "ORD001", BoxNumber: "B001"}}
	svc := outbound.NewService(okCaller(t, "getShippingBoxNumber", want))
	got, err := svc.GetShippingBoxNumber(context.Background(), &outbound.GetShippingBoxNumberRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
