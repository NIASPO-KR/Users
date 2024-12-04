package users

import (
	"context"
	"fmt"

	"users/internal/converters"
	"users/internal/models/dto"
	"users/internal/ports/repository"
	"users/internal/usecase"
)

type cartsUseCase struct {
	cartsRepository repository.CartsRepository
	cartsConverter  *converters.CartsConverter
}

func NewCartsUseCase(
	cartsRepository repository.CartsRepository,
) usecase.CartsUseCase {
	return &cartsUseCase{
		cartsRepository: cartsRepository,
		cartsConverter:  converters.NewCartsConverter(),
	}
}

func (cuc *cartsUseCase) GetCarts(ctx context.Context) ([]dto.ItemCount, error) {
	carts, err := cuc.cartsRepository.GetCartByUserID(ctx, usecase.MockUsername)
	if err != nil {
		return nil, fmt.Errorf("repo get cart by user id: %w", err)
	}

	return cuc.cartsConverter.ToItemCountDTOs(carts), nil
}

func (cuc *cartsUseCase) UpdateCartItem(ctx context.Context, newCartItem dto.UpdateCartItem) error {
	err := cuc.cartsRepository.UpdateCartItem(ctx, cuc.cartsConverter.ToUpdateCartItemEntity(newCartItem))
	if err != nil {
		return fmt.Errorf("repo update cart: %w", err)
	}

	return nil
}
