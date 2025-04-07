package service

import (
	"context"
	grpcClient "github.com/leebrouse/Gorder/common/client"
	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/order/adapters"
	"github.com/leebrouse/Gorder/order/adapters/grpc"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

// Order application
func NewApplication(ctx context.Context) (app.Application, func()) {
	//init
	stockClient, closeStockClient, err := grpcClient.NewStockGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	stockGRPC := grpc.NewStockGRPC(stockClient)
	return newApplication(ctx, stockGRPC), func() {
		_ = closeStockClient()
	}

}

func newApplication(_ context.Context, stockGRPC query.StockService) app.Application {
	//init orderRepo && stockGRPC
	orderRepo := adapters.NewMemoryOrderRepository()
	//stockGRPC := grpc.NewStockGRPC()
	//init logger
	logger := logrus.NewEntry(logrus.StandardLogger())
	//init metricsClient
	metricsClient := metrics.NewTodoMetrics()
	return app.Application{
		Commend: app.Commend{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, stockGRPC, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricsClient),
		},
		Queries: app.Queries{GetCustomOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient)},
	}
}
