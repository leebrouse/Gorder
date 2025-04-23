package command

import (
	"encoding/json"
	"errors"

	"github.com/leebrouse/Gorder/common/broker"
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/app/query"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	amqp "github.com/rabbitmq/amqp091-go"
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
	channel   *amqp.Channel
}

// Handle 实现了命令处理的核心逻辑
func (c createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	// TODO: 实际应该通过gRPC调用库存服务验证商品可用性
	validItems, err := c.validate(ctx, cmd.Items)
	if err != nil {
		return nil, err
	}

	//创建order
	o, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID,
		Items:      validItems,
	})
	if err != nil {
		return nil, err
	}

	//OrderCreated queue declare
	q, err := c.channel.QueueDeclare(
		broker.EventOrderCreated,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// marshal the order
	marshalledOrder, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	//send message
	err = c.channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent, //持久化消息
		Body:         marshalledOrder,
	})
	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{OrderID: o.ID}, nil
}

// validate
func (c createOrderHandler) validate(ctx context.Context, items []*orderpb.ItemWithQuantity) ([]*orderpb.Item, error) {
	if len(items) == 0 {
		return nil, errors.New(" must have at least one item")
	}
	items = packItems(items)
	resp, err := c.stockGRPC.CheckIfItemsInStock(ctx, items)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}

//合并相同id的order中的quantity 值如
/**
	{
		"id"=2
		"quantity"=30
	}
	{
		"id"=3
		"quantity"=30
	}
	{
		"id"=2
		"quantity"=20
	}
**/
func packItems(items []*orderpb.ItemWithQuantity) []*orderpb.ItemWithQuantity {
	//用 map 来进行去重合并
	merged := make(map[string]int32)
	for _, item := range items {
		merged[item.ID] += item.Quantity
	}

	//合并order 之后的结果
	var res []*orderpb.ItemWithQuantity
	for id, quantity := range merged {
		res = append(res, &orderpb.ItemWithQuantity{ID: id, Quantity: quantity})
	}
	return res
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
	channel *amqp.Channel,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CreateOrderHandler {
	// 防御性编程：检查必要依赖
	if orderRepo == nil {
		panic("nil orderRepo") // 启动时快速失败
	}

	if stockGRPC == nil {
		panic("nil stockGRPC")
	}

	if channel == nil {
		panic("nil channel")
	}
	// 应用装饰器链（通常包含日志记录、性能监控、重试等）
	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{
			orderRepo: orderRepo,
			stockGRPC: stockGRPC,
			channel:   channel,
		}, // 核心处理器 1.order 2.stock
		logger,       // 日志装饰器
		metricClient, // 监控装饰器
		// 可以添加更多装饰器...
	)
}
