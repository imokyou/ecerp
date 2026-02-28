# ecerp - 易仓ERP Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/imokyou/ecerp.svg)](https://pkg.go.dev/github.com/imokyou/ecerp)

企业级易仓ERP开放平台 Go SDK，模块化子包架构，覆盖 **342 个 API 接口**（全部经官方文档逐一验证）。

## 特性

- 🔐 MD5 自动签名
- 📦 22 个模块化子包
- 🎯 342 个 API 接口全覆盖（官方文档验证）
- 🔒 完整的 Go 类型定义
- ⚡ Context 超时/取消支持
- 🛡️ 企业级：参数校验、OOM 防护、连接池、slog 日志
- 🔄 客户端并发安全，支持复用

## 安装

```bash
go get github.com/imokyou/ecerp
```

## 快速开始

```go
package main

import (
    "context"
    "log"

    "github.com/imokyou/ecerp"
    "github.com/imokyou/ecerp/order"
    "github.com/imokyou/ecerp/inventory"
)

func main() {
    client, err := ecerp.NewClient("app_key", "app_secret", "service_id")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()

    // 查询订单
    orderSvc := order.NewService(client)
    orders, _ := orderSvc.GetOrderList(ctx, &order.GetOrderListRequest{
        PageRequest: order.PageRequest{Page: 1, PageSize: 20},
    })
    _ = orders

    // 查询库存
    invSvc := inventory.NewService(client)
    inv, _ := invSvc.GetProductInventoryNew(ctx, &inventory.GetProductInventoryRequest{
        ProductSKU: "SKU001",
    })
    _ = inv
}
```

## 全部模块 (22 个子包, 342 个接口)

| 子包 | 说明 | 接口数 |
|------|------|--------|
| `product` | 产品查询/创建/编辑、SPU、品牌、品类、条码、箱规、物流/海关属性、图片 | **46** |
| `order` | 订单全生命周期：创建/审核/取消/退件/RMA/拣货/追踪/拦截/面单/结算 | **40** |
| `basicdata` | 营业执照、工单、物流地址、PDA权限、角色权限、产品等级、部门、组织、品类、QC、包材、供应商 | **34** |
| `goods` | 移除订单、Amazon报告(结算/仓储/赔偿/库存/退货)、Listing(Amazon/Walmart/Wayfair)、Review/Feedback | **34** |
| `firstmile` | 头程总单/备货计划、FBA货件/发货管理、海外仓头程、订柜/装箱/拣货/出货、费用 | **26** |
| `fbasta` | FBA STA 入库计划、装箱、分仓、运输、标签、配送窗口 | **20** |
| `purchase` | 采购单/申购单(新旧)、QC/收货异常、附件、变更、ETA、1688匹配、批量导入 | **20** |
| `finance` | 财务利润报表、付款管理、库存分类账、FBA报告(进销存/库存/调整/销售/退货/移除)、交易明细 | **20** |
| `inventory` | 库存查询(新版/团队/库位/FBA)、批次库存、盘点、统计、导入、分销同步、库龄、库位 | **19** |
| `expense` | 付款单(确认/收款/付款)、采购结算、FNSKU成本、费用试算、订单成本明细、服务商账单 | **14** |
| `inbound` | 入库单同步/查询、收货/质检、上架、退回入库、装箱信息、仓库成本、收货明细 | **12** |
| `warehouse` | 仓库信息(全部/单个/运输方式)、库位/库位类型/分区(查询/创建) | **11** |
| `provider` | 服务商订单/库存/标签/状态、上传跟踪号、费用更新 | **8** |
| `supplier` | 供应商信息同步/查询、供应商产品、KPI列表、承运商 | **7** |
| `amazonads` | 广告赔偿/原始报告、发票下载、广告明细、任务状态、店铺列表 | **6** |
| `outbound` | 出库单同步/查询、新品入库、次品上架、出库明细 | **6** |
| `user` | 用户创建/编辑(单个/批量)、平台账号列表 | **5** |
| `compass` | 订单统计(V1/V2)、产品销量汇总/销售情况 | **4** |
| `transfer` | 调拨单创建/查询/编辑 | **3** |
| `packing` | 快捷组包揽件、组包条件、揽件单标签 | **3** |
| `fba` | FBA 退件、FBA 赔偿 | **2** |
| `exchange` | 汇率编辑/查询 | **2** |

## 错误处理

```go
if apiErr, ok := ecerp.IsAPIError(err); ok {
    switch {
    case apiErr.IsAuthError():
        log.Fatal("认证失败")
    case apiErr.IsRateLimitError():
        time.Sleep(time.Second)
    case apiErr.IsServerError():
        log.Println("服务端错误，稍后重试")
    }
}
```

## 配置

```go
client, _ := ecerp.NewClient("key", "secret", "service",
    ecerp.WithTimeout(60*time.Second),
    ecerp.WithLogger(slog.Default()),
    ecerp.WithUserAgent("my-app/1.0"),
)
defer client.Close()
```

## License

MIT License
