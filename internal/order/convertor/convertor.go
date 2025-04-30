package convertor

import (
	client "github.com/leebrouse/Gorder/common/client/order"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	domain "github.com/leebrouse/Gorder/order/domain/order"
	"github.com/leebrouse/Gorder/order/entity"
)

// ==============================
// 类型定义（Convertor）
// ==============================
type OrderConvertor struct{}            // 订单转换器
type ItemConvertor struct{}             // 单个商品转换器（带 Name/PriceID）
type ItemWithQuantityConvertor struct{} // 简化商品信息转换器（仅 ID 和数量）

// ==============================
// ItemWithQuantityConvertor 转换逻辑
// ==============================

// 实体数组 → proto 数组
func (c *ItemWithQuantityConvertor) EntitiesToProtos(items []*entity.ItemWithQuantity) (res []*orderpb.ItemWithQuantity) {
	for _, i := range items {
		res = append(res, c.EntityToProto(i))
	}
	return
}

// 单个实体 → proto
func (c *ItemWithQuantityConvertor) EntityToProto(i *entity.ItemWithQuantity) *orderpb.ItemWithQuantity {
	return &orderpb.ItemWithQuantity{
		ID:       i.ID,
		Quantity: i.Quantity,
	}
}

// proto 数组 → 实体数组
func (c *ItemWithQuantityConvertor) ProtosToEntities(items []*orderpb.ItemWithQuantity) (res []*entity.ItemWithQuantity) {
	for _, i := range items {
		res = append(res, c.ProtoToEntity(i))
	}
	return
}

// 单个 proto → 实体
func (c *ItemWithQuantityConvertor) ProtoToEntity(i *orderpb.ItemWithQuantity) *entity.ItemWithQuantity {
	return &entity.ItemWithQuantity{
		ID:       i.ID,
		Quantity: i.Quantity,
	}
}

// client 数组 → 实体数组
func (c *ItemWithQuantityConvertor) ClientsToEntities(items []client.ItemWithQuantity) (res []*entity.ItemWithQuantity) {
	for _, i := range items {
		res = append(res, c.ClientToEntity(i))
	}
	return
}

// 单个 client → 实体
func (c *ItemWithQuantityConvertor) ClientToEntity(i client.ItemWithQuantity) *entity.ItemWithQuantity {
	return &entity.ItemWithQuantity{
		ID:       i.Id,
		Quantity: i.Quantity,
	}
}

// ==============================
// OrderConvertor 转换逻辑
// ==============================

// 实体 → proto
func (c *OrderConvertor) EntityToProto(o *domain.Order) *orderpb.Order {
	c.check(o)
	return &orderpb.Order{
		ID:          o.ID,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		Items:       NewItemConvertor().EntitiesToProtos(o.Items), // 这里默认使用 Item 类型
		PaymentLink: o.PaymentLink,
	}
}

// proto → 实体
func (c *OrderConvertor) ProtoToEntity(o *orderpb.Order) *domain.Order {
	c.check(o)
	return &domain.Order{
		ID:          o.ID,
		CustomerID:  o.CustomerID,
		Status:      o.Status,
		PaymentLink: o.PaymentLink,
		Items:       NewItemConvertor().ProtosToEntities(o.Items),
	}
}

// client → 实体
func (c *OrderConvertor) ClientToEntity(o *client.Order) *domain.Order {
	c.check(o)
	return &domain.Order{
		ID:          o.Id,
		CustomerID:  o.CustomerId,
		Status:      o.Status,
		PaymentLink: o.PaymentLink,
		Items:       NewItemConvertor().ClientsToEntities(o.Items),
	}
}

// 实体 → client
func (c *OrderConvertor) EntityToClient(o *domain.Order) *client.Order {
	c.check(o)
	return &client.Order{
		Id:          o.ID,
		CustomerId:  o.CustomerID,
		Status:      o.Status,
		PaymentLink: o.PaymentLink,
		Items:       NewItemConvertor().EntitiesToClients(o.Items),
	}
}

// 空值检查
func (c *OrderConvertor) check(o interface{}) {
	if o == nil {
		panic("connot convert nil order") // 建议拼写修正为 "cannot"
	}
}

// ==============================
// ItemConvertor 转换逻辑
// ==============================

// 实体数组 → proto 数组
func (c *ItemConvertor) EntitiesToProtos(items []*entity.Item) (res []*orderpb.Item) {
	for _, i := range items {
		res = append(res, c.EntityToProto(i))
	}
	return
}

// proto 数组 → 实体数组
func (c *ItemConvertor) ProtosToEntities(items []*orderpb.Item) (res []*entity.Item) {
	for _, i := range items {
		res = append(res, c.ProtoToEntity(i))
	}
	return
}

// client 数组 → 实体数组
func (c *ItemConvertor) ClientsToEntities(items []client.Item) (res []*entity.Item) {
	for _, i := range items {
		res = append(res, c.ClientToEntity(i))
	}
	return
}

// 实体数组 → client 数组
func (c *ItemConvertor) EntitiesToClients(items []*entity.Item) (res []client.Item) {
	for _, i := range items {
		res = append(res, c.EntityToClient(i))
	}
	return
}

// 单个实体 → proto
func (c *ItemConvertor) EntityToProto(i *entity.Item) *orderpb.Item {
	return &orderpb.Item{
		ID:       i.ID,
		Name:     i.Name,
		Quantity: i.Quantity,
		PriceID:  i.PriceID,
	}
}

// proto → 实体
func (c *ItemConvertor) ProtoToEntity(i *orderpb.Item) *entity.Item {
	return &entity.Item{
		ID:       i.ID,
		Name:     i.Name,
		Quantity: i.Quantity,
		PriceID:  i.PriceID,
	}
}

// client → 实体
func (c *ItemConvertor) ClientToEntity(i client.Item) *entity.Item {
	return &entity.Item{
		ID:       i.Id,
		Name:     i.Name,
		Quantity: i.Quantity,
		PriceID:  i.PriceId,
	}
}

// 实体 → client
func (c *ItemConvertor) EntityToClient(i *entity.Item) client.Item {
	return client.Item{
		Id:       i.ID,
		Name:     i.Name,
		Quantity: i.Quantity,
		PriceId:  i.PriceID,
	}
}
