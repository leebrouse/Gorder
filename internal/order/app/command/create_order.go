package command

import (
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/app/query"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// CreateOrder 定义了创建订单命令的请求结构
type CreateOrder struct {
	CustomerID string                      // 客户ID
	Items      []*orderpb.ItemWithQuantity // 订单项列表（包含商品ID和数量）
}

// CreateOrderResult 定义了创建订单命令的响应结构
type CreateOrderResult struct {
	OrderID string // 创建成功后返回的订单ID
}

// CreateOrderHandler 是经过装饰器包装的命令处理器接口
type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

// createOrderHandler 是实际的命令处理器实现
type createOrderHandler struct {
	orderRepo domain.Repository  // 订单领域仓储接口
	stockGRPC query.StockService //添加库存服务gRPC客户端
}

// Handle 实现了命令处理的核心逻辑
func (c createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	// TODO: 实际应该通过gRPC调用库存服务验证商品可用性
	//Test CheckIfItemsInStock function
	logrus.Info("Test CheckIfItemsInStock")
	err := c.stockGRPC.CheckIfItemsInStock(ctx, cmd.Items)
	//log
	logrus.Info("createOrderHandler||err from stockGRPC", err)

	//Test GetItems function
	logrus.Info("Test CheckIfItemsInStock")
	resp, _ := c.stockGRPC.GetItems(ctx, []string{"123"})
	//log
	logrus.Info("GetItems||resp from stockGRPC", resp)

	// 当前使用测试数据模拟库存响应
	var stockResponse []*orderpb.Item
	for _, item := range cmd.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			ID:       item.ID,       // 商品ID
			Quantity: item.Quantity, // 商品数量
		})
	}

	// 调用仓储创建订单
	createOrder, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID, // 设置客户ID
		Items:      stockResponse,  // 设置订单项
	})
	if err != nil {
		return nil, err // 创建失败时返回错误
	}

	// 返回创建成功的订单ID
	return &CreateOrderResult{OrderID: createOrder.ID}, nil
}

// NewCreateOrderHandler 创建经过装饰的命令处理器
// 参数:
//   - orderRepo: 订单仓储实现（不可为nil）
//   - logger: 日志记录器
//   - metricClient: 指标监控客户端
//
// 返回:
//   - 经过日志、监控等装饰器包装的命令处理器
func NewCreateOrderHandler(
	orderRepo domain.Repository,
	stockGRPC query.StockService,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CreateOrderHandler {
	// 防御性编程：检查必要依赖
	if orderRepo == nil {
		panic("nil orderRepo") // 启动时快速失败
	}
	// 应用装饰器链（通常包含日志记录、性能监控、重试等）
	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{orderRepo: orderRepo, stockGRPC: stockGRPC}, // 核心处理器 1.order 2.stock
		logger,       // 日志装饰器
		metricClient, // 监控装饰器
		// 可以添加更多装饰器...
	)
}
