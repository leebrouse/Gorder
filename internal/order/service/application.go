package service

import (
	"context"

	"github.com/leebrouse/Gorder/common/broker"
	grpcClient "github.com/leebrouse/Gorder/common/client"
	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/order/adapters"
	"github.com/leebrouse/Gorder/order/adapters/grpc"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Order application
func NewApplication(ctx context.Context) (app.Application, func()) {
	//init,order server 需要调用stock server的方法需要注册grpc client
	stockClient, closeStockClient, err := grpcClient.NewStockGRPCClient(ctx)
	if err != nil {
		panic(err)
	}

	//init rabbitmq
	ch, closeCh := broker.Connect(
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetString("rabbitmq.port"),
	)

	stockGRPC := grpc.NewStockGRPC(stockClient)
	return newApplication(ctx, stockGRPC, ch), func() {
		_ = closeStockClient()
		_ = closeCh()
		_ = ch.Close()
	}

}

func newApplication(_ context.Context, stockGRPC query.StockService, ch *amqp.Channel) app.Application {
	//init orderRepo && stockGRPC
	orderRepo := adapters.NewMemoryOrderRepository()
	//stockGRPC := grpc.NewStockGRPC()
	//init logger
	logger := logrus.NewEntry(logrus.StandardLogger())
	//init metricsClient
	metricsClient := metrics.NewTodoMetrics()
	return app.Application{
		Commend: app.Commend{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, stockGRPC, ch, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricsClient),
		},
		Queries: app.Queries{GetCustomOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient)},
	}
}
