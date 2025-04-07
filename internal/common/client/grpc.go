package client

import (
	"github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/genproto/stockpb"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// init viper
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

// New stock GRPC client
func NewStockGRPCClient(ctx context.Context) (client stockpb.StockServiceClient, close func() error, err error) {
	//	read from the config file
	grpcAddr := viper.GetString("stock.grpc-addr")
	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error {
			return nil
		}, err
	}
	conn, err := grpc.NewClient(grpcAddr, opts...)
	if err != nil {
		return nil, func() error {
			return nil
		}, err
	}
	return stockpb.NewStockServiceClient(conn), conn.Close, nil

	//grpc.DialOption()
}

func grpcDialOpts(addr string) ([]grpc.DialOption, error) {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
}
