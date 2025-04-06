package app

import (
	"github.com/leebrouse/Gorder/order/app/command"
	"github.com/leebrouse/Gorder/order/app/query"
)

// Application
type Application struct {
	//	commend
	Commend Commend
	//	queries
	Queries Queries
}

// Commend
type Commend struct {
	CreateOrder command.CreateOrderHandler
	UpdateOrder command.UpdateOrderHandler
}

// Queries
type Queries struct {
	GetCustomOrder query.GetCustomerOrderHandler
}
