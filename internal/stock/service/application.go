package service

import (
	"context"

	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/stock/adapters"
	"github.com/leebrouse/Gorder/stock/app"
	"github.com/leebrouse/Gorder/stock/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	//init orderRepo && stockGRPC
	stockRepo := adapters.NewMemoryStockRepository()
	//stockGRPC := grpc.NewStockGRPC()
	//init logger
	logger := logrus.NewEntry(logrus.StandardLogger())
	//init metricsClient
	metricsClient := metrics.NewTodoMetrics()
	return app.Application{
		Commend: app.Commend{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, logger, metricsClient),
			GetItems:            query.NewGetCustomerOrderHandler(stockRepo, logger, metricsClient),
		},
	}
}
