package app

import "github.com/leebrouse/Gorder/stock/app/query"

// Application
type Application struct {
	//	commend
	Commend Commend
	//	queries
	Queries Queries
}

// Commend
type Commend struct{}

// Queries
type Queries struct {
	CheckIfItemsInStock query.CheckIfItemsInStockHandler
	GetItems            query.GetItemsHandler
}
