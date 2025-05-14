package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/broker"
	_ "github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/discovery"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/logging"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/leebrouse/Gorder/common/tracing"
	"github.com/leebrouse/Gorder/order/infrastructure/consumer"
	"github.com/leebrouse/Gorder/order/ports"
	"github.com/leebrouse/Gorder/order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	logging.Init()
}

func main() {
	serviceName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown, err := tracing.InitJaegerProvider(viper.GetString("jaeger.url"), serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shutdown(ctx)

	application, cleanup := service.NewApplication(ctx)
	defer cleanup()

	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

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
	go consumer.NewConsumer(application).Listen(ch)

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		router.StaticFile("/success", "../../public/success.html")
		ports.RegisterHandlersWithOptions(router, HTTPServer{
			app: application,
		}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}
