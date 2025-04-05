package models

type CartItem struct {
	ID        uint   `json:"productID"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Total     int    `json:"total"`
}
