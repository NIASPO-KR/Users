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
	if newCartItem.Count == 0 {
		if err := cuc.cartsRepository.DeleteCartItem(ctx, newCartItem.ItemID); err != nil {
			return fmt.Errorf("repo delete cart item: %w", err)
		}
		return nil
	}

	cartItemCounts, err := cuc.cartsRepository.GetCartByUserID(ctx, usecase.MockUsername)
	if err != nil {
		return fmt.Errorf("repo get cart by user id: %w", err)
	}

	hasItem := false
	for _, cartItemCount := range cartItemCounts {
		if cartItemCount.ItemID == newCartItem.ItemID {
			hasItem = true
			break
		}
	}

	if !hasItem {
		if err = cuc.cartsRepository.CreateCartItem(ctx, cuc.cartsConverter.ToUpdateCartItemEntity(newCartItem)); err != nil {
			return fmt.Errorf("repo create cart item: %w", err)
		}
	} else {
		if err = cuc.cartsRepository.UpdateCartItem(ctx, cuc.cartsConverter.ToUpdateCartItemEntity(newCartItem)); err != nil {
			return fmt.Errorf("repo update cart: %w", err)
		}
	}

	return nil
}
