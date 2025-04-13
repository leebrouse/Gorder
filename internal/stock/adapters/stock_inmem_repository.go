package adapters

import (
	"context"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/stock/domain/stock"
	"sync"
)

// MemoryStockRepository 是一个内存中的库存仓储，负责对库存数据进行 CRUD 操作。
type MemoryStockRepository struct {
	// 使用读写互斥锁保护并发访问 store
	lock *sync.RWMutex
	// store 用于存储库存商品，键为商品 ID，值为对应的 orderpb.Item 对象
	store map[string]*orderpb.Item
}

// stub 为预置的测试数据，模拟一个库存商品
var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
	"item1": {
		ID:       "item1",
		Name:     "stub item1",
		Quantity: 10000,
		PriceID:  "stub_item1_price_id",
	},
	"item2": {
		ID:       "item2",
		Name:     "stub item2",
		Quantity: 10000,
		PriceID:  "stub_item2_price_id",
	},
	"item3": {
		ID:       "item3",
		Name:     "stub item3",
		Quantity: 10000,
		PriceID:  "stub_item3_price_id",
	},
}

// NewMemoryStockRepository 创建一个新的 MemoryStockRepository 实例，并初始化库存数据
func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{}, // 初始化读写锁
		store: stub,            // 使用预置的测试数据作为库存初始数据
	}
}

// GetItems 根据传入的商品 ID 列表查询对应的库存商品
// 如果所有 ID 都能找到，则返回对应的商品列表；
// 如果有缺失的 ID，则返回找到的商品列表和一个包含缺失 ID 的 NotFoundError 错误
func (m MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	// 加锁，防止并发写入或读取冲突
	m.lock.Lock()
	defer m.lock.Unlock()

	var (
		// 存储查询到的商品列表
		res []*orderpb.Item
		// 存储未在库存中找到的商品 ID
		missing []string
	)

	// 遍历传入的所有商品 ID
	for _, id := range ids {
		if item, ok := m.store[id]; ok {
			// 如果库存中存在该商品，将其加入结果列表
			res = append(res, item)
		} else {
			// 如果库存中不存在该商品，记录下缺失的 ID
			missing = append(missing, id)
		}
	}

	// 如果存在缺失的商品，则返回结果列表和一个 NotFoundError 错误
	if len(missing) > 0 {
		return res, domain.NotFoundError{Missing: missing}
	}

	// 如果所有商品都能查询到，则返回结果列表和 nil 错误
	return res, nil
}
