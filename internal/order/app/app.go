package app

import (
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateOrder command.CreateOrderHandler
	UpdateOrder command.UpdateOrderHandler
}

type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}
