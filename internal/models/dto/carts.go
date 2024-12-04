package dto

type UpdateCartItem struct {
	ItemCount
	UserID string `json:"userID"`
}
