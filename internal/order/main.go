package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/discovery"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/leebrouse/Gorder/order/ports"
	"github.com/leebrouse/Gorder/order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

// init order
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	//Test viper
	//log.Printf("%v", viper.Get("order"))
	serviceName := viper.GetString("order.service-name")

	//create an application with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	//Run gRPC Server in goroutine
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	//Run Http Server
	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{application}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
