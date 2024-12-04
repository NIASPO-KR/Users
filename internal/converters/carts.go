package converters

import (
	"users/internal/models/dto"
	"users/internal/models/entities"
)

type CartsConverter struct{}

func NewCartsConverter() *CartsConverter {
	return &CartsConverter{}
}

func (c *CartsConverter) ToItemCountDTO(cartItem entities.ItemCount) dto.ItemCount {
	return dto.ItemCount{
		ItemID: cartItem.ItemID,
		Count:  cartItem.Count,
	}
}

func (c *CartsConverter) ToItemCountDTOs(items []entities.ItemCount) []dto.ItemCount {
	res := make([]dto.ItemCount, len(items))
	for i, item := range items {
		res[i] = c.ToItemCountDTO(item)
	}

	return res
}

func (c *CartsConverter) ToItemCountEntity(cartItem dto.ItemCount) entities.ItemCount {
	return entities.ItemCount{
		ItemID: cartItem.ItemID,
		Count:  cartItem.Count,
	}
}

func (c *CartsConverter) ToItemCountEntities(items []dto.ItemCount) []entities.ItemCount {
	res := make([]entities.ItemCount, len(items))
	for i, item := range items {
		res[i] = c.ToItemCountEntity(item)
	}

	return res
}

func (c *CartsConverter) ToUpdateCartItemEntity(item dto.UpdateCartItem) entities.UpdateCartItem {
	return entities.UpdateCartItem{
		ItemCount: entities.ItemCount{
			ItemID: item.ItemID,
			Count:  item.Count,
		},
		UserID: item.UserID,
	}
}
