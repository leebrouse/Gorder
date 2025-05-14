package query

import (
	"context"

	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
)

type StockService interface {
	CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error)
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
}
