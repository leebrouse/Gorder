package service

import (
	"context"
	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/order/adapters"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

// Order application
func NewApplication(ctx context.Context) app.Application {
	//init orderRepo
	orderRepo := adapters.NewMemoryOrderRepository()
	//init logger
	logger := logrus.NewEntry(logrus.StandardLogger())
	//init metricsClient
	metricsClient := metrics.NewTodoMetrics()

	return app.Application{

		Commend: app.Commend{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricsClient),
		},
		Queries: app.Queries{GetCustomOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient)},
	}
}
