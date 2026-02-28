// Package main 展示分页遍历的最佳实践
//
// 当需要获取大量数据时，必须使用分页遍历。
// 本示例展示如何安全、高效地遍历所有页面的数据。
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imokyou/ecerp"
	"github.com/imokyou/ecerp/order"
)

func main() {
	client := ecerp.MustNewClient(
		os.Getenv("ECERP_APP_KEY"),
		os.Getenv("ECERP_APP_SECRET"),
		os.Getenv("ECERP_SERVICE_ID"),
	)
	defer client.Close()

	ctx := context.Background()
	orderSvc := order.NewService(client)

	// ═════════════════════════════
	// 方式一：简单分页遍历
	// ═════════════════════════════
	fmt.Println("📖 方式一：简单分页遍历")
	simplePagination(ctx, orderSvc)

	// ═════════════════════════════
	// 方式二：带超时控制的分页
	// ═════════════════════════════
	fmt.Println("\n⏱️  方式二：带超时控制的分页")
	paginationWithTimeout(ctx, orderSvc)

	// ═════════════════════════════
	// 方式三：带限速的分页（避免429）
	// ═════════════════════════════
	fmt.Println("\n🐌 方式三：带限速的分页")
	paginationWithRateLimit(ctx, orderSvc)
}

// simplePagination 最简单的分页遍历
func simplePagination(ctx context.Context, svc *order.Service) {
	page := 1
	pageSize := 100 // 建议每页 50-100 条
	total := 0

	for {
		orders, err := svc.GetOrderList(ctx, &order.GetOrderListRequest{
			PageRequest: order.PageRequest{
				Page:     page,
				PageSize: pageSize,
			},
		})
		if err != nil {
			log.Printf("第 %d 页查询失败: %v\n", page, err)
			break
		}

		// 没有更多数据了
		if len(orders) == 0 {
			break
		}

		total += len(orders)
		fmt.Printf("  第 %d 页: 获取 %d 条 (累计 %d)\n", page, len(orders), total)

		// 如果返回数量小于 pageSize，说明是最后一页
		if len(orders) < pageSize {
			break
		}

		page++
	}

	fmt.Printf("  ✅ 遍历完成，共 %d 条订单\n", total)
}

// paginationWithTimeout 带总体超时控制的分页
func paginationWithTimeout(ctx context.Context, svc *order.Service) {
	// 总体超时 5 分钟（防止网络问题导致无限等待）
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	page := 1
	total := 0

	for {
		// 检查是否超时
		select {
		case <-ctx.Done():
			fmt.Printf("  ⏱️  超时中断，已获取 %d 条\n", total)
			return
		default:
		}

		orders, err := svc.GetOrderList(ctx, &order.GetOrderListRequest{
			PageRequest: order.PageRequest{Page: page, PageSize: 100},
		})
		if err != nil {
			if ctx.Err() != nil {
				fmt.Println("  ⏱️  请求超时")
			} else {
				log.Printf("  ❌ 第 %d 页失败: %v\n", page, err)
			}
			break
		}

		if len(orders) == 0 {
			break
		}

		total += len(orders)
		page++
	}

	fmt.Printf("  ✅ 完成，共 %d 条\n", total)
}

// paginationWithRateLimit 带限速的分页（避免触发 429）
func paginationWithRateLimit(ctx context.Context, svc *order.Service) {
	page := 1
	total := 0
	interval := 200 * time.Millisecond // 每次请求间隔 200ms（即每秒最多 5 次）

	for {
		orders, err := svc.GetOrderList(ctx, &order.GetOrderListRequest{
			PageRequest: order.PageRequest{Page: page, PageSize: 100},
		})
		if err != nil {
			if apiErr, ok := ecerp.IsAPIError(err); ok && apiErr.IsRateLimitError() {
				// 被限流了，等 2 秒再重试
				fmt.Printf("  ⏳ 被限流，等待 2 秒后重试 (第 %d 页)\n", page)
				time.Sleep(2 * time.Second)
				continue // 重试当前页，不 page++
			}
			log.Printf("  ❌ 第 %d 页失败: %v\n", page, err)
			break
		}

		if len(orders) == 0 {
			break
		}

		total += len(orders)
		fmt.Printf("  第 %d 页: %d 条\n", page, len(orders))
		page++

		// 限速：等待一小段时间再请求下一页
		time.Sleep(interval)
	}

	fmt.Printf("  ✅ 完成，共 %d 条\n", total)
}
