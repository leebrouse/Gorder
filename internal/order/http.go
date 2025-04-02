package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/order/app"
)

type HTTPServer struct {
	app app.Application
}

// (POST /customer/{customerID}/orders)
func (s HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {

}

// (GET /customer/{customerID}/orders/{ordersID})
func (s HTTPServer) GetCustomerCustomerIDOrdersOrdersID(c *gin.Context, customerID string, ordersID string) {

}
