package user_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp/user"
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

func TestCreateUser_Success(t *testing.T) {
	called := false
	svc := user.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "createUser" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.CreateUser(context.Background(), &user.CreateUserRequest{UserName: "test_user"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestEditUser_Success(t *testing.T) {
	svc := user.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "editUser" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.EditUser(context.Background(), &user.EditUserRequest{UserID: 1, UserName: "new_name"})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestBatchCreateUser_Success(t *testing.T) {
	svc := user.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "batchCreateUser" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.BatchCreateUser(context.Background(), []user.CreateUserRequest{{UserName: "user1"}, {UserName: "user2"}})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestGetUserAccountList_Success(t *testing.T) {
	want := []user.PlatformAccount{{AccountID: 1, AccountName: "亚马逊店铺A"}}
	svc := user.NewService(okCaller(t, "getUserAccountList", want))
	got, err := svc.GetUserAccountList(context.Background(), &user.PageRequest{Page: 1, PageSize: 10})
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
