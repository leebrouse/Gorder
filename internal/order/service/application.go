package service

import (
	"context"
	"github.com/leebrouse/Gorder/order/app"
)

// Order application
func NewApplication(ctx context.Context) app.Application {
	//orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}
