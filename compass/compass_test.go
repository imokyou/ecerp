package compass_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/compass"
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

// compass 实际方法: GetOrderStatisticsV2 / GetOrderStatistics / GetProductSaleSummary / GetProductSale
// 返回 []map[string]interface{}
func TestGetOrderStatisticsV2_Success(t *testing.T) {
	want := []map[string]interface{}{{"date": "2024-01-01", "orders": 100}}
	svc := compass.NewService(okCaller(t, "getOrderStatisticsV2", want))
	got, err := svc.GetOrderStatisticsV2(context.Background(), &compass.OrderStatisticsRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetProductSale_Success(t *testing.T) {
	want := []map[string]interface{}{{"sku": "SKU001", "sales": 50}}
	svc := compass.NewService(okCaller(t, "getProductSale", want))
	got, err := svc.GetProductSale(context.Background(), &compass.ProductSaleRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
