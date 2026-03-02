package fba_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/fba"
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

// fba 包实际只有 GetFbaReturn / GetFbaReimbursement
// Return 字段: ReturnID, OrderCode, FNSKU, SKU, Quantity
func TestGetFbaReturn_Success(t *testing.T) {
	want := []fba.Return{{ReturnID: 1, OrderCode: "ORD001", SKU: "SKU001"}}
	svc := fba.NewService(okCaller(t, "getFbaReturn", want))
	got, err := svc.GetFbaReturn(context.Background(), &fba.GetFbaReturnRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].ReturnID != 1 {
		t.Errorf("ReturnID 不符")
	}
}

// Reimbursement 字段: ReimbursementID, CaseID, SKU, Quantity
func TestGetFbaReimbursement_Success(t *testing.T) {
	want := []fba.Reimbursement{{ReimbursementID: "RB001", SKU: "SKU001"}}
	svc := fba.NewService(okCaller(t, "getFbaReimbursement", want))
	got, err := svc.GetFbaReimbursement(context.Background(), &fba.GetFbaReimbursementRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
