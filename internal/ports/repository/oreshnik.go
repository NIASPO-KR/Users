package repository

import (
	"context"

	"users/internal/models/entities"
)

type CartsRepository interface {
	GetCartByUserID(ctx context.Context, userID string) ([]entities.CartItem, error)
	UpdateCartItem(ctx context.Context, newCartItem entities.UpdateCartItem) error
}
