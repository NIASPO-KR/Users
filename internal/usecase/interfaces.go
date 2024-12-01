package usecase

import (
	"context"

	"users/internal/models/dto"
)

type CartsUseCase interface {
	GetCarts(ctx context.Context) ([]dto.CartItem, error)
	UpdateCartItem(ctx context.Context, newCartItem dto.UpdateCartItem) error
}
