package entities

type ItemCount struct {
	ItemID string `db:"item_id"`
	Count  int    `db:"count"`
}
