package domain

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type Processor interface {
	CreatePaymentLink(ctx context.Context, order *orderpb.Order) (string, error)
}
