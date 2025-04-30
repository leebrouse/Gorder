package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common"
	client "github.com/leebrouse/Gorder/common/client/order"
	"github.com/leebrouse/Gorder/order/app"
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/dto"
	"github.com/leebrouse/Gorder/order/app/query"
	"github.com/leebrouse/Gorder/order/convertor"
)

// HTTPServer 结构体继承了 BaseResponse，用于处理 HTTP 请求，并持有业务逻辑层 app.Application 的引用。
type HTTPServer struct {
	common.BaseResponse
	app app.Application
}

// PostCustomerCustomerIDOrders 处理客户创建订单的 POST 请求
func (H HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {
	var (
		req  client.CreateOrderRequest // 请求体对象，包含客户提交的订单信息（如商品ID、数量等）
		resp dto.CreateOrderResponse   // order response
		err  error                     // 用于错误捕获
	)

	// 在函数结束时统一处理响应输出，无论成功或失败
	defer func() {
		H.Response(c, err, &resp)
	}()

	// 尝试将请求体绑定到 req，如果绑定失败直接返回错误
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	// 调用业务逻辑层的 CreateOrder handler 创建订单
	r, err := H.app.Commend.CreateOrder.Handle(c.Request.Context(), command.CreateOrder{
		CustomerID: req.CustomerId,
		Items:      convertor.NewItemWithQuantityConvertor().ClientsToEntities(req.Items),
	})
	if err != nil {
		return
	}

	// 构造响应体
	resp = dto.NewCreateOrderResponse(
		req.CustomerId,
		r.OrderID,
		fmt.Sprintf("http://localhost:8283/success?customerID=%s&orderID=%s", req.CustomerId, r.OrderID),
	)
}

// GetCustomerCustomerIDOrdersOrdersID 查询客户订单信息的 GET 请求
func (H HTTPServer) GetCustomerCustomerIdOrdersOrdersId(c *gin.Context, customerId string, ordersID string) {
	var (
		err  error       // 错误变量
		resp interface{} // 响应数据（泛型接口，可适配不同返回类型）
	)

	// 统一响应处理
	defer func() {
		H.Response(c, err, resp)
	}()

	// 调用业务逻辑层查询订单信息
	o, err := H.app.Queries.GetCustomOrder.Handle(c.Request.Context(), query.GetCustomerOrder{
		OrderID:    ordersID,
		CustomerID: customerId,
	})
	if err != nil {
		return
	}

	// 将实体对象转换为客户端响应对象
	resp = convertor.NewOrderConvertor().EntityToClient(o)
}
