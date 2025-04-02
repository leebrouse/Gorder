package main

import (
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/leebrouse/Gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	//serviceType
	serviceType := viper.GetString("stock.server-to-run")
	switch serviceType {
	//1.grpc
	case "grpc":
		//Run grpc server
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer()
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
