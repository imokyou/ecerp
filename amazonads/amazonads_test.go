package amazonads_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/amazonads"
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

// amazonads 实际方法: ReimbursementDownload / ReportDownload / AdInvoice...
// 方法均返回 []map[string]interface{}
func TestReimbursementDownload_Success(t *testing.T) {
	want := []map[string]interface{}{{"reimbursement_id": "R001"}}
	svc := amazonads.NewService(okCaller(t, "ReimbursementDownload", want))
	got, err := svc.ReimbursementDownload(context.Background(), &amazonads.ReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// TaskStatus 字段: TaskID, Status
func TestGetTasksStatus_Success(t *testing.T) {
	want := []amazonads.TaskStatus{{TaskID: "T001", Status: "completed"}}
	svc := amazonads.NewService(okCaller(t, "GetTasksStatus", want))
	got, err := svc.GetTasksStatus(context.Background(), &amazonads.ReportRequest{})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

// AdStore 字段参见 amazonads.go
func TestGetAuthAdStoreSiteList_Success(t *testing.T) {
	want := []amazonads.AdStore{{AccountID: 1, AccountName: "测试广告账号"}}
	svc := amazonads.NewService(okCaller(t, "GetAuthAdStoreSiteList", want))
	got, err := svc.GetAuthAdStoreSiteList(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
