package query

import (
	"github.com/leebrouse/Gorder/common/decorator"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// GetCustomerOrder 定义查询客户订单的请求参数
// 包含客户ID和订单ID两个必要参数
type GetCustomerOrder struct {
	CustomerID string // 客户唯一标识符
	OrderID    string // 订单唯一标识符
}

// GetCustomerOrderHandler 是经过装饰器包装的查询处理器接口
// 泛型参数：[请求类型, 返回类型(domain.Order指针)]
type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

// getCustomerOrderHandler 实际的查询处理器实现
type getCustomerOrderHandler struct {
	orderRepo domain.Repository // 订单领域仓储接口
}

// Handle 实现查询处理的核心逻辑
func (g getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	// 通过仓储获取订单数据
	// 同时验证订单是否属于指定客户（双重校验）
	o, err := g.orderRepo.Get(ctx, query.OrderID, query.CustomerID)
	if err != nil {
		// 错误可能包含：订单不存在/不属于该客户/数据库错误等
		return nil, err
	}
	return o, nil
}

// NewGetCustomerOrderHandler 创建装饰后的查询处理器
// 参数:
//   - orderRepo: 订单仓储实现（必须非nil）
//   - logger: 结构化日志记录器
//   - metricClient: 监控指标客户端
//
// 返回:
//   - 应用了日志、监控等装饰器的查询处理器
func NewGetCustomerOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) GetCustomerOrderHandler {
	// 启动时依赖检查
	if orderRepo == nil {
		panic("nil orderRepo") // 快速失败，避免运行时空指针
	}

	// 应用装饰器链（通常包含日志、缓存、监控等）
	return decorator.ApplyQueryDecorators[GetCustomerOrder, *domain.Order](
		getCustomerOrderHandler{orderRepo: orderRepo},
		logger,
		metricClient,
	)
}
