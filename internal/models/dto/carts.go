package dto

type CartItem struct {
	ItemID string `json:"itemID"`
	Count  int    `json:"count"`
}

type UpdateCartItem struct {
	CartItem
	UserID string `json:"userID"`
}
