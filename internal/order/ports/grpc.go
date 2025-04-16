package ports

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	app app.Application
}

// New GRPCServer factory pattern
func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

// 业务逻辑层
func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*empty.Empty, error) {
	_, err := G.app.Commend.CreateOrder.Handle(ctx, command.CreateOrder{
		CustomerID: request.CustomerID,
		Items:      request.Item,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &empty.Empty{}, err
}

func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	o, err := G.app.Queries.GetCustomOrder.Handle(ctx, query.GetCustomerOrder{
		CustomerID: request.CustomerID,
		OrderID:    request.OrderID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return o.ToProto(), nil
}

func (G GRPCServer) UpdateOrder(ctx context.Context, request *orderpb.Order) (*emptypb.Empty, error) {
	logrus.Infof("order_grpc||request_in||request=%+v", request)
	//create a new order for updating the order status and paymentLink
	order, err := domain.NewOrder(
		request.ID,
		request.CustomerID,
		request.Status,
		request.PaymentLink,
		request.Items,
	)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	_, err = G.app.Commend.UpdateOrder.Handle(ctx, command.UpdateOrder{
		Order: order,
		UpdateFn: func(ctx context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil
		},
	})
	return nil, err
}
