package domain

import (
	"context"

	"github.com/leebrouse/Gorder/common/genproto/orderpb"
)

type Processor interface {
	CreatePaymentLink(context.Context, *orderpb.Order) (string, error)
}

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*orderpb.Item
}
