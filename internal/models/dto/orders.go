package dto

import "errors"

type UpdateOrder struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type CreateOrder struct {
	Items      []ItemCount `json:"items"`
	PostomatID string      `json:"postomatID"`
	PaymentID  string      `json:"paymentID"`
}

func (d *CreateOrder) IsValid() error {
	if len(d.Items) == 0 {
		return errors.New("empty items")
	}

	return nil
}

type Order struct {
	CreateOrder
	ID     int    `json:"id"`
	Status string `json:"status"`
}
