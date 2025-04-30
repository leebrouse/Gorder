package dto

// 返回order响应结构体
type CreateOrderResponse struct {
	CustomerID  string `json:"customer_id"`
	OrderID     string `json:"order_id"`
	RedirectURL string `json:"redirect_url"` // 下单成功后的跳转地址
}

// new order response
func NewCreateOrderResponse(customerID string, orderID string, redirectURL string) CreateOrderResponse {
	return CreateOrderResponse{
		CustomerID:  customerID,
		OrderID:     orderID,
		RedirectURL: redirectURL,
	}
}
