package fbasta_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/fbasta"
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

// fbasta 包实际使用 InboundPlan / CreateInboundPlan / GetInboundPlan / CancelInboundPlan
func TestCreateInboundPlan_Success(t *testing.T) {
	called := false
	svc := fbasta.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "createInboundPlan" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateInboundPlan(context.Background(), &fbasta.CreateInboundPlanRequest{DestinationMarketplaces: []string{"ATVPDKIKX0DER"}})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

// InboundPlan 字段: InboundPlanID, Name, Status
func TestGetInboundPlan_Success(t *testing.T) {
	want := &fbasta.InboundPlan{InboundPlanID: "PLAN001", Status: "ACTIVE"}
	svc := fbasta.NewService(okCaller(t, "getInboundPlan", want))
	got, err := svc.GetInboundPlan(context.Background(), "PLAN001")
	if err != nil || got == nil {
		t.Fatalf("err=%v", err)
	}
	if got.InboundPlanID != "PLAN001" {
		t.Errorf("InboundPlanID 不符")
	}
}

func TestCancelInboundPlan_Success(t *testing.T) {
	svc := fbasta.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "cancelInboundPlan" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CancelInboundPlan(context.Background(), "PLAN001")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
