package response

import (
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
)

type CreateBookResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"imageUrl"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	PublishedAt string    `json:"publishAt"`
	Authors     []string  `json:"authors"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	CategoryID  uint      `json:"categoryID"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GetBookResponse struct {
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
	Total int           `json:"total"`
	Data  []models.Book `json:"data"`
}

type GetListBookByCategoryResponse struct {
	Quantity int           `json:"quantity" `
	Limit    int           `json:"limit"`
	Page     int           `json:"page"`
	Data     []models.Book `json:"data"`
}
