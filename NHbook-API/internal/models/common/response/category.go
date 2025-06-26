package response

import "time"

type CategoryResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListCategoryResponse struct {
	Total      int                `json:"total"`
	Categories []CategoryResponse `json:"categories"`
}
