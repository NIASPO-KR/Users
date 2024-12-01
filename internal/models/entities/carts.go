package entities

type CartItem struct {
	ItemID string `db:"item_id"`
	Count  int    `db:"count"`
}

type UpdateCartItem struct {
	CartItem
	UserID string `db:"user_id"`
}
