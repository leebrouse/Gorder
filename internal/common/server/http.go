package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	//using the serviceName and get its elements under the service
	addr := viper.Sub(serviceName).GetString("http-addr")
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)

	//Router 组
	apiRouter.Group("/api")

	//Test router
	//apiRouter.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, "pong")
	//})

	//run gin server
	if err := apiRouter.Run(addr); err != nil {

	}
}
