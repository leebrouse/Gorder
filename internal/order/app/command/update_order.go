package command

import (
	"github.com/leebrouse/Gorder/common/decorator"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// request
type UpdateOrder struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, any]

type updateOrderHandler struct {
	orderRepo domain.Repository
	//stockGRPC
}

// implement CommandHandler
func (c updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (any, error) {
	if cmd.UpdateFn == nil {
		logrus.Warnf("updateOrderHandler got nil updateFn,orderID=%v", cmd.Order)
		cmd.UpdateFn = func(_ context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil
		}
	}
	err := c.orderRepo.Update(ctx, cmd.Order, cmd.UpdateFn)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewUpdateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) UpdateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[UpdateOrder, any](
		updateOrderHandler{orderRepo: orderRepo},
		logger,
		metricClient,
	)
}
