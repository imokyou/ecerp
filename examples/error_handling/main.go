// Package main 展示错误处理的最佳实践
//
// 本示例演示:
// 1. 区分 API 业务错误 vs 系统错误
// 2. 按错误码分类处理
// 3. 带重试的错误处理
// 4. 使用 DoRaw 调试
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
	// 示例 1: 基础错误处理
	// ═════════════════════════════
	fmt.Println("=== 示例 1: 基础错误处理 ===")
	basicErrorHandling(ctx, orderSvc)

	// ═════════════════════════════
	// 示例 2: 带重试的调用
	// ═════════════════════════════
	fmt.Println("\n=== 示例 2: 带重试的调用 ===")
	retryExample(ctx, orderSvc)

	// ═════════════════════════════
	// 示例 3: 使用 DoRaw 调试
	// ═════════════════════════════
	fmt.Println("\n=== 示例 3: DoRaw 调试模式 ===")
	debugWithDoRaw(ctx, client)
}

// basicErrorHandling 基础错误处理
func basicErrorHandling(ctx context.Context, svc *order.Service) {
	orders, err := svc.GetOrderList(ctx, &order.GetOrderListRequest{
		PageRequest: order.PageRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		// 第一层：判断是否是 API 业务错误
		apiErr, isAPI := ecerp.IsAPIError(err)
		if !isAPI {
			// 不是 API 错误 → 网络/超时/JSON解析等系统错误
			log.Printf("❌ 系统错误: %v\n", err)
			return
		}

		// 第二层：按错误码分类处理
		switch {
		case apiErr.IsAuthError():
			// 401 / 403
			fmt.Println("🔐 认证失败！请检查:")
			fmt.Println("   1. app_key 是否正确")
			fmt.Println("   2. app_secret 是否正确")
			fmt.Println("   3. service_id 是否有此接口权限")
			fmt.Println("   4. 应用是否已上线/审批通过")

		case apiErr.IsRateLimitError():
			// 429
			fmt.Println("⏳ 请求频率超限，请降低调用频率")

		case apiErr.IsNotFound():
			// 404
			fmt.Println("🔍 资源不存在")

		case apiErr.IsServerError():
			// 500+
			fmt.Println("💥 易仓服务器内部错误，建议稍后自动重试")

		default:
			// 其他业务错误 (如参数验证失败等)
			fmt.Printf("⚠️  业务错误 [%d]: %s\n", apiErr.Code, apiErr.Message)
		}
		return
	}

	fmt.Printf("✅ 成功获取 %d 条订单\n", len(orders))
}

// retryExample 带重试的错误处理
func retryExample(ctx context.Context, svc *order.Service) {
	maxRetries := 3
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		orders, err := svc.GetOrderList(ctx, &order.GetOrderListRequest{
			PageRequest: order.PageRequest{Page: 1, PageSize: 10},
		})
		if err == nil {
			fmt.Printf("✅ 第 %d 次尝试成功，获取 %d 条订单\n", attempt, len(orders))
			return
		}

		lastErr = err

		// 判断是否应该重试
		apiErr, isAPI := ecerp.IsAPIError(err)
		if isAPI {
			switch {
			case apiErr.IsAuthError():
				// 认证错误不重试（重试也没用）
				fmt.Println("🔐 认证错误，不再重试")
				return

			case apiErr.IsRateLimitError():
				// 限流：等久一点再试
				wait := time.Duration(attempt) * 2 * time.Second
				fmt.Printf("⏳ 限流，等待 %v 后重试 (第 %d/%d 次)\n", wait, attempt, maxRetries)
				time.Sleep(wait)

			case apiErr.IsServerError():
				// 服务器错误：指数退避重试
				wait := time.Duration(attempt) * time.Second
				fmt.Printf("💥 服务器错误，等待 %v 后重试 (第 %d/%d 次)\n", wait, attempt, maxRetries)
				time.Sleep(wait)

			default:
				// 其他业务错误不重试
				fmt.Printf("⚠️  业务错误 [%d]: %s，不再重试\n", apiErr.Code, apiErr.Message)
				return
			}
		} else {
			// 网络错误等，短暂等待后重试
			wait := time.Duration(attempt) * 500 * time.Millisecond
			fmt.Printf("❌ 网络错误，等待 %v 后重试 (第 %d/%d 次): %v\n", wait, attempt, maxRetries, err)
			time.Sleep(wait)
		}
	}

	fmt.Printf("❌ 重试 %d 次后仍然失败: %v\n", maxRetries, lastErr)
}

// debugWithDoRaw 使用 DoRaw 调试 API 调用
func debugWithDoRaw(ctx context.Context, client *ecerp.Client) {
	// DoRaw 返回原始响应，适合调试
	resp, err := client.DoRaw(ctx, "getOrderList", map[string]interface{}{
		"page":      1,
		"page_size": 5,
	})

	if err != nil {
		if apiErr, ok := ecerp.IsAPIError(err); ok {
			// 即使有 API 错误，resp 仍然不为 nil（包含原始数据）
			fmt.Printf("API错误: code=%d, msg=%s\n", apiErr.Code, apiErr.Message)
			if resp != nil {
				fmt.Printf("原始响应: %s\n", string(resp.Data))
			}
			return
		}
		log.Printf("请求错误: %v\n", err)
		return
	}

	fmt.Printf("✅ 响应码: %d\n", resp.Code)
	fmt.Printf("   消息: %s\n", resp.Message)

	// 打印原始 JSON 数据（截断显示）
	data := string(resp.Data)
	if len(data) > 500 {
		data = data[:500] + "..."
	}
	fmt.Printf("   数据: %s\n", data)
}
