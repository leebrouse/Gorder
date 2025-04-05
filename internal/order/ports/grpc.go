package ports

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/app"
	"log"
)

type GRPCServer struct {
	app app.Application
}

// New GRPCServer factory pattern
func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*empty.Empty, error) {
	//TODO implement me
	//panic("implement me")
	log.Print("CreateOrder Function\n")
	return nil, nil
}

func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	//TODO implement me
	log.Print("GetOrder Function\n")
	return nil, nil
}

func (G GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*empty.Empty, error) {
	//TODO implement me
	log.Print("UpdateOrder Function\n")
	return nil, nil
}
