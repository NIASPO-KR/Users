package entities

type UpdateCartItem struct {
	ItemCount
	UserID string `db:"user_id"`
}
