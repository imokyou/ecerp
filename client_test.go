package ecerp_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/basicdata"
	"github.com/imokyou/ecerp/order"
	"github.com/imokyou/ecerp/product"
)

func newTestClient(t *testing.T, serverURL string) *ecerp.Client {
	t.Helper()
	client, err := ecerp.NewClient("test_key", "test_secret", "test_service",
		ecerp.WithBaseURL(serverURL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func TestNewClient_Validation(t *testing.T) {
	tests := []struct {
		name      string
		appKey    string
		appSecret string
		serviceID string
		wantErr   bool
	}{
		{"缺少 appKey", "", "secret", "service", true},
		{"缺少 appSecret", "key", "", "service", true},
		{"缺少 serviceID", "key", "secret", "", true},
		{"全部合法", "key", "secret", "service", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ecerp.NewClient(tt.appKey, tt.appSecret, tt.serviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMustNewClient_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustNewClient 应在参数为空时 panic")
		}
	}()
	ecerp.MustNewClient("", "", "")
}

func TestClientDo_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("期望 POST 请求, 收到 %s", r.Method)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("期望 Content-Type application/json, 收到 %s", ct)
		}
		if ua := r.Header.Get("User-Agent"); ua == "" {
			t.Error("User-Agent 不应为空")
		}

		var req ecerp.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("解析请求体失败: %v", err)
		}

		if req.AppKey != "test_key" {
			t.Errorf("期望 app_key=test_key, 收到 %s", req.AppKey)
		}
		if req.Sign == "" {
			t.Error("sign 不应为空")
		}

		resp := ecerp.Response{
			Code:    200,
			Message: "success",
			Data:    json.RawMessage(`[{"warehouse_code":"WH001","warehouse_name":"深圳仓"}]`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	svc := basicdata.NewService(client)
	warehouses, err := svc.GetAllWarehouse(context.Background())
	if err != nil {
		t.Fatalf("GetAllWarehouse() error = %v", err)
	}

	if len(warehouses) != 1 {
		t.Fatalf("期望 1 个仓库, 收到 %d", len(warehouses))
	}
	if warehouses[0].WarehouseCode != "WH001" {
		t.Errorf("期望 warehouse_code=WH001, 收到 %s", warehouses[0].WarehouseCode)
	}
}

func TestClientDo_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ecerp.Response{Code: 401, Message: "签名验证失败"}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	svc := basicdata.NewService(client)
	_, err := svc.GetAllWarehouse(context.Background())
	if err == nil {
		t.Fatal("期望返回错误, 实际为 nil")
	}

	apiErr, ok := ecerp.IsAPIError(err)
	if !ok {
		t.Fatalf("期望 APIError 类型, 收到 %T", err)
	}
	if apiErr.Code != 401 {
		t.Errorf("期望 code=401, 收到 %d", apiErr.Code)
	}
	if !apiErr.IsAuthError() {
		t.Error("期望 IsAuthError 返回 true")
	}
}

func TestClientDo_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Bad Gateway"))
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	err := client.Do(context.Background(), "getWarehouse", nil, nil)
	if err == nil {
		t.Fatal("期望返回 HTTP 错误")
	}
}

func TestClientDo_EmptyMethod(t *testing.T) {
	client := ecerp.MustNewClient("key", "secret", "service")

	err := client.Do(context.Background(), "", nil, nil)
	if err == nil {
		t.Fatal("空 method 应返回错误")
	}
}

func TestClientDo_ClosedClient(t *testing.T) {
	client := ecerp.MustNewClient("key", "secret", "service")
	client.Close()

	err := client.Do(context.Background(), "test", nil, nil)
	if err == nil {
		t.Fatal("已关闭的客户端应返回错误")
	}
}

func TestClientDo_WithBizContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ecerp.Request
		json.NewDecoder(r.Body).Decode(&req)

		if req.BizContent == "{}" || req.BizContent == "" {
			t.Error("biz_content 不应为空")
		}

		var bizContent map[string]interface{}
		if err := json.Unmarshal([]byte(req.BizContent), &bizContent); err != nil {
			t.Fatalf("解析 biz_content 失败: %v", err)
		}

		if page, ok := bizContent["page"]; !ok || page.(float64) != 1 {
			t.Errorf("期望 page=1, 收到 %v", page)
		}

		resp := ecerp.Response{Code: 200, Message: "success", Data: json.RawMessage(`[]`)}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	svc := order.NewService(client)
	_, err := svc.GetOrderList(context.Background(), &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 10},
		Status:      1,
	})
	if err != nil {
		t.Fatalf("GetOrderList() error = %v", err)
	}
}

func TestClientDo_ProductModule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ecerp.Request
		json.NewDecoder(r.Body).Decode(&req)

		if req.InterfaceMethod != "getWmsProductList" {
			t.Errorf("期望 interface_method=getWmsProductList, 收到 %s", req.InterfaceMethod)
		}

		resp := ecerp.Response{
			Code:    200,
			Message: "success",
			Data:    json.RawMessage(`[{"product_sku":"SKU001","product_title":"测试产品"}]`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	svc := product.NewService(client)
	products, err := svc.GetWmsProductList(context.Background(), &product.GetProductListRequest{
		PageRequest: product.PageRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		t.Fatalf("GetProductList() error = %v", err)
	}
	if len(products) != 1 || products[0].ProductSKU != "SKU001" {
		t.Errorf("产品数据不符合预期: %+v", products)
	}
}

func TestClientDoRaw_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ecerp.Response{Code: 200, Message: "success", Data: json.RawMessage(`{"key":"value"}`)}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	resp, err := client.DoRaw(context.Background(), "getCountryList", nil)
	if err != nil {
		t.Fatalf("DoRaw() error = %v", err)
	}
	if resp.Code != 200 {
		t.Errorf("期望 code=200, 收到 %d", resp.Code)
	}
}
