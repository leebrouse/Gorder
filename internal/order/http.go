package main

import "github.com/gin-gonic/gin"

type HTTPServer struct {
}

// (POST /customer/{customerID}/orders)
func (s HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {

}

// (GET /customer/{customerID}/orders/{ordersID})
func (s HTTPServer) GetCustomerCustomerIDOrdersOrdersID(c *gin.Context, customerID string, ordersID string) {

}
