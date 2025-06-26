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
	ErrBindBody         = errors.New("binding body error")
	ErrTransaction      = errors.New("transaction error")
	ErrGetBook          = errors.New("get book error")
	ErrNotFoundCategory = errors.New("not found category")
	ErrCreateBook       = errors.New("create book unsuccess")
	ErrGetFile          = errors.New("get file error")
	ErrStrConvert       = errors.New("convert string error")
	ErrGetListBook      = errors.New("get list book error")
)

type BookHandler struct {
	bookService services.IBookService
}

func NewBookHandler(bs services.IBookService) *BookHandler {
	return &BookHandler{
		bookService: bs,
	}
}

func (bh *BookHandler) CreateBook(c *gin.Context) {
	var req request.CreateBookRequest
	// Step 1: Binding body
	if err := c.ShouldBind(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrBindBody, err).Error())
		return
	}

	file, err := c.FormFile("thumbnail")

	req.Thumbnail = file

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetFile, err).Error())
		return
	}

	// Step 2: Create book
	book, err := bh.bookService.CreateBook(&req)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrCreateBook, err).Error())
		return
	}

	// Step 3: Create response

	res := &response.NewBookResponse{}

	copier.Copy(&res, book)

	utils.WriteResponse(c, http.StatusCreated, "Create book success", res, nil)

}

// GetBooks godoc
// @Summary Get list of books
// @Description get all books
// @Tags books
// @Accept  json
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {array} string
// @Router /books [get]
func (bh *BookHandler) GetListBook(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	categoryID := c.DefaultQuery("category_id", "0")
	authorID := c.DefaultQuery("author_id", "0")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrStrConvert, err).Error())
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrStrConvert, err).Error())
		return
	}

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrStrConvert, err).Error())
		return
	}

	authorIDInt, err := strconv.Atoi(authorID)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrStrConvert, err).Error())
		return
	}

	query := &request.QueryLimit{
		Limit: limitInt,
		Page:  pageInt,
	}

	books, err := bh.bookService.GetListBook(query, categoryIDInt, authorIDInt)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetBook, err).Error())
		return
	}

	booksData := []response.BookData{}
	copier.Copy(&booksData, books)

	res := &response.GetBookResponse{
		Limit: limitInt,
		Page:  pageInt,
		Total: len(books),
		Data:  booksData,
	}

	utils.WriteResponse(c, http.StatusOK, "Get All Book Success", res, nil)
}

// GetBookDetail godoc
// @Summary Get book detail
// @Description get book detail by ID
// @Tags books
// @Accept  json
// @Security ApiKeyAuth
// @Produce  json
// @Param bookID path int true "Book ID"
// @Success 200 {object} response.BookData
// @Router /books/{bookID} [get]
// GetBookDetail retrieves the details of a book by its ID.
func (bh *BookHandler) GetBookDetail(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("bookID"))
	book, err := bh.bookService.GetBookByID(bookID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrGetBook, err).Error())
	}

	// Step 3: Create Response
	utils.WriteResponse(c, http.StatusOK, "Get book Success", book, nil)

}
