// Package main 展示多模块组合使用：完整的亚马逊订单处理流程
//
// 本示例模拟一个完整的跨境电商业务流程:
// 1. 查询订单 → 2. 查库存 → 3. 审核订单 → 4. 查运单 → 5. FBA管理
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/fba"
	"github.com/imokyou/ecerp/goods"
	"github.com/imokyou/ecerp/inventory"
	"github.com/imokyou/ecerp/order"
)

func main() {
	// 创建客户端
	client := ecerp.MustNewClient(
		os.Getenv("ECERP_APP_KEY"),
		os.Getenv("ECERP_APP_SECRET"),
		os.Getenv("ECERP_SERVICE_ID"),
		ecerp.WithTimeout(30*time.Second),
		ecerp.WithLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))),
	)
	defer client.Close()

	ctx := context.Background()

	// ═════════════════════════════
	// 场景：Amazon 订单处理流水线
	// ═════════════════════════════

	fmt.Println("🚀 === Amazon 订单处理流水线 ===")
	fmt.Println()

	// Step 1: 获取待处理订单
	fmt.Println("📋 Step 1: 获取待处理订单...")
	orderSvc := order.NewService(client)
	orders, err := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 50},
	})
	if err != nil {
		log.Fatalf("获取订单失败: %v", err)
	}
	fmt.Printf("   → 获取到 %d 条待处理订单\n\n", len(orders))

	// Step 2: 检查库存
	fmt.Println("📦 Step 2: 检查产品库存...")
	invSvc := inventory.NewService(client)
	invList, err := invSvc.GetProductInventoryNew(ctx, &inventory.GetProductInventoryRequest{
		PageRequest: inventory.PageRequest{Page: 1, PageSize: 100},
	})
	if err != nil {
		log.Fatalf("查询库存失败: %v", err)
	}
	fmt.Printf("   → 获取到 %d 条库存记录\n\n", len(invList))

	// Step 3: 查询Amazon Listing
	fmt.Println("🏷️  Step 3: 查询Amazon Listing...")
	goodsSvc := goods.NewService(client)
	listings, err := goodsSvc.AmazonListing(ctx, &goods.ReportRequest{
		PageRequest: goods.PageRequest{Page: 1, PageSize: 20},
	})
	if err != nil {
		handleErr("查询Listing", err)
	} else {
		fmt.Printf("   → 获取到 %d 条Listing\n\n", len(listings))
	}

	// Step 4: 查询FBA退件
	fmt.Println("📤 Step 4: 查询FBA退件记录...")
	fbaSvc := fba.NewService(client)
	returns, err := fbaSvc.GetFbaReturn(ctx, &fba.GetFbaReturnRequest{
		PageRequest: fba.PageRequest{Page: 1, PageSize: 20},
	})
	if err != nil {
		handleErr("查询FBA退件", err)
	} else {
		fmt.Printf("   → 获取到 %d 条FBA退件\n\n", len(returns))
	}

	// Step 5: 查询FBA赔偿
	fmt.Println("💰 Step 5: 查询FBA赔偿...")
	reimbursements, err := fbaSvc.GetFbaReimbursement(ctx, &fba.GetFbaReimbursementRequest{
		PageRequest: fba.PageRequest{Page: 1, PageSize: 20},
	})
	if err != nil {
		handleErr("查询FBA赔偿", err)
	} else {
		fmt.Printf("   → 获取到 %d 条FBA赔偿记录\n\n", len(reimbursements))
	}

	fmt.Println("✅ === 流水线执行完毕 ===")
}

func handleErr(op string, err error) {
	if apiErr, ok := ecerp.IsAPIError(err); ok {
		fmt.Printf("   ⚠️  [%s] API错误(code=%d): %s\n\n", op, apiErr.Code, apiErr.Message)
	} else {
		fmt.Printf("   ❌ [%s] 错误: %v\n\n", op, err)
	}
}
