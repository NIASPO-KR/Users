package oreshnik

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"users/internal/errs"
	"users/internal/infrastructure/database/postgres"
	"users/internal/models/entities"
	"users/internal/ports/repository"
)

const (
	cartsTable = "carts"
)

type cartsRepository struct {
	db *postgres.Postgres
}

func NewCartsRepository(
	db *postgres.Postgres,
) repository.CartsRepository {
	return &cartsRepository{
		db: db,
	}
}

func (r *cartsRepository) CreateCartItem(ctx context.Context, cartItem entities.UpdateCartItem) error {
	qb := r.db.Builder.
		Insert(cartsTable).
		Columns(
			"item_id",
			"count",
		).
		Values(
			cartItem.ItemID,
			cartItem.Count,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := r.db.SqlxDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %w", err)
	}

	return nil
}

func (r *cartsRepository) GetCartByUserID(ctx context.Context, userID string) ([]entities.ItemCount, error) {
	qb := r.db.Builder.Select(
		"item_id",
		"count",
	).
		From(cartsTable).
		Where(squirrel.Eq{"user_id": userID})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql %w", err)
	}

	var cartItems []entities.ItemCount
	if err := r.db.SqlxDB().SelectContext(ctx, &cartItems, query, args...); err != nil {
		return nil, fmt.Errorf("select %w", err)
	}

	return cartItems, nil
}

func (r *cartsRepository) UpdateCartItem(ctx context.Context, newCartItem entities.UpdateCartItem) error {
	qb := r.db.Builder.Update(cartsTable).
		Set("count", newCartItem.Count).
		Where(squirrel.Eq{
			"item_id": newCartItem.ItemID,
			"user_id": newCartItem.UserID,
		})

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := r.db.SqlxDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *cartsRepository) DeleteCartItem(ctx context.Context, cartItemID string) error {
	qb := r.db.Builder.
		Delete(cartsTable).
		Where(squirrel.Eq{"item_id": cartItemID})

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := r.db.SqlxDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
