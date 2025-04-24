package main

import (
	"github.com/labstack/gommon/log"
	"github.com/leebrouse/Gorder/common/broker"
	"github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/logging"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/leebrouse/Gorder/common/tracing"
	"github.com/leebrouse/Gorder/payment/infrastructure/consumer"
	"github.com/leebrouse/Gorder/payment/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// Viper init
func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}
func main() {
	serviceName := viper.GetString("payment.service-name")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start jaeger for Distributed Tracing in the Payment server
	shutdown, err := tracing.InitJaegerProvider(viper.GetString("jaeger.url"), serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shutdown(ctx)

	serverType := viper.GetString("payment.server-to-run")

	application, cleanup := service.NewApplication(ctx)
	defer cleanup()

	//init rabbitmq
	ch, closeCh := broker.Connect(
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetString("rabbitmq.port"),
	)
	defer func() {
		_ = ch.Close()
		_ = closeCh()
	}()
	//listen the rabbitmq for consuming the message in goroutines
	go consumer.NewConsumer(application).Listen(ch)

	//no register the payment service
	paymentHandler := NewPaymentHandler(ch)
	switch serverType {
	case "http":
		server.RunHTTPServer(serviceName, paymentHandler.RegisterRoutes)
	case "grpc":
		logrus.Panic("unsupported type")
	default:
		logrus.Panic("Unexpected server type")
	}
}
