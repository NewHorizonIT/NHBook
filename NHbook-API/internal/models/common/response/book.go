package response

import (
	"time"
)

type NewBookResponse struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	ImageURL    string       `json:"imageUrl"`
	Price       int          `json:"price"`
	Description string       `json:"description"`
	PublishedAt time.Time    `json:"publishedAt"`
	Authors     AuthorData   `json:"authors"`
	Stock       int          `json:"stock"`
	Category    CategoryData `json:"category"`
	CreatedAt   time.Time    `json:"createdAt"`
}

type GetBookResponse struct {
	Limit int        `json:"limit"`
	Page  int        `json:"page"`
	Total int        `json:"total"`
	Data  []BookData `json:"data"`
}
type BookData struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	ImageURL    string       `json:"thumbnail"`
	Price       int          `json:"price"`
	Description string       `json:"description"`
	Stock       int          `json:"stock"`
	Category    CategoryData `json:"category"`
	Authors     []AuthorData `json:"authors"`
	PublishedAt time.Time    `json:"publishedAt"`
}

type CategoryData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthorData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
