package order

import "github.com/leebrouse/Gorder/common/genproto/orderpb"

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*orderpb.Item
}
