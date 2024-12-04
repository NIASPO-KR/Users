package converters

import (
	"users/internal/models/dto"
	"users/internal/models/entities"
)

type OrdersConverter struct {
	cartsConverter *CartsConverter
}

func NewOrdersConverter() *OrdersConverter {
	return &OrdersConverter{
		cartsConverter: NewCartsConverter(),
	}
}

func (c *OrdersConverter) ToOrderDTO(order entities.Order) dto.Order {
	return dto.Order{
		CreateOrder: dto.CreateOrder{
			Items:      c.cartsConverter.ToItemCountDTOs(order.Items),
			PostomatID: order.PostomatID,
			PaymentID:  order.PaymentID,
		},
		ID:     order.ID,
		Status: order.Status,
	}
}

func (c *OrdersConverter) ToOrderDTOs(orders []entities.Order) []dto.Order {
	res := make([]dto.Order, len(orders))
	for i, order := range orders {
		res[i] = c.ToOrderDTO(order)
	}

	return res
}

func (c *OrdersConverter) ToOrderUpdateEntity(order dto.UpdateOrder) entities.UpdateOrder {
	return entities.UpdateOrder{
		ID:     order.ID,
		Status: order.Status,
	}
}

func (c *OrdersConverter) ToOrderCreateEntity(order dto.CreateOrder) entities.CreateOrder {
	return entities.CreateOrder{
		Items:      c.cartsConverter.ToItemCountEntities(order.Items),
		PostomatID: order.PostomatID,
		PaymentID:  order.PaymentID,
	}
}
