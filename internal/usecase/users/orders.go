package users

import (
	"context"
	"fmt"

	"users/internal/converters"
	"users/internal/models/dto"
	"users/internal/ports/repository"
	"users/internal/usecase"
)

type ordersUseCase struct {
	ordersRepository repository.OrdersRepository
	ordersConverter  *converters.OrdersConverter
}

func NewOrdersUseCase(
	ordersRepository repository.OrdersRepository,
) usecase.OrdersUseCase {
	return &ordersUseCase{
		ordersRepository: ordersRepository,
		ordersConverter:  converters.NewOrdersConverter(),
	}
}

func (ouc *ordersUseCase) CreateOrder(ctx context.Context, order dto.CreateOrder) (int, error) {
	id, err := ouc.ordersRepository.CreateOrder(ctx, ouc.ordersConverter.ToOrderCreateEntity(order))
	if err != nil {
		return 0, fmt.Errorf("repo create order: %w", err)
	}

	return id, nil
}

func (ouc *ordersUseCase) UpdateOrderStatus(ctx context.Context, order dto.UpdateOrder) error {
	if err := ouc.ordersRepository.UpdateOrderStatus(ctx, ouc.ordersConverter.ToOrderUpdateEntity(order)); err != nil {
		return fmt.Errorf("repo update order: %w", err)
	}

	return nil
}

func (ouc *ordersUseCase) GetOrders(ctx context.Context) ([]dto.Order, error) {
	orders, err := ouc.ordersRepository.GetOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get orders: %w", err)
	}

	return ouc.ordersConverter.ToOrderDTOs(orders), nil
}
