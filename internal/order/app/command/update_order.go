package command

import (
	"github.com/leebrouse/Gorder/common/decorator"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// UpdateOrder 定义更新订单的命令结构
// 采用函数式更新模式，支持原子化更新操作
type UpdateOrder struct {
	Order    *domain.Order                                               // 需要更新的订单对象
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error) // 更新函数，定义如何修改订单
}

// UpdateOrderHandler 是经过装饰器包装的命令处理器接口
// 使用泛型定义，any表示不需要返回具体结果
type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, any]

// updateOrderHandler 实际的命令处理器实现
type updateOrderHandler struct {
	orderRepo domain.Repository // 订单领域仓储接口
	// stockGRPC 可以添加库存服务客户端（当前未实现）
}

// Handle 实现命令处理的核心逻辑
func (c updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (any, error) {
	// 防御性检查：如果UpdateFn为空则使用默认空操作函数
	if cmd.UpdateFn == nil {
		logrus.Warnf("updateOrderHandler got nil updateFn, orderID=%v", cmd.Order.ID)
		cmd.UpdateFn = func(_ context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil // 默认返回原订单不做修改
		}
	}

	// 调用仓储执行更新操作
	// UpdateFn允许在仓储事务内执行原子化更新
	err := c.orderRepo.Update(ctx, cmd.Order, cmd.UpdateFn)
	if err != nil {
		return nil, err // 更新失败返回错误
	}

	// 成功时返回nil（any类型的零值）
	return nil, nil
}

// NewUpdateOrderHandler 创建装饰后的命令处理器
// 参数:
//   - orderRepo: 订单仓储实现（必须非nil）
//   - logger: 日志记录器
//   - metricClient: 监控指标客户端
//
// 返回:
//   - 经过日志、监控等装饰器包装的命令处理器
func NewUpdateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) UpdateOrderHandler {
	// 启动时依赖检查
	if orderRepo == nil {
		panic("nil orderRepo") // 快速失败机制
	}

	// 应用装饰器链（包含日志、监控等）
	return decorator.ApplyCommandDecorators[UpdateOrder, any](
		updateOrderHandler{orderRepo: orderRepo},
		logger,
		metricClient,
	)
}
