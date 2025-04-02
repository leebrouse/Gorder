package ports

import (
	context "context"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/leebrouse/Gorder/stock/app"
)

type GRPCServer struct {
	app app.Application
}

// New GRPCServer factory pattern
func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (s GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
