package query

import (
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/stock/domain/stock"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CheckIfItemsInStock struct {
	Items []*orderpb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*orderpb.Item]

// getItemsOrderHandler 实际的查询处理器实现
type checkIfItemsInStockHandler struct {
	stockRepo domain.Repository // 订单领域仓储接口
}

func (h checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*orderpb.Item, error) {
	var res []*orderpb.Item
	for _, i := range query.Items {
		res = append(res, &orderpb.Item{
			ID:       i.ID,
			Quantity: i.Quantity,
		})
	}
	return res, nil
}

func NewCheckIfItemsInStockHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	// 启动时依赖检查
	if stockRepo == nil {
		panic("nil stockRepo") // 快速失败，避免运行时空指针
	}

	// 应用装饰器链（通常包含日志、缓存、监控等）
	return decorator.ApplyQueryDecorators[CheckIfItemsInStock, []*orderpb.Item](
		checkIfItemsInStockHandler{stockRepo: stockRepo},
		logger,
		metricClient,
	)
}
