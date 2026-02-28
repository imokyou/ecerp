// Package main 展示 ecerp SDK 基础使用方法
//
// 运行前请设置环境变量:
//
//	export ECERP_APP_KEY="your_app_key"
//	export ECERP_APP_SECRET="your_app_secret"
//	export ECERP_SERVICE_ID="your_service_id"
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/basicdata"
	"github.com/imokyou/ecerp/inventory"
	"github.com/imokyou/ecerp/order"
	"github.com/imokyou/ecerp/product"
	"github.com/imokyou/ecerp/purchase"
)

func main() {
	// ═══════════════════════════════════════════
	// 第一步：创建客户端
	// ═══════════════════════════════════════════
	client, err := ecerp.NewClient(
		getEnvOrDefault("ECERP_APP_KEY", "demo_key"),
		getEnvOrDefault("ECERP_APP_SECRET", "demo_secret"),
		getEnvOrDefault("ECERP_SERVICE_ID", "demo_service"),
		ecerp.WithTimeout(60*time.Second),        // 超时60秒
		ecerp.WithLogger(slog.Default()),         // 开启日志
		ecerp.WithUserAgent("ecerp-example/1.0"), // 自定义UA
	)
	if err != nil {
		log.Fatalf("❌ 创建客户端失败: %v", err)
	}
	defer client.Close()
	fmt.Println("✅ 客户端创建成功")

	// ═══════════════════════════════════════════
	// 第二步：创建各模块服务（可按需创建）
	// ═══════════════════════════════════════════
	basicSvc := basicdata.NewService(client)
	orderSvc := order.NewService(client)
	productSvc := product.NewService(client)
	invSvc := inventory.NewService(client)
	purchaseSvc := purchase.NewService(client)

	ctx := context.Background()

	// ═══════════════════════════════════════════
	// 示例 1: 基础数据 — 获取仓库列表
	// ═══════════════════════════════════════════
	fmt.Println("\n📦 === 获取仓库列表 ===")
	warehouses, err := basicSvc.GetAllWarehouse(ctx)
	if err != nil {
		handleError("获取仓库", err)
	} else {
		fmt.Printf("  共 %d 个仓库:\n", len(warehouses))
		for i, w := range warehouses {
			if i >= 5 {
				fmt.Printf("  ... 还有 %d 个\n", len(warehouses)-5)
				break
			}
			fmt.Printf("  [%d] %s (%s)\n", i+1, w.WarehouseName, w.WarehouseCode)
		}
	}

	// ═══════════════════════════════════════════
	// 示例 2: 订单 — 分页获取订单列表
	// ═══════════════════════════════════════════
	fmt.Println("\n🛒 === 获取订单列表 ===")
	orders, err := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		handleError("获取订单", err)
	} else {
		fmt.Printf("  获取到 %d 条订单\n", len(orders))
	}

	// ═══════════════════════════════════════════
	// 示例 3: 产品 — 按SKU查询产品
	// ═══════════════════════════════════════════
	fmt.Println("\n📋 === 查询产品列表 ===")
	products, err := productSvc.GetWmsProductList(ctx, &product.GetProductListRequest{
		PageRequest: product.PageRequest{Page: 1, PageSize: 5},
	})
	if err != nil {
		handleError("获取产品", err)
	} else {
		fmt.Printf("  获取到 %d 个产品\n", len(products))
		for _, p := range products {
			fmt.Printf("  SKU: %s | 名称: %s\n", p.ProductSKU, p.ProductName)
		}
	}

	// ═══════════════════════════════════════════
	// 示例 4: 库存 — 查询库存
	// ═══════════════════════════════════════════
	fmt.Println("\n📊 === 查询库存 ===")
	inv, err := invSvc.GetProductInventoryNew(ctx, &inventory.GetProductInventoryRequest{
		PageRequest: inventory.PageRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		handleError("查询库存", err)
	} else {
		fmt.Printf("  获取到 %d 条库存记录\n", len(inv))
	}

	// ═══════════════════════════════════════════
	// 示例 5: 采购 — 创建采购单
	// ═══════════════════════════════════════════
	fmt.Println("\n🏷️  === 创建采购单（演示，不实际提交）===")
	_ = purchaseSvc // 实际使用时取消注释:
	// err = purchaseSvc.SyncPurchaseOrders(ctx, &purchase.SyncPurchaseOrderRequest{
	//     WarehouseCode: "WH001",
	//     SupplierCode:  "SUP001",
	//     Currency:      "CNY",
	//     Items: []purchase.PurchaseOrderItem{
	//         {SKU: "SKU001", Quantity: 100, UnitPrice: 9.99},
	//     },
	// })
	fmt.Println("  (采购单创建代码已准备，取消注释即可使用)")

	fmt.Println("\n🎉 所有示例执行完毕！")
}

// handleError 统一错误处理
func handleError(operation string, err error) {
	if apiErr, ok := ecerp.IsAPIError(err); ok {
		switch {
		case apiErr.IsAuthError():
			fmt.Printf("  ❌ [%s] 认证失败(code=%d): %s\n", operation, apiErr.Code, apiErr.Message)
			fmt.Println("     → 请检查 ECERP_APP_KEY / ECERP_APP_SECRET / ECERP_SERVICE_ID")
		case apiErr.IsRateLimitError():
			fmt.Printf("  ⏳ [%s] 请求频率限制(code=%d): %s\n", operation, apiErr.Code, apiErr.Message)
			fmt.Println("     → 请稍后重试")
		case apiErr.IsServerError():
			fmt.Printf("  💥 [%s] 服务器错误(code=%d): %s\n", operation, apiErr.Code, apiErr.Message)
		default:
			fmt.Printf("  ⚠️  [%s] 业务错误(code=%d): %s\n", operation, apiErr.Code, apiErr.Message)
		}
	} else {
		fmt.Printf("  ❌ [%s] 系统错误: %v\n", operation, err)
	}
}

// getEnvOrDefault 从环境变量获取值，不存在则使用默认值
func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
