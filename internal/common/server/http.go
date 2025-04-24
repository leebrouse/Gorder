package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	//using the serviceName and get its elements under the service
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		panic("empty http address")
	}
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	setMiddlewares(apiRouter)
	//Router 组
	apiRouter.Group("/api")
	//run gin server
	if err := apiRouter.Run(addr); err != nil {
		panic("Http server failed to run")
	}
}

func setMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	// 接入 OpenTelemetry 的链路追踪功能
	r.Use(otelgin.Middleware("default_server"))
	//r.Use()
}
