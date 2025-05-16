package service

import (
	"context"

	grpcClient "github.com/leebrouse/Gorder/common/client"
	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/payment/adapters"
	"github.com/leebrouse/Gorder/payment/app"
	"github.com/leebrouse/Gorder/payment/app/command"
	"github.com/leebrouse/Gorder/payment/domain"
	"github.com/leebrouse/Gorder/payment/infrastructure/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	orderClient, closeOrderClient, err := grpcClient.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	//memoryProcessor := processor.NewInmemProcessor()
	stripeProcessor := processor.NewStripeProcessor(viper.GetString("stripe-key"))
	return newApplication(ctx, orderGRPC, stripeProcessor), func() {
		_ = closeOrderClient()
	}
}

func newApplication(_ context.Context, orderGRPC command.OrderService, processor domain.Processor) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NewPrometheusMetricsClient(&metrics.PrometheusMetricsClientConfig{
		Host:        viper.GetString("payment.metrics_http_addr"),
		ServiceName: viper.GetString("payment.service-name"),
	})
	return app.Application{
		Commands: app.Commands{
			CreatePayment: command.NewCreatePaymentHandler(processor, orderGRPC, logger, metricsClient),
		},
	}
}
