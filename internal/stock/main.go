package main

import (
	"log"

	"github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/discovery"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/leebrouse/Gorder/common/logging"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/leebrouse/Gorder/stock/ports"
	"github.com/leebrouse/Gorder/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Viper init
func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	serviceName := viper.GetString("stock.service-name")
	//serviceType
	serviceType := viper.GetString("stock.server-to-run")
	//log
	logrus.Info(serviceType)
	//create application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)

	//register service
	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

	// choose the service type
	switch serviceType {
	//1.grpc
	case "grpc":
		//Run grpc server
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	//2.Run Http server
	case "http":
		//	TODO:
	//3.default (panic )
	default:
		panic("Unexpected server type")
	}

}
