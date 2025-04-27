package convertor

import (
	client "github.com/leebrouse/Gorder/common/client/order"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/order/entity"
)

type OrderConvertor struct{}
type ItemConvertor struct{}
type ItemWithQuantityConvertor struct{}

// Main Convertor function
func (c *OrderConvertor) EntityToProto(o entity.Order) *orderpb.Order {
	return &orderpb.Order{
		ID:          o.ID,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		Items:       NewItemConvertor().EntitiesToProto(o.Items),
		PaymentLink: o.PaymentLink,
	}
}

func (c *OrderConvertor) ProtoToEntity(o *orderpb.Order) *entity.Order {
	return &entity.Order{
		ID:          o.ID,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		Items:       NewItemConvertor().ProtosToEntities(o.Items),
		PaymentLink: o.PaymentLink,
	}
}

func (c *OrderConvertor) ClientToEntity(o *client.Order) *entity.Order {
	return &entity.Order{
		ID:          o.Id,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		Items:       NewItemConvertor().ClientsToEntities(o.Items),
		PaymentLink: o.PaymentLink,
	}
}

func (c *OrderConvertor) EntityToClient(o *entity.Order) *client.Order {
	return &client.Order{
		Id:          o.ID,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		Items:       NewItemConvertor().EntitiesToClients(o.Items),
		PaymentLink: o.PaymentLink,
	}
}

// Convertor Helper
func (c *ItemConvertor) EntitiesToProto(items []*entity.Item) (res []*orderpb.Item) {
	panic("")
}

func (c *ItemConvertor) ProtosToEntities(items []*orderpb.Item) (res []*entity.Item) {
	panic("")
}

func (c *ItemConvertor) ClientsToEntities(items []client.Item) (res []*entity.Item) {
	panic("")
}

func (c *ItemConvertor) EntitiesToClients(items []*entity.Item) []client.Item {
	panic("")
}
