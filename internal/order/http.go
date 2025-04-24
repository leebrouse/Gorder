package main

import (
	"fmt"
	"github.com/leebrouse/Gorder/common/tracing"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
)

type HTTPServer struct {
	app app.Application
}

// (POST /customer/{customerID}/orders)
func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	// 手动创建一个子 Span
	ctx, span := tracing.Start(c, "PostCustomerCustomerIDOrders")
	defer span.End()
	//create order 请求体
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	r, err := H.app.Commend.CreateOrder.Handle(ctx, command.CreateOrder{
		CustomerID: req.CustomerID,
		Items:      req.Item,
	})
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err})
		return
	}
	//create traceID
	//traceID := tracing.TraceID(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"trace":        tracing.TraceID(ctx),
		"customer_id":  req.CustomerID,
		"order_id":     r.OrderID,
		"redirect_url": fmt.Sprintf("http://localhost:8283/success?customerID=%s&orderID=%s", req.CustomerID, r.OrderID),
	})
}

// (GET /customer/{customerID}/orders/{ordersID})
func (H HTTPServer) GetCustomerCustomerIDOrdersOrdersID(c *gin.Context, customerID string, ordersID string) {
	// 手动创建一个子 Span
	ctx, span := tracing.Start(c, "GetCustomerCustomerIDOrdersOrdersID")
	defer span.End()

	order, err := H.app.Queries.GetCustomOrder.Handle(ctx, query.GetCustomerOrder{
		OrderID:    ordersID,
		CustomerID: customerID,
	})
	if err != nil {
		//fail to get order
		c.JSON(http.StatusNoContent, gin.H{"error": err})
		return
	}

	//create traceID
	//traceID := tracing.TraceID(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"traceID": tracing.TraceID(ctx),
		"data": gin.H{
			"Order": order,
		},
	})
}
