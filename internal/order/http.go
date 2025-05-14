package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common"
	client "github.com/leebrouse/Gorder/common/client/order"
	"github.com/leebrouse/Gorder/common/consts"
	"github.com/leebrouse/Gorder/common/handler/errors"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/dto"
	"github.com/leebrouse/Gorder/order/app/query"
	"github.com/leebrouse/Gorder/order/convertor"
)

type HTTPServer struct {
	common.BaseResponse
	app app.Application
}

func (H HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerID string) {
	var (
		req  client.CreateOrderRequest
		resp dto.CreateOrderResponse
		err  error
	)
	defer func() {
		H.Response(c, err, &resp)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		err = errors.NewWithError(consts.ErrnoBindRequestError, err)
		return
	}
	if err = H.validate(req); err != nil {
		err = errors.NewWithError(consts.ErrnoRequestValidateError, err)
		return
	}
	r, err := H.app.Commands.CreateOrder.Handle(c.Request.Context(), command.CreateOrder{
		CustomerID: req.CustomerId,
		Items:      convertor.NewItemWithQuantityConvertor().ClientsToEntities(req.Items),
	})
	if err != nil {
		//err = errors.NewWithError()
		return
	}
	resp = dto.CreateOrderResponse{
		OrderID:     r.OrderID,
		CustomerID:  req.CustomerId,
		RedirectURL: fmt.Sprintf("http://localhost:8282/success?customerID=%s&orderID=%s", req.CustomerId, r.OrderID),
	}
}

func (H HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerID string, orderID string) {
	var (
		err  error
		resp interface{}
	)
	defer func() {
		H.Response(c, err, resp)
	}()

	o, err := H.app.Queries.GetCustomerOrder.Handle(c.Request.Context(), query.GetCustomerOrder{
		OrderID:    orderID,
		CustomerID: customerID,
	})
	if err != nil {
		return
	}

	resp = convertor.NewOrderConvertor().EntityToClient(o)
}

func (H HTTPServer) validate(req client.CreateOrderRequest) error {
	for _, v := range req.Items {
		if v.Quantity <= 0 {
			return fmt.Errorf("quantity must be positive, got %d from %s", v.Quantity, v.Id)
		}
	}
	return nil
}
