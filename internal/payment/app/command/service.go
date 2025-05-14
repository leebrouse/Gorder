package command

import (
	"context"

	"github.com/leebrouse/Gorder/common/genproto/orderpb"
)

type OrderService interface {
	UpdateOrder(ctx context.Context, order *orderpb.Order) error
}
