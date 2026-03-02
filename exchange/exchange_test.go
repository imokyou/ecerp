package exchange_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/exchange"
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

// exchange 包实际: GetCurrencyList / EditCurrency
// CurrencyRate 字段参见 exchange.go
func TestGetCurrencyList_Success(t *testing.T) {
	want := []exchange.CurrencyRate{{CurrencyCode: "USD", Rate: 7.2}}
	svc := exchange.NewService(okCaller(t, "getCurrencyList", want))
	got, err := svc.GetCurrencyList(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
	if got[0].CurrencyCode != "USD" {
		t.Errorf("CurrencyCode 不符")
	}
}

func TestEditCurrency_Success(t *testing.T) {
	svc := exchange.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "editCurrency" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.EditCurrency(context.Background(), &exchange.AddCurrencyRateRequest{CurrencyCode: "USD", Rate: 7.2})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
