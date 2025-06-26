package handlers

import "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"

type AuthorHandler struct {
	authorService services.IAuthorService
}

func NewAuthorHandler(as services.IAuthorService) *AuthorHandler {
	return &AuthorHandler{
		authorService: as,
	}
}
