package adapters

import (
	"context"

	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/tracing"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

type OrderGRPC struct {
	client orderpb.OrderServiceClient
}

func NewOrderGRPC(client orderpb.OrderServiceClient) *OrderGRPC {
	return &OrderGRPC{client: client}
}

func (o OrderGRPC) UpdateOrder(ctx context.Context, order *orderpb.Order) (err error) {
	defer func() {
		if err != nil {
			logrus.Infof("payment_adapter||update_order,err=%v", err)
		}
	}()

	ctx, span := tracing.Start(ctx, "order_grpc.update_order")
	defer span.End()

	_, err = o.client.UpdateOrder(ctx, order)
	return status.Convert(err).Err()
}
