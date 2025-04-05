package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/query"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

// (POST /customer/{customerID}/orders)
func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {

}

// (GET /customer/{customerID}/orders/{ordersID})
func (H HTTPServer) GetCustomerCustomerIDOrdersOrdersID(c *gin.Context, customerID string, ordersID string) {
	order, err := H.app.Queries.GetCustomOrder.Handle(c, query.GetCustomerOrder{
		OrderID:    "fake-ID",
		CustomerID: "fake-customer-id",
	})
	if err != nil {
		//fail to get order
		c.JSON(http.StatusNoContent, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}
