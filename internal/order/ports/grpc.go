package ports

import (
	context "context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}
