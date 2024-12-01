package converters

import (
	"users/internal/models/dto"
	"users/internal/models/entities"
)

type CartsConverter struct{}

func NewCartsConverter() *CartsConverter {
	return &CartsConverter{}
}

func (c *CartsConverter) ToCartItemDTO(cartItem entities.CartItem) dto.CartItem {
	return dto.CartItem{
		ItemID: cartItem.ItemID,
		Count:  cartItem.Count,
	}
}

func (c *CartsConverter) ToCartItemDTOs(items []entities.CartItem) []dto.CartItem {
	res := make([]dto.CartItem, len(items))
	for i, item := range items {
		res[i] = c.ToCartItemDTO(item)
	}

	return res
}

func (c *CartsConverter) ToUpdateCartItemEntity(item dto.UpdateCartItem) entities.UpdateCartItem {
	return entities.UpdateCartItem{
		CartItem: entities.CartItem{
			ItemID: item.ItemID,
			Count:  item.Count,
		},
		UserID: item.UserID,
	}
}
