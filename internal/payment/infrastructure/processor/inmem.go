package processor

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type InmemProcessor struct {
}

func NewInmemProcessor() *InmemProcessor {
	return &InmemProcessor{}
}

func (i InmemProcessor) CreatePaymentLink(ctx context.Context, order *orderpb.Order) (string, error) {
	return "inmem-payment-link", nil
}
