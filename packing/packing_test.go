package packing_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/packing"
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

// PackageOrder 字段: OrderCode, PackageCode, WarehouseCode, ShippingMethod, Status
func TestGetPackageOrderCondition_Success(t *testing.T) {
	want := &packing.PackageOrder{OrderCode: "ORD001", PackageCode: "PKG001"}
	svc := packing.NewService(okCaller(t, "getPackageOrderCondition", want))
	got, err := svc.GetPackageOrderCondition(context.Background(), &packing.GetPackageOrderConditionRequest{OrderCode: "ORD001"})
	if err != nil || got == nil {
		t.Fatalf("err=%v", err)
	}
	if got.PackageCode != "PKG001" {
		t.Errorf("PackageCode 不符")
	}
}

// PackageOrderLabel 字段: OrderCode, LabelURL, LabelType
func TestGetPackageOrderLabel_Success(t *testing.T) {
	want := &packing.PackageOrderLabel{OrderCode: "ORD001", LabelURL: "https://example.com/label.pdf"}
	svc := packing.NewService(okCaller(t, "getPackageOrderLabel", want))
	got, err := svc.GetPackageOrderLabel(context.Background(), &packing.GetPackageOrderLabelRequest{OrderCode: "ORD001"})
	if err != nil || got == nil {
		t.Fatalf("err=%v", err)
	}
	if got.LabelURL == "" {
		t.Errorf("LabelURL 不应为空")
	}
}

func TestCreatePackageOrderNew_Success(t *testing.T) {
	svc := packing.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "createPackageOrderNew" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreatePackageOrderNew(context.Background(), &packing.FastPackageOrderRequest{OrderCode: "ORD001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
