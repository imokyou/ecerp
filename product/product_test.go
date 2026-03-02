package product_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/product"
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

func errCaller(err error) *mockCaller {
	return &mockCaller{doFn: func(_ context.Context, _ string, _ interface{}, _ interface{}) error { return err }}
}

func TestSyncProduct_Success(t *testing.T) {
	called := false
	svc := product.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		called = true
		if m != "syncProduct" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncProduct(context.Background(), &product.SyncProductRequest{ProductSKU: "SKU001"})
	if err != nil || !called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}

func TestGetProductBySku_Success(t *testing.T) {
	// Product 实际字段：ProductSKU, ProductName
	want := &product.Product{ProductSKU: "SKU001", ProductName: "测试产品"}
	svc := product.NewService(okCaller(t, "getProductBySku", want))
	got, err := svc.GetProductBySku(context.Background(), "SKU001")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if got == nil || got.ProductSKU != "SKU001" {
		t.Errorf("数据不符: %+v", got)
	}
}

func TestGetProductBySku_Error_ReturnsNil(t *testing.T) {
	svc := product.NewService(errCaller(&ecerp.APIError{Code: 404, Message: "not found"}))
	_, err := svc.GetProductBySku(context.Background(), "GHOST")
	// GetProductBySku 出错时返回 err 非 nil，我们只验证 err 不为 nil
	if err == nil {
		t.Fatal("期望返回错误")
	}
}

func TestGetWmsProductList_Success(t *testing.T) {
	want := []product.Product{{ProductSKU: "SKU001"}, {ProductSKU: "SKU002"}}
	svc := product.NewService(okCaller(t, "getWmsProductList", want))
	got, err := svc.GetWmsProductList(context.Background(), &product.GetProductListRequest{PageRequest: product.PageRequest{Page: 1, PageSize: 10}})
	if err != nil || len(got) != 2 {
		t.Fatalf("err=%v len=%d", err, len(got))
	}
}

func TestSyncBatchProduct_Success(t *testing.T) {
	svc := product.NewService(&mockCaller{doFn: func(_ context.Context, m string, _ interface{}, _ interface{}) error {
		if m != "syncBatchProduct" {
			t.Errorf("wrong method: %s", m)
		}
		return nil
	}})
	err := svc.SyncBatchProduct(context.Background(), &product.SyncBatchProductRequest{Products: []product.SyncProductRequest{{ProductSKU: "A"}}})
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestGetSaleStatus_Success(t *testing.T) {
	// SaleStatus 实际字段：StatusID, StatusName
	want := []product.SaleStatus{{StatusID: 1, StatusName: "在售"}}
	svc := product.NewService(okCaller(t, "getSaleStatus", want))
	got, err := svc.GetSaleStatus(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetSkuRelation_Success(t *testing.T) {
	// SkuRelation 实际字段参见源码
	want := []product.SkuRelation{{SKU: "PARENT", RelatedSKU: "CHILD"}}
	svc := product.NewService(okCaller(t, "getSkuRelation", want))
	got, err := svc.GetSkuRelation(context.Background(), "PARENT")
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetProductColor_Success(t *testing.T) {
	want := []product.ProductColor{{ColorName: "红色"}}
	svc := product.NewService(okCaller(t, "getProductColor", want))
	got, err := svc.GetProductColor(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}

func TestGetProductSize_Success(t *testing.T) {
	want := []product.ProductSize{{SizeName: "L"}}
	svc := product.NewService(okCaller(t, "getProductSize", want))
	got, err := svc.GetProductSize(context.Background())
	if err != nil || len(got) == 0 {
		t.Fatalf("err=%v", err)
	}
}
