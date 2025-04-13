package adapters

import (
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strconv"
	"sync"
	"time"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

// New a Order Repository for CRUD Operation
func NewMemoryOrderRepository() *MemoryOrderRepository {
	//Test data
	s := make([]*domain.Order, 0)
	s = append(s, &domain.Order{
		ID:          "fake-ID",
		CustomerID:  "fake-customer-id",
		Status:      "fake-status",
		PaymentLink: "fake-payment-link",
		Items:       nil,
	})

	return &MemoryOrderRepository{
		lock: &sync.RWMutex{},
		//store: make([]*domain.Order, 0),
		store: s,
	}
}

// Impelment Repository interface
func (m *MemoryOrderRepository) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	//TODO implement me
	m.lock.Lock()
	defer m.lock.Unlock()
	//create a new order
	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		Items:       order.Items,
		PaymentLink: order.PaymentLink,
	}
	m.store = append(m.store, newOrder)
	// 记录 Debug 级别的日志，方便调试 MemoryOrderRepository 的存储过程
	logrus.WithFields(logrus.Fields{
		"input_order":        order,   // 记录当前传入的订单对象
		"store_after_create": m.store, // 记录订单存储后的状态
	}).Info("memory_order_repo_create") // 日志消息，标识日志来源
	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(_ context.Context, id, customerID string) (*domain.Order, error) {
	//TODO implement me
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, o := range m.store {
		if o.ID == id && o.CustomerID == customerID {
			// 记录 Debug 级别日志，标记找到匹配的订单
			logrus.Debugf("memory_order_repo_repo_get||found||id=%s||customerID=%s||res=%+v", id, customerID, *o)
			return o, nil
		}
	}
	return nil, domain.NotFoundError{OrderID: id}
}

func (m *MemoryOrderRepository) Update(ctx context.Context, order *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	//TODO implement me
	//add lock and unlock
	m.lock.Lock()
	defer m.lock.Unlock()

	//Using found as a flag to ensure having the required order in the Repository
	found := false
	//TO Find the required Order
	for i, o := range m.store {
		if o.ID == order.ID && o.CustomerID == order.CustomerID {
			found = true //find order
			//	update the order
			updateOrder, err := updateFn(ctx, o)
			if err != nil {
				return err
			}
			m.store[i] = updateOrder
		}
	}

	if !found {
		return domain.NotFoundError{OrderID: order.ID}
	}

	return nil
}
