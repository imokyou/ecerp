// Package order_test provides unit tests for the order sub-package.
package order_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/order"
)

// ─────────────────────────────────────────────
// Mock Caller
// ─────────────────────────────────────────────

// mockCaller 实现 ecerp.Caller 接口，用于单元测试。
// 通过 doFn 可以灵活模拟各种响应。
type mockCaller struct {
	doFn func(ctx context.Context, method string, bizContent interface{}, result interface{}) error
}

func (m *mockCaller) Do(ctx context.Context, method string, bizContent interface{}, result interface{}) error {
	if m.doFn != nil {
		return m.doFn(ctx, method, bizContent, result)
	}
	return nil
}

// successCaller 返回一个把 v 作为结果写入 result 的 mockCaller
func successCaller(t *testing.T, method string, v interface{}) *mockCaller {
	t.Helper()
	return &mockCaller{
		doFn: func(_ context.Context, m string, _ interface{}, result interface{}) error {
			if m != method {
				t.Errorf("期望 method=%s, 收到=%s", method, m)
			}
			if result == nil || v == nil {
				return nil
			}
			// 通过 JSON round-trip 填充 result
			b, _ := json.Marshal(v)
			return json.Unmarshal(b, result)
		},
	}
}

// errorCaller 返回固定错误的 mockCaller
func errorCaller(err error) *mockCaller {
	return &mockCaller{
		doFn: func(_ context.Context, _ string, _ interface{}, _ interface{}) error {
			return err
		},
	}
}

// ─────────────────────────────────────────────
// SyncOrder
// ─────────────────────────────────────────────

func TestSyncOrder_Success(t *testing.T) {
	svc := order.NewService(&mockCaller{
		doFn: func(_ context.Context, method string, _ interface{}, _ interface{}) error {
			if method != "syncOrder" {
				t.Errorf("期望 method=syncOrder, 收到 %s", method)
			}
			return nil
		},
	})

	err := svc.SyncOrder(context.Background(), &order.SyncOrderRequest{
		OrderCode:     "ORD001",
		WarehouseCode: "WH001",
	})
	if err != nil {
		t.Fatalf("SyncOrder() 不应返回错误, 收到: %v", err)
	}
}

func TestSyncOrder_Error(t *testing.T) {
	apiErr := &ecerp.APIError{Code: 400, Message: "参数错误"}
	svc := order.NewService(errorCaller(apiErr))

	err := svc.SyncOrder(context.Background(), &order.SyncOrderRequest{})
	if err == nil {
		t.Fatal("期望返回错误")
	}
	var gotErr *ecerp.APIError
	if !errors.As(err, &gotErr) {
		t.Fatalf("期望 *ecerp.APIError, 收到 %T", err)
	}
	if gotErr.Code != 400 {
		t.Errorf("期望 code=400, 收到 %d", gotErr.Code)
	}
}

// ─────────────────────────────────────────────
// GetOrderList
// ─────────────────────────────────────────────

