package service

import (
	"context"

	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/stock/adapters"
	"github.com/leebrouse/Gorder/stock/app"
	"github.com/leebrouse/Gorder/stock/app/query"
	"github.com/leebrouse/Gorder/stock/infrastructure/integration"
	"github.com/leebrouse/Gorder/stock/infrastructure/persistent"
	"github.com/sirupsen/logrus"
)

func NewApplication(_ context.Context) app.Application {
	//stockRepo := adapters.NewMemoryStockRepository()
	db := persistent.NewMySQL()
	stockRepo := adapters.NewMySQLStockRepository(db)
	logger := logrus.NewEntry(logrus.StandardLogger())
	stripeAPI := integration.NewStripeAPI()
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, stripeAPI, logger, metricsClient),
			GetItems:            query.NewGetItemsHandler(stockRepo, logger, metricsClient),
		},
	}
}
