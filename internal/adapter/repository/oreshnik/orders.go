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
	ordersTable      = "orders"
	ordersItemsTable = "orders_items"
)

type ordersRepository struct {
	db *postgres.Postgres
}

func NewOrdersRepository(
	db *postgres.Postgres,
) repository.OrdersRepository {
	return &ordersRepository{
		db: db,
	}
}

func (r *ordersRepository) CreateOrder(ctx context.Context, order entities.CreateOrder) (int, error) {
	tx, err := r.db.SqlxDB().Beginx()
	if err != nil {
		return 0, fmt.Errorf("transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	qb := r.db.Builder.Insert(ordersTable).
		Columns(
			"postomat_id",
			"payment_id",
		).Values(
		order.PostomatID,
		order.PaymentID,
	).Suffix(
		"RETURNING id",
	)

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	var orderID int
	if err = tx.QueryRowxContext(ctx, query, args...).Scan(&orderID); err != nil {
		return 0, fmt.Errorf("scan: %w", err)
	}

	for _, itemCount := range order.Items {
		itemQb := r.db.Builder.Insert(ordersItemsTable).
			Columns(
				"order_id",
				"item_id",
				"count",
			).Values(
			orderID,
			itemCount.ItemID,
			itemCount.Count,
		)

		itemQuery, itemArgs, err := itemQb.ToSql()
		if err != nil {
			return 0, fmt.Errorf("to sql: %w", err)
		}

		if _, err = tx.ExecContext(ctx, itemQuery, itemArgs...); err != nil {
			return 0, fmt.Errorf("exec: %w", err)
		}
	}

	return orderID, nil
}

func (r *ordersRepository) UpdateOrderStatus(ctx context.Context, order entities.UpdateOrder) error {
	qb := r.db.Builder.Update(ordersTable).
		Set("status", order.Status).
		Where(squirrel.Eq{
			"id": order.ID,
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

func (r *ordersRepository) GetOrders(ctx context.Context) ([]entities.Order, error) {
	qb := r.db.Builder.Select(
		"o.id",
		"o.postomat_id",
		"o.payment_id",
		"o.status",
		"oi.item_id",
		"oi.count",
	).
		From("orders AS o").
		Join("orders_items AS oi ON o.id = oi.order_id")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	rows, err := r.db.SqlxDB().QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	ordersMap := make(map[int]*entities.Order)
	for rows.Next() {
		var (
			order  entities.Order
			itemID string
			count  int
		)
		if err = rows.Scan(&order.ID, &order.PostomatID, &order.PaymentID, &order.Status, &itemID, &count); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		if existingOrder, exists := ordersMap[order.ID]; exists {
			existingOrder.Items = append(existingOrder.Items, entities.ItemCount{
				ItemID: itemID,
				Count:  count,
			})
		} else {
			order.Items = append(order.Items, entities.ItemCount{
				ItemID: itemID,
				Count:  count,
			})
			ordersMap[order.ID] = &order
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	var orders []entities.Order
	for _, order := range ordersMap {
		orders = append(orders, *order)
	}

	return orders, nil
}
