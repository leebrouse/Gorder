package service

import (
	"github.com/leebrouse/Gorder/common/client"
	"github.com/leebrouse/Gorder/common/metrics"
	"github.com/leebrouse/Gorder/payment/adapters"
	"github.com/leebrouse/Gorder/payment/app"
	"github.com/leebrouse/Gorder/payment/app/command"
	"github.com/leebrouse/Gorder/payment/domain"
	"github.com/leebrouse/Gorder/payment/infrastructure/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// Order application
func NewApplication(ctx context.Context) (app.Application, func()) {
	orderClient, closeOrderClient, err := client.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	//introduce the stripe api
	stripeProcess := processor.NewStripeProcessor(viper.GetString("stripe_key"))
	return newApplication(ctx, orderGRPC, stripeProcess), func() {
		_ = closeOrderClient()
	}
}

func newApplication(_ context.Context, orderGRPC command.OrderService, processor domain.Processor) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NewTodoMetrics()
	return app.Application{
		Commend: app.Commend{CreatePayment: command.NewCreatePaymentHandler(processor, orderGRPC, logger, metricsClient)},
	}
}
