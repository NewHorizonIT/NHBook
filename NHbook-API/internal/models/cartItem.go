package models

type CartItem struct {
	ID        uint    `json:"productID"`
	Thumbnail string  `json:"thumbnail"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
