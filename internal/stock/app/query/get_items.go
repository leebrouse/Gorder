package query

import (
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/stock/domain/stock"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// GetItems 是查询请求结构体，表示根据一组 ItemIDs 获取对应的库存项
type GetItems struct {
	ItemIDs []string
}

// GetItemsHandler 是一个泛型接口别名，表示处理 GetItems 请求，返回 []*orderpb.Item 的处理器
type GetItemsHandler decorator.QueryHandler[GetItems, []*orderpb.Item]

// getItemsHandler 是真正处理查询逻辑的结构体实现
type getItemsHandler struct {
	stockRepo domain.Repository // 注入的库存仓库接口，实现数据获取
}

// Handle 实现 QueryHandler 接口中的方法，执行从仓库中获取对应 Item 的逻辑
func (g getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*orderpb.Item, error) {
	// 从仓库中查询所有指定 ID 的库存项
	items, err := g.stockRepo.GetItems(ctx, query.ItemIDs)
	if err != nil {
		return nil, err // 查询出错，直接返回错误
	}
	return items, nil // 成功则返回查询结果
}

// NewGetCustomerOrderHandler 构造一个带有装饰器（日志、监控等）的 GetItemsHandler
func NewGetCustomerOrderHandler(
	stockRepo domain.Repository, // 依赖注入的库存仓库接口
	logger *logrus.Entry, // 日志组件（如 logrus）
	metricClient decorator.MetricsClient, // 指标采集组件（用于监控）
) GetItemsHandler {

	// 启动时依赖检查，确保仓库不为 nil，避免后续运行时 panic
	if stockRepo == nil {
		panic("nil stockRepo")
	}

	// 使用装饰器链增强实际的 Handler，实现链式扩展：
	// - 日志记录
	// - 监控埋点
	// - 可能还包括缓存、熔断、限流等功能
	// 这里必须指定泛型类型：请求是 GetItems，响应是 []*orderpb.Item
	return decorator.ApplyQueryDecorators[GetItems, []*orderpb.Item](
		getItemsHandler{stockRepo: stockRepo}, // 实际业务逻辑处理器
		logger,                                // 日志器
		metricClient,                          // 指标采集器
	)
}
