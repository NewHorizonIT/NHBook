package response

import "time"

type CreateBookResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"imageUrl"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	CategoryID  uint      `json:"categoryID"`
	CreatedAt   time.Time `json:"createdAt"`
}
