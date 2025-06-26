package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var (
	ErrBindCategory   = errors.New("bind body category error")
	ErrCreateCtegory  = errors.New("create category error")
	ErrGetCategory    = errors.New("get category error")
	ErrParam          = errors.New("param invalid")
	ErrUpdateCategory = errors.New("update category error")
)

const (
	categoyPrivate = 0
	categoryPublic = 1
	categoryAll    = 3
)

type CategoryHandler struct {
	categoryService services.ICategoryService
}

// @Summary Create a new category
// @Description Create a new category with name and status
// @Tags category
// @Accept json
// @Produce json
// @Param category body request.CategoryRequest true "Category details"
// @Success 201 {object} utils.ResponseSuccess{data=response.CategoryResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /categories [post]
// @Security ApiKeyAuth
// CreateCategory creates a new category
func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	// Step 1: Binding data
	var req request.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrBindCategory, err).Error())
		return
	}

	// Step 2: Call service create CreateCategory
	category, err := ch.categoryService.CreateCategory(&req)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrCreateCtegory, err).Error())
		return
	}

	res := &response.CategoryResponse{}
	copier.Copy(res, category)

	utils.WriteResponse(c, http.StatusCreated, "Create category success", res, nil)

}

// @Summary Get category detail
// @Description Get category detail by categoryID
// @Tags category
// @Accept json
// @Produce json
// @Param categoryID path int true "Category ID"
// @Success 200 {object} utils.ResponseSuccess{data=response.CategoryResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /categories/{categoryID} [get]
// @Security ApiKeyAuth
// GetcategoryDetail retrieves the details of a category by its ID
func (ch *CategoryHandler) GetcategoryDetail(c *gin.Context) {
	// Step1: Get param CategoryID
	id, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Step 2: Call Service GetCategoryByID
	category, err := ch.categoryService.GetCategoryByID(id)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetCategory, err).Error())
		return
	}

	// Step 3: Craete response
	res := &response.CategoryResponse{}
	copier.Copy(res, category)

	utils.WriteResponse(c, http.StatusOK, "Get Category success", res, nil)
}

// @Summary Get list category
// @Description Get list category by status
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} utils.ResponseSuccess{data=[]response.CategoryResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /categories [get]
// @Security ApiKeyAuth
// GetListCategory retrieves all categories
func (ch *CategoryHandler) GetListCategory(c *gin.Context) {
	// Step 1: Call service getAllCategory
	categories, err := ch.categoryService.GetAllCategory(categoryAll)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetCategory, err).Error())
		return
	}

	// Step 2: Create Response
	res := []response.CategoryResponse{}
	copier.Copy(&res, &categories)

	utils.WriteResponse(c, http.StatusOK, "Get list category Success", res, nil)
}

// @Summary Get list category by status
// @Description Get list category by status
// @Tags category
// @Accept json
// @Produce json
// @Param status path int true "Category status (0: private, 1: public, 3: all)"
// @Success 200 {object} utils.ResponseSuccess{data=[]response.CategoryResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /categories/status/{status} [get]
// @Security ApiKeyAuth
// GetListCategoryByStatus retrieves categories by their status
func (ch *CategoryHandler) GetListCategoryByStatus(c *gin.Context) {
	// Step 1: Call service getAllCategory
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, utils.FormatError(ErrParam, err).Error())
		return
	}
	categories, err := ch.categoryService.GetAllCategory(status)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetCategory, err).Error())
		return
	}

	// Step 2: Create Response
	res := []response.CategoryResponse{}
	copier.Copy(&res, &categories)

	utils.WriteResponse(c, http.StatusOK, "Get list category Success", res, nil)
}

// @Summary Update a category
// @Description Update a category with name and status
// @Tags category
// @Accept json
// @Produce json
// @Param category body request.CategoryUpdate true "Category update details"
// @Success 200 {object} utils.ResponseSuccess{data=response.CategoryResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /categories [put]
// @Security ApiKeyAuth
// UpdateCategory updates an existing category
func (ch *CategoryHandler) UpdateCategory(c *gin.Context) {
	var req request.CategoryUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrBindCategory, err).Error())
		return
	}

	category, err := ch.categoryService.UpdateCategory(&req)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrUpdateCategory, err).Error())
		return
	}

	res := &response.CategoryResponse{}
	copier.Copy(res, category)

	utils.WriteResponse(c, http.StatusOK, "Update Category Success", res, nil)
}

func NewCategoryHandler(cs services.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: cs,
	}
}
