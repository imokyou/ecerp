package purchase_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/purchase"
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

func TestSyncPurchaseOrders_Success(t *testing.T) {
	called := false
	svc := purchase.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "syncPurchaseOrders" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncPurchaseOrders(context.Background(), &purchase.SyncPurchaseOrderRequest{PurchaseCode: "PO001"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestGetPurchaseOrders_Success(t *testing.T) {
	// PurchaseOrder 字段: PurchaseCode, WarehouseCode, SupplierCode, Status...
	want := []purchase.PurchaseOrder{{PurchaseCode: "PO001", Status: 1}}
	svc := purchase.NewService(okCaller(t, "getPurchaseOrders", want))
	got, err := svc.GetPurchaseOrders(context.Background(), &purchase.GetPurchaseOrdersRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].PurchaseCode != "PO001" {
		t.Errorf("PurchaseCode 不符")
	}
}

func TestGetPurchaseRequestOrdersNew_Success(t *testing.T) {
	want := []purchase.PurchaseRequestOrder{{RequestCode: "REQ001"}}
	svc := purchase.NewService(okCaller(t, "getPurchaseRequestOrdersNew", want))
	got, err := svc.GetPurchaseRequestOrdersNew(context.Background(), &purchase.GetPurchaseRequestOrdersRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// PurchasePlan 实际字段: PlanID, SKU, ProductName, Quantity
func TestGetPurchasePlan_Success(t *testing.T) {
	want := []purchase.PurchasePlan{{PlanID: 1, SKU: "SKU001"}}
	svc := purchase.NewService(okCaller(t, "getPurchasePlan", want))
	got, err := svc.GetPurchasePlan(context.Background(), &purchase.GetPurchasePlanRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestPurchaseForceCompletion_Success(t *testing.T) {
	svc := purchase.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "purchaseForceCompletion" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	if err := svc.PurchaseForceCompletion(context.Background(), "PO001"); err != nil {
		t.Fatalf("err=%v", err)
	}
}

// PurchaseOrderFile 字段: FileID, FileName, FileURL, PurchaseCode
func TestGetPurchaseOrderFiles_Success(t *testing.T) {
	want := []purchase.PurchaseOrderFile{{FileName: "invoice.pdf", PurchaseCode: "PO001"}}
	svc := purchase.NewService(okCaller(t, "getPurchaseOrderFiles", want))
	got, err := svc.GetPurchaseOrderFiles(context.Background(), "PO001")
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestVerifyPurchase_Success(t *testing.T) {
	svc := purchase.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "verifyPurchase" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.VerifyPurchase(context.Background(), &purchase.VerifyPurchaseRequest{PurchaseCode: "PO001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestRevocationPurchase_Success(t *testing.T) {
	svc := purchase.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "revocationPurchase" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.RevocationPurchase(context.Background(), &purchase.RevocationPurchaseRequest{PurchaseCode: "PO001"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
