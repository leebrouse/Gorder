package main

import (
	"github.com/labstack/gommon/log"
	"github.com/leebrouse/Gorder/common/config"
	"github.com/leebrouse/Gorder/common/logging"
	"github.com/leebrouse/Gorder/common/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Viper init
func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

}
func main() {
	serverType := viper.GetString("payment.server-to-run")
	paymentHandler := NewPaymentHandler()
	switch serverType {
	case "http":
		server.RunHTTPServer(viper.GetString("payment.service-name"), paymentHandler.RegisterRoutes)
	case "grpc":
		logrus.Panic("unsupported type")
	default:
		logrus.Panic("Unexpected server type")
	}
}
