package adapters

import (
	"context"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/stock/domain/stock"
	"sync"
)

type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
}

// New a Stock Repository for CRUD Operation
func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: stub,
	}
}

func (m MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var (
		res     []*orderpb.Item
		missing []string
	)

	for _, id := range ids {
		if item, ok := m.store[id]; ok {
			res = append(res, item)
		} else {
			missing = append(missing, id)
		}
	}

	if len(missing) > 0 { // 如果有缺失的 ID，则返回错误
		return res, domain.NotFoundError{Missing: missing}
	}

	return res, nil // 所有 ID 都找到了，正常返回
}
