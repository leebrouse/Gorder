package command

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type OrderService interface {
	UpdateOrder(ctx context.Context, order *orderpb.Order) error
}
