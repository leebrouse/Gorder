package order

import (
	"fmt"

	"golang.org/x/net/context"
)

// order service entity interface for grpc service
type Repository interface {
	Create(context.Context, *Order) (*Order, error)
	Get(ctx context.Context, id, customerID string) (*Order, error)
	Update(
		ctx context.Context,
		o *Order,
		updateFn func(context.Context, *Order) (*Order, error),
	) error
}

type NotFoundError struct {
	OrderID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("order '%s' not found", e.OrderID)
}
