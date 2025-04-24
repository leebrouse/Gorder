package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/broker"
	"github.com/leebrouse/Gorder/common/config"
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
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// init order
func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	//logrus.Fatal(viper.GetString("STRIPE_KEY"))
	serviceName := viper.GetString("order.service-name")

	//create an application with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Start jaeger for Distributed Tracing in the order Server
	shutdown, err := tracing.InitJaegerProvider(viper.GetString("jaeger.url"), serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shutdown(ctx)

	application, cleanup := service.NewApplication(ctx)
	defer cleanup()

	//register service To Consul
	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

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
	//consume the message form the payment service
	go consumer.NewConsumer(application).Listen(ch)

	//Run gRPC Server in goroutine
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	//Run Http Server
	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		router.StaticFile("/success", "../../public/success.html")
		ports.RegisterHandlersWithOptions(router, HTTPServer{application}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
