package handlers

import "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"

type CategoryHandler struct {
	categoryService services.ICategoryService
}

func NewCategoryHandler(cs services.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: cs,
	}
}
