# ecerp SDK 使用指南（保姆级）

> 易仓ERP开放平台 Go SDK，覆盖 **22 个模块、342 个 API 接口**。
> 本文档从零开始，手把手教你接入。

---

## 目录

1. [准备工作](#1-准备工作)
2. [安装](#2-安装)
3. [快速上手（5分钟）](#3-快速上手5分钟)
4. [核心概念](#4-核心概念)
5. [客户端配置详解](#5-客户端配置详解)
6. [错误处理](#6-错误处理)
7. [分页查询](#7-分页查询)
8. [常见业务场景](#8-常见业务场景)
9. [生产环境最佳实践](#9-生产环境最佳实践)
10. [FAQ](#10-faq)

---

## 1. 准备工作

### 1.1 获取 API 凭证

在使用 SDK 之前，你需要从 **易仓ERP开放平台** 获取三个凭证：

| 参数 | 说明 | 获取位置 |
|------|------|----------|
| `app_key` | 应用Key | 应用管理 → 应用详情 |
| `app_secret` | 应用密钥 | 应用管理 → 应用详情 |
| `service_id` | 授权服务ID | 授权管理 → 服务授权 |

> [!IMPORTANT]
> `app_secret` 是签名密钥，**绝对不要**提交到代码仓库！建议通过环境变量或密钥管理服务注入。

### 1.2 环境要求

- Go 1.18+（支持泛型）
- 网络：可访问 `openapi-web.eccang.com`

---

## 2. 安装

```bash
go get github.com/imokyou/ecerp
```

---

## 3. 快速上手（5分钟）

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/imokyou/ecerp"
    "github.com/imokyou/ecerp/order"
)

func main() {
    // 第一步：创建客户端
    client, err := ecerp.NewClient(
        os.Getenv("ECERP_APP_KEY"),     // 从环境变量读取
        os.Getenv("ECERP_APP_SECRET"),
        os.Getenv("ECERP_SERVICE_ID"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close() // 别忘了关闭！

    // 第二步：创建模块服务
    orderSvc := order.NewService(client)

    // 第三步：调用 API
    ctx := context.Background()
    orders, err := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
        PageRequest: order.PageRequest{Page: 1, PageSize: 10},
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("获取到 %d 条订单\n", len(orders))
}
```

**就这么简单！三步搞定。**

---

## 4. 核心概念

### 4.1 架构图

```
┌─────────────────────────────────────────────┐
│                你的业务代码                    │
├──────┬──────┬──────┬──────┬──────┬───────────┤
│order │product│inven-│first │pur-  │  ... 共22  │
│      │       │tory  │mile  │chase │  个模块    │
├──────┴──────┴──────┴──────┴──────┴───────────┤
│          ecerp.Client (核心客户端)              │
│  • 自动签名 (MD5)                              │
│  • 连接池管理                                   │
│  • 超时/取消控制                                │
│  • 错误处理                                     │
│  • 日志记录                                     │
├─────────────────────────────────────────────┤
│          HTTP POST → eccang.com               │
└─────────────────────────────────────────────┘
```

### 4.2 三层结构

| 层级 | 说明 | 示例 |
|------|------|------|
| **Client 层** | 管理凭证、签名、连接 | `ecerp.NewClient(...)` |
| **Service 层** | 每个模块一个 Service | `order.NewService(client)` |
| **方法层** | 每个 API 一个 Go 方法 | `orderSvc.GetOrderList(ctx, req)` |

### 4.3 设计原则

- **一个客户端，多个服务**：创建一个 `Client`，可同时用于所有 22 个模块
- **并发安全**：`Client` 和所有 `Service` 都是并发安全的
- **函数式选项**：通过 `With*` 函数配置客户端

---

## 5. 客户端配置详解

### 5.1 最简配置

```go
client, err := ecerp.NewClient("key", "secret", "service_id")
```

### 5.2 全量配置

```go
import (
    "log/slog"
    "net/http"
    "time"

    "github.com/imokyou/ecerp"
)

client, err := ecerp.NewClient(
    "key", "secret", "service_id",

    // 超时设置（默认30秒）
    ecerp.WithTimeout(60 * time.Second),

    // 自定义日志（默认无日志）
    ecerp.WithLogger(slog.Default()),

    // 自定义 User-Agent（默认 ecerp-go-sdk/1.0）
    ecerp.WithUserAgent("my-app/2.0"),

    // 自定义 API 地址（默认正式环境）
    ecerp.WithBaseURL("http://sandbox.eccang.com/openApi/api/unity"),

    // 自定义 HTTP 客户端（高级：代理/证书等）
    ecerp.WithHTTPClient(&http.Client{
        Timeout: 90 * time.Second,
        Transport: &http.Transport{
            MaxIdleConnsPerHost: 20,
        },
    }),
)
```

### 5.3 MustNewClient（初始化阶段推荐）

```go
// 失败时直接 panic，适用于 main() 或 init() 中
client := ecerp.MustNewClient("key", "secret", "service_id",
    ecerp.WithTimeout(60 * time.Second),
)
```

### 5.4 配置项一览

| 选项 | 默认值 | 说明 |
|------|--------|------|
| `WithTimeout` | 30s | 请求超时时间 |
| `WithLogger` | nil | slog 日志记录器 |
| `WithUserAgent` | `ecerp-go-sdk/1.0` | UA 标识 |
| `WithBaseURL` | 正式环境URL | API 地址 |
| `WithHTTPClient` | 自动创建（带连接池） | 自定义 HTTP 客户端 |
| `WithCharset` | `UTF-8` | 字符集 |
| `WithVersion` | `V1.0.0` | API 版本 |
| `WithSignType` | `MD5` | 签名类型 |

---

## 6. 错误处理

### 6.1 基础用法

```go
orders, err := orderSvc.GetOrderList(ctx, req)
if err != nil {
    // 判断是否是 API 业务错误（有错误码）
    if apiErr, ok := ecerp.IsAPIError(err); ok {
        fmt.Printf("API错误: code=%d, msg=%s\n", apiErr.Code, apiErr.Message)
    } else {
        // 网络错误、超时、JSON解析错误等
        fmt.Printf("系统错误: %v\n", err)
    }
    return
}
```

### 6.2 按错误类型处理

```go
if apiErr, ok := ecerp.IsAPIError(err); ok {
    switch {
    case apiErr.IsAuthError():       // 401/403 - 认证失败
        log.Fatal("请检查 appKey/appSecret/serviceID 是否正确")

    case apiErr.IsRateLimitError():  // 429 - 频率限制
        time.Sleep(time.Second)      // 等一秒重试
        // retry...

    case apiErr.IsNotFound():        // 404 - 资源不存在
        log.Println("未找到该资源")

    case apiErr.IsServerError():     // 500+ - 服务端错误
        log.Println("易仓服务器异常，稍后重试")

    default:
        log.Printf("业务错误: %d - %s\n", apiErr.Code, apiErr.Message)
    }
}
```

### 6.3 常见错误码

| 错误码 | 含义 | 处理建议 |
|--------|------|----------|
| 200 | 成功 | 正常处理 |
| 400 | 参数错误 | 检查请求参数 |
| 401 | 未授权 | 检查凭证 |
| 403 | 无权限 | 联系管理员开通权限 |
| 404 | 资源不存在 | 检查 ID 是否正确 |
| 429 | 请求频率限制 | 降低请求频率 |
| 500 | 服务器内部错误 | 稍后重试 |

---

## 7. 分页查询

大部分列表接口都支持分页，统一使用 `PageRequest`：

```go
// 第一页
orders, err := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
    PageRequest: order.PageRequest{
        Page:     1,    // 页码，从 1 开始
        PageSize: 50,   // 每页条数（建议 ≤ 100）
    },
})

// 遍历所有页
page := 1
for {
    orders, err := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
        PageRequest: order.PageRequest{Page: page, PageSize: 100},
    })
    if err != nil {
        log.Fatal(err)
    }
    if len(orders) == 0 {
        break // 没有更多数据了
    }

    // 处理当前页数据
    for _, o := range orders {
        processOrder(o)
    }

    page++
}
```

---

## 8. 常见业务场景

### 8.1 订单管理

```go
orderSvc := order.NewService(client)

// 查询订单列表
orders, _ := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
    PageRequest: order.PageRequest{Page: 1, PageSize: 20},
})

// 创建/同步订单
err := orderSvc.SyncOrder(ctx, &order.SyncOrderRequest{
    // 填充订单信息...
})

// 审核订单
err = orderSvc.VerifyOrder(ctx, &order.VerifyOrderRequest{
    OrderCode: "ORD-2024-001",
})
```

### 8.2 产品管理

```go
productSvc := product.NewService(client)

// 按SKU查询产品
p, _ := productSvc.GetProductBySku(ctx, &product.GetProductBySkuRequest{
    ProductSKU: "SKU001",
})

// 创建产品
err := productSvc.CreateProduct(ctx, &product.CreateProductRequest{
    ProductSKU:  "NEW-SKU-001",
    ProductName: "新产品",
})
```

### 8.3 库存管理

```go
invSvc := inventory.NewService(client)

// 查询库存
inv, _ := invSvc.GetProductInventoryNew(ctx, &inventory.GetProductInventoryRequest{
    ProductSKU: "SKU001",
})

// 导入库存
err := invSvc.ImportInventory(ctx, &inventory.ImportInventoryRequest{
    // ...
})
```

### 8.4 采购管理

```go
purchaseSvc := purchase.NewService(client)

// 创建采购单
err := purchaseSvc.SyncPurchaseOrders(ctx, &purchase.SyncPurchaseOrderRequest{
    WarehouseCode: "WH001",
    SupplierCode:  "SUP001",
    Currency:      "CNY",
    Items: []purchase.PurchaseOrderItem{
        {SKU: "SKU001", Quantity: 100, UnitPrice: 9.99},
        {SKU: "SKU002", Quantity: 200, UnitPrice: 19.99},
    },
})
```

### 8.5 FBA STA 管理

```go
fbastaSvc := fbasta.NewService(client)

// 创建入库计划
err := fbastaSvc.CreateInboundPlan(ctx, &fbasta.CreateInboundPlanRequest{
    // ...
})

// 获取标签
labels, _ := fbastaSvc.GetLabels(ctx, &fbasta.GetLabelsRequest{
    ShipmentID: "SHIP001",
})
```

---

## 9. 生产环境最佳实践

### 9.1 客户端复用

```go
// ✅ 正确：全局创建一个客户端
var client *ecerp.Client

func init() {
    client = ecerp.MustNewClient(
        os.Getenv("ECERP_APP_KEY"),
        os.Getenv("ECERP_APP_SECRET"),
        os.Getenv("ECERP_SERVICE_ID"),
        ecerp.WithTimeout(60 * time.Second),
        ecerp.WithLogger(slog.Default()),
    )
}

// ❌ 错误：每次请求创建新客户端（浪费连接池）
func handleRequest() {
    client, _ := ecerp.NewClient(...)  // 不要这样做！
}
```

### 9.2 Context 超时控制

```go
// 为每个请求设置独立的超时
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

orders, err := orderSvc.GetOrderList(ctx, req)
```

### 9.3 优雅关闭

```go
func main() {
    client := ecerp.MustNewClient(...)
    defer client.Close() // 程序退出时释放连接

    // 或者监听信号
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigCh
        client.Close()
        os.Exit(0)
    }()
}
```

### 9.4 环境变量管理凭证

```bash
# .env 文件（不要提交到 Git！）
export ECERP_APP_KEY="your_key"
export ECERP_APP_SECRET="your_secret"
export ECERP_SERVICE_ID="your_service_id"
```

```go
// .gitignore
.env
```

### 9.5 日志开启

```go
// 开发环境：开启 Debug 日志
handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})
logger := slog.New(handler)

client, _ := ecerp.NewClient("key", "secret", "sid",
    ecerp.WithLogger(logger),
)
// 每次 API 调用都会输出请求/响应日志
```

---

## 10. FAQ

### Q: 支持哪些 Go 版本？
Go 1.18+（因为使用了泛型 `PageResponse[T]`）。

### Q: SDK 是线程安全的吗？
是的，`Client` 和所有 `Service` 都是并发安全的，可以在多个 goroutine 中共享使用。

### Q: 如何使用沙箱环境测试？
```go
client, _ := ecerp.NewClient("key", "secret", "sid",
    ecerp.WithBaseURL("http://sandbox.eccang.com/openApi/api/unity"),
)
```

### Q: 签名是自动处理的吗？
是的，SDK 会自动为每次请求生成 MD5 签名，你不需要关心签名逻辑。

### Q: 如何处理大量数据的分页查询？
参考 [分页查询](#7-分页查询) 章节，使用循环遍历所有页。

### Q: 如何查看完整的请求/响应？
使用 `DoRaw` 方法获取原始响应：
```go
resp, err := client.DoRaw(ctx, "getOrderList", reqBody)
fmt.Printf("Code: %d, Data: %s\n", resp.Code, string(resp.Data))
```

### Q: 接口返回的数据结构在哪里看？
每个模块的 `.go` 文件中都有完整的结构体定义，对应官方文档的请求/响应参数。

---

## 22 个模块速查表

| 模块 | 导入路径 | 接口数 |
|------|---------|:------:|
| 产品 | `github.com/imokyou/ecerp/product` | 46 |
| 订单 | `github.com/imokyou/ecerp/order` | 40 |
| 基础数据 | `github.com/imokyou/ecerp/basicdata` | 34 |
| 商品 | `github.com/imokyou/ecerp/goods` | 34 |
| 头程 | `github.com/imokyou/ecerp/firstmile` | 26 |
| FBA STA | `github.com/imokyou/ecerp/fbasta` | 20 |
| 采购 | `github.com/imokyou/ecerp/purchase` | 20 |
| 财务 | `github.com/imokyou/ecerp/finance` | 20 |
| 库存 | `github.com/imokyou/ecerp/inventory` | 19 |
| 费用 | `github.com/imokyou/ecerp/expense` | 14 |
| 入库 | `github.com/imokyou/ecerp/inbound` | 12 |
| 仓库 | `github.com/imokyou/ecerp/warehouse` | 11 |
| 服务商 | `github.com/imokyou/ecerp/provider` | 8 |
| 供应商 | `github.com/imokyou/ecerp/supplier` | 7 |
| 亚马逊广告 | `github.com/imokyou/ecerp/amazonads` | 6 |
| 出库 | `github.com/imokyou/ecerp/outbound` | 6 |
| 用户 | `github.com/imokyou/ecerp/user` | 5 |
| 数据罗盘 | `github.com/imokyou/ecerp/compass` | 4 |
| 调拨 | `github.com/imokyou/ecerp/transfer` | 3 |
| 组包揽件 | `github.com/imokyou/ecerp/packing` | 3 |
| FBA | `github.com/imokyou/ecerp/fba` | 2 |
| 汇率 | `github.com/imokyou/ecerp/exchange` | 2 |
