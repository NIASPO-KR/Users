package usecase

import (
	"context"

	"users/internal/models/dto"
)

type CartsUseCase interface {
	GetCarts(ctx context.Context) ([]dto.ItemCount, error)
	UpdateCartItem(ctx context.Context, newCartItem dto.UpdateCartItem) error
}

type OrdersUseCase interface {
	CreateOrder(ctx context.Context, order dto.CreateOrder) (int, error)
	UpdateOrderStatus(ctx context.Context, order dto.UpdateOrder) error
	GetOrders(ctx context.Context) ([]dto.Order, error)
}
