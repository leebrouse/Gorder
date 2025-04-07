package query

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type StockService interface {
	CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) error
	GetItems(ctx context.Context, itemsIDs []string) ([]*orderpb.Item, error)
}
