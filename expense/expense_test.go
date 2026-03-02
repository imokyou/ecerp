package expense_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/expense"
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

// expense 包实际方法: PaymentRecordsVerify / GetSuPaymentAccount / GetReceivePaymentBank / GetPurchaseCompany / GetPaymentBank
func TestGetSuPaymentAccount_Success(t *testing.T) {
	want := []expense.PaymentAccount{{AccountID: 1, AccountName: "招商银行账号"}}
	svc := expense.NewService(okCaller(t, "getSuPaymentAccount", want))
	got, err := svc.GetSuPaymentAccount(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetPurchaseCompany_Success(t *testing.T) {
	want := []expense.PurchaseCompany{{CompanyID: 1, CompanyName: "测试公司"}}
	svc := expense.NewService(okCaller(t, "getPurchaseCompany", want))
	got, err := svc.GetPurchaseCompany(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestPaymentRecordsVerify_Success(t *testing.T) {
	svc := expense.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "paymentRecordsVerify" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.PaymentRecordsVerify(context.Background(), &expense.PaymentRecordsVerifyRequest{})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
