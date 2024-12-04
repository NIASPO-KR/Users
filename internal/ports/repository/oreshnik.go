package repository

import (
	"context"

	"users/internal/models/entities"
)

type CartsRepository interface {
	GetCartByUserID(ctx context.Context, userID string) ([]entities.ItemCount, error)
	UpdateCartItem(ctx context.Context, newCartItem entities.UpdateCartItem) error
}

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entities.CreateOrder) (int, error)
	UpdateOrderStatus(ctx context.Context, order entities.UpdateOrder) error
	GetOrders(ctx context.Context) ([]entities.Order, error)
}
