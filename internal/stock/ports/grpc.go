package ports

import (
	context "context"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/leebrouse/Gorder/stock/app"
	"github.com/sirupsen/logrus"
)

type GRPCServer struct {
	app app.Application
}

// New GRPCServer factory pattern
func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (s GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	logrus.Info("rpc_request_in, stock_GetItems")
	defer func() {
		logrus.Info("rpc_request_out, stock_GetItems")
	}()

	//Test date
	fake := []*orderpb.Item{
		{
			ID: "fake-item-from-GetItems",
		},
	}
	return &stockpb.GetItemsResponse{Items: fake}, nil
}

func (s GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	logrus.Info("rpc_request_in, stock_CheckIfItemsInStock")
	defer func() {
		logrus.Info("rpc_request_out, stock_CheckIfItemsInStock")
	}()
	return nil, nil
}
