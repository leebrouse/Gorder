package adapters

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/tracing"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type OrderGRPC struct {
	client orderpb.OrderServiceClient
}

func NewOrderGRPC(client orderpb.OrderServiceClient) *OrderGRPC {
	return &OrderGRPC{client: client}
}

func (o OrderGRPC) UpdateOrder(ctx context.Context, order *orderpb.Order) error {
	ctx, span := tracing.Start(ctx, "order_grpc.update_order")
	defer span.End()

	_, err := o.client.UpdateOrder(ctx, order)
	logrus.Infof("payment_adapter||update_order,err=%v", err)
	return err
}
