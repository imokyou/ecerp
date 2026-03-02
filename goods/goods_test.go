package goods_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/goods"
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

// goods 包所有方法返回 []map[string]interface{}, 请求用 ReportRequest
func TestNewRemovalOrderList_Success(t *testing.T) {
	want := []map[string]interface{}{{"order_code": "RMV001"}}
	svc := goods.NewService(okCaller(t, "NewRemovalOrderList", want))
	got, err := svc.NewRemovalOrderList(context.Background(), &goods.ReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestAmazonSettlementReport_Success(t *testing.T) {
	want := []map[string]interface{}{{"settlement_id": "S001"}}
	svc := goods.NewService(okCaller(t, "AmazonSettlementReport", want))
	got, err := svc.AmazonSettlementReport(context.Background(), &goods.ReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
