package entities

type UpdateOrder struct {
	ID     int
	Status string
}

type CreateOrder struct {
	Items      []ItemCount `db:"items"`
	PostomatID string      `db:"postomat_id"`
	PaymentID  string      `db:"payment_id"`
}

type Order struct {
	CreateOrder
	ID     int    `db:"id"`
	Status string `db:"status"`
}
