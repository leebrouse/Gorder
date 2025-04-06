package command

import (
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// request
type CreateOrder struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

// response
type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

type createOrderHandler struct {
	orderRepo domain.Repository
	//stockGRPC
}

// implement CommandHandler
func (c createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	//Todo: stock gRPC to get items
	//Test data
	var stockResponse []*orderpb.Item
	for _, item := range cmd.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}
	createOrder, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID,
		Items:      stockResponse,
	})
	if err != nil {
		return nil, err
	}
	return &CreateOrderResult{OrderID: createOrder.ID}, nil
}

func NewCreateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CreateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{orderRepo: orderRepo},
		logger,
		metricClient,
	)
}
