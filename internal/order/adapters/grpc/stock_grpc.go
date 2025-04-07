package grpc

import (
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type StockGRPC struct {
	client stockpb.StockServiceClient
}

func NewStockGRPC(client stockpb.StockServiceClient) *StockGRPC {
	return &StockGRPC{client: client}
}

// implement check function
func (s StockGRPC) CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) error {
	resp, err := s.client.CheckIfItemsInStock(ctx, &stockpb.CheckIfItemsInStockRequest{Items: items})
	//log operation:
	logrus.Info("stock_grpc response", resp)
	return err
}

// implement get function
func (s StockGRPC) GetItems(ctx context.Context, itemsIDs []string) ([]*orderpb.Item, error) {
	resp, err := s.client.GetItems(ctx, &stockpb.GetItemsRequest{ItemIDs: itemsIDs})
	if err != nil {
		return nil, err
	}
	//log operation:
	logrus.Info("stock_grpc response", resp)
	return resp.Items, err
}
