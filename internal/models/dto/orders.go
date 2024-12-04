package dto

type UpdateOrder struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type CreateOrder struct {
	Items      []ItemCount `json:"items"`
	PostomatID string      `json:"postomatID"`
	PaymentID  string      `json:"paymentID"`
}

func (d CreateOrder) IsValid() bool {
	return len(d.Items) != 0
}

type Order struct {
	CreateOrder
	ID     int    `json:"id"`
	Status string `json:"status"`
}
