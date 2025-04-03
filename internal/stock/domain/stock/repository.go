package stock

import (
	"fmt"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"golang.org/x/net/context"
	"strings"
)

// Stock service Repository for gRpc
type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

// implement error interface
func (e NotFoundError) Error() string {
	return fmt.Sprintf("these items not found in stock: %s", strings.Join(e.Missing, ","))
}