func TestGetOrderList_Success(t *testing.T) {
	want := []order.Order{
		{OrderCode: "ORD001", Status: 1, Currency: "USD"},
		{OrderCode: "ORD002", Status: 2, Currency: "GBP"},
	}
	svc := order.NewService(successCaller(t, "getOrderList", want))

	got, err := svc.GetOrderList(context.Background(), &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		t.Fatalf("GetOrderList() 不应返回错误: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("期望2条订单, 收到 %d", len(got))
	}
	if got[0].OrderCode != "ORD001" {
		t.Errorf("期望 OrderCode=ORD001, 收到 %s", got[0].OrderCode)
	}
}

func TestGetOrderList_Empty(t *testing.T) {
	svc := order.NewService(successCaller(t, "getOrderList", []order.Order{}))
	got, err := svc.GetOrderList(context.Background(), &order.GetOrderListRequest{})
	if err != nil {
		t.Fatalf("GetOrderList() 不应返回错误: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("期望空列表, 收到 %d 条", len(got))
	}
}

// ─────────────────────────────────────────────
// GetOrderInfo
// ─────────────────────────────────────────────

func TestGetOrderInfo_Success(t *testing.T) {
	want := order.Order{OrderCode: "ORD999", Status: 3, BuyerName: "张三"}
	svc := order.NewService(successCaller(t, "getOrderInfo", want))

	got, err := svc.GetOrderInfo(context.Background(), "ORD999")
	if err != nil {
		t.Fatalf("GetOrderInfo() 不应返回错误: %v", err)
	}
	if got == nil {
		t.Fatal("GetOrderInfo() 不应返回 nil")
	}
	if got.OrderCode != "ORD999" {
		t.Errorf("期望 OrderCode=ORD999, 收到 %s", got.OrderCode)
	}
}

func TestGetOrderInfo_Error_ReturnsNil(t *testing.T) {
	// 验证 I5 修复：error 时返回 nil 而非非空指针
	svc := order.NewService(errorCaller(&ecerp.APIError{Code: 404, Message: "未找到"}))

	got, err := svc.GetOrderInfo(context.Background(), "NOT_EXIST")
	if err == nil {
		t.Fatal("期望返回错误")
	}
	if got != nil {
		t.Errorf("错误时期望返回 nil, 收到非空指针: %+v", got)
	}
}

// ─────────────────────────────────────────────
// CancelOrder
// ─────────────────────────────────────────────

func TestCancelOrder_Success(t *testing.T) {
	svc := order.NewService(&mockCaller{
		doFn: func(_ context.Context, method string, bizContent interface{}, _ interface{}) error {
			if method != "cancelOrder" {
				t.Errorf("期望 method=cancelOrder, 收到 %s", method)
			}
			return nil
		},
	})
	err := svc.CancelOrder(context.Background(), &order.CancelOrderRequest{
		OrderCode: "ORD001",
		Reason:    "客户取消",
	})
	if err != nil {
		t.Fatalf("CancelOrder() 不应返回错误: %v", err)
	}
}

// ─────────────────────────────────────────────
// UploadTrackingNoList (批量操作边界测试)
// ─────────────────────────────────────────────

func TestUploadTrackingNoList_WithItems(t *testing.T) {
	called := false
	svc := order.NewService(&mockCaller{
		doFn: func(_ context.Context, method string, bizContent interface{}, _ interface{}) error {
			called = true
			if method != "uploadTrackingNoList" {
				t.Errorf("期望 method=uploadTrackingNoList, 收到 %s", method)
			}
			return nil
		},
	})
	err := svc.UploadTrackingNoList(context.Background(), &order.UploadTrackingNoListRequest{
		List: []order.TrackingNoItem{
			{OrderCode: "ORD001", TrackingNumber: "TN001", ShippingMethod: "DHL"},
			{OrderCode: "ORD002", TrackingNumber: "TN002"},
		},
	})
	if err != nil {
		t.Fatalf("UploadTrackingNoList() 不应返回错误: %v", err)
	}
	if !called {
		t.Error("期望 Caller.Do 被调用，实际未调用")
	}
}

// ─────────────────────────────────────────────
// GetRmaReason (无参接口测试)
// ─────────────────────────────────────────────

func TestGetRmaReason_Success(t *testing.T) {
	want := []order.RmaReason{
		{ReasonID: 1, ReasonName: "商品损坏"},
		{ReasonID: 2, ReasonName: "发错货"},
	}
	svc := order.NewService(successCaller(t, "getRmaReason", want))

	got, err := svc.GetRmaReason(context.Background())
	if err != nil {
		t.Fatalf("GetRmaReason() 不应返回错误: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("期望2条原因, 收到 %d", len(got))
	}
}

// ─────────────────────────────────────────────
// 集成测试（真实 httptest.Server）
// ─────────────────────────────────────────────

func newIntegrationClient(t *testing.T, serverURL string) *ecerp.Client {
	t.Helper()
	c, err := ecerp.NewClient("key", "secret", "svc",
		ecerp.WithBaseURL(serverURL),
		ecerp.WithDisableRetry(), // 集成测试禁用重试加快速度
	)
	if err != nil {
		t.Fatalf("NewClient 失败: %v", err)
	}
	return c
}

func TestGetOrderList_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ecerp.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("解析请求失败: %v", err)
		}
		if req.InterfaceMethod != "getOrderList" {
			t.Errorf("期望 interface_method=getOrderList, 收到 %s", req.InterfaceMethod)
		}
		resp := ecerp.Response{
			Code:    200,
			Message: "success",
			Data:    json.RawMessage(`[{"order_code":"INT001","status":1}]`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newIntegrationClient(t, server.URL)
	defer client.Close()

	svc := order.NewService(client)
	orders, err := svc.GetOrderList(context.Background(), &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 20},
	})
	if err != nil {
		t.Fatalf("GetOrderList() 集成测试失败: %v", err)
	}
	if len(orders) != 1 || orders[0].OrderCode != "INT001" {
		t.Errorf("集成测试数据不符预期: %+v", orders)
	}
}

func TestGetOrderInfo_Integration_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ecerp.Response{Code: 404, Message: "订单不存在"}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newIntegrationClient(t, server.URL)
	defer client.Close()

	svc := order.NewService(client)
	got, err := svc.GetOrderInfo(context.Background(), "GHOST")
	if err == nil {
		t.Fatal("期望返回错误")
	}
	if got != nil {
		t.Errorf("错误时 GetOrderInfo 应返回 nil, 收到: %+v", got)
	}
	apiErr, ok := ecerp.IsAPIError(err)
	if !ok {
		t.Fatalf("期望 APIError, 收到 %T", err)
	}
	if !apiErr.IsNotFound() {
		t.Errorf("期望 IsNotFound=true, code=%d", apiErr.Code)
	}
}
