package ports

import (
	context "context"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
)

type GRPCServer struct{}

// New GRPCServer factory pattern
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (s GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
