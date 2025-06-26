package request

import (
	"mime/multipart"
)

type CreateBookRequest struct {
	Title       string                `json:"title" form:"title"`
	Authors     []string              `json:"authors" form:"authors"`
	Thumbnail   *multipart.FileHeader `json:"imageUrl" form:"thumbnail"`
	Price       int                   `json:"price" form:"price"`
	Description string                `json:"description" form:"description"`
	Stock       int                   `json:"stock" form:"stock"`
	CategoryID  int                   `json:"categoryId" form:"categoryID"`
	PublishedAt string                `json:"publishedAt" form:"publishedAt"`
}
