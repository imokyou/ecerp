package finance_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/finance"
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

// finance 包方法返回 []map[string]interface{}, 请求用 FinancialReportRequest
func TestGetFinancialOrderReportDetail_Success(t *testing.T) {
	want := []map[string]interface{}{{"order_code": "ORD001", "amount": 100.5}}
	svc := finance.NewService(okCaller(t, "getFinancialOrderReportDetail", want))
	got, err := svc.GetFinancialOrderReportDetail(context.Background(), &finance.FinancialReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetFeeTransferRecords_Success(t *testing.T) {
	want := []map[string]interface{}{{"record_id": "R001"}}
	svc := finance.NewService(okCaller(t, "getFeeTransferRecords", want))
	got, err := svc.GetFeeTransferRecords(context.Background(), &finance.PageRequest{Page: 1, PageSize: 10})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetFbaInventoryReport_Success(t *testing.T) {
	want := []map[string]interface{}{{"sku": "SKU001", "qty": 50}}
	svc := finance.NewService(okCaller(t, "getFbaInventoryReport", want))
	got, err := svc.GetFbaInventoryReport(context.Background(), &finance.FinancialReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
