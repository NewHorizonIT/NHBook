package services

import (
	"errors"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/jinzhu/copier"
)

var (
	ErrBookExist           = errors.New("book isn't exist")
	ErrCreateBookUnSuccess = errors.New("create book unsuccess")
)

type IBookService interface {
	CreateBook(book *request.CreateBookRequest) (*response.BookData, error)
	GetListBook(query *request.QueryLimit, categoryID int, authorID int) ([]models.Book, error)
	CheckBookExist(bookID int) (bool, error)
	GetBookByID(bookID int) (*response.BookData, error)
}

type bookService struct {
	bookRepo        repositories.IBookRepository
	authorService   IAuthorService
	categoryService ICategoryService
}

// GetBookByID implements IBookService.
func (b *bookService) GetBookByID(bookID int) (*response.BookData, error) {
	book, err := b.bookRepo.GetBookByID(bookID)

	if err != nil {
		return nil, err
	}

	bookData := &response.BookData{}
	copier.Copy(&bookData, &book)
	return bookData, nil
}

// CheckBookExist implements IBookService.
func (b *bookService) CheckBookExist(bookID int) (bool, error) {
	isExist, err := b.bookRepo.IsExistBook(bookID)

	if err != nil {
		return false, err
	}

	if !isExist {
		return false, ErrBookExist
	}

	return isExist, nil
}

// GetListBook implements IBookService.
func (b *bookService) GetListBook(query *request.QueryLimit, categoryID int, authorID int) ([]models.Book, error) {
	var books []models.Book

	books, err := b.bookRepo.GetListBook(query, categoryID, authorID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// CreateBook implements IBookService.
func (b *bookService) CreateBook(book *request.CreateBookRequest) (*response.BookData, error) {
	// Fix Create Book

	// Start transaction
	tx := global.MySQL.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Check category is exists
	category, err := b.categoryService.CheckCategoryExitsByID(book.CategoryID, tx)

	if err != nil || category == nil {
		tx.Rollback()
		return nil, err
	}

	// Step 2: Check Author is invalid
	var authors []models.Author
	for _, authorName := range book.Authors {
		author, err := b.authorService.CheckAuthorExists(authorName, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		authors = append(authors, *author)
	}

	// Step 3: Paser Time publish
	pushlishAtParsed, err := time.Parse(time.DateOnly, book.PublishedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 4: Save thumbnail into cloudinary
	src, err := book.Thumbnail.Open()
	defer func() {
		src.Close()
	}()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	fileName := "book_" + utils.RandomString(12)

	uploadURL, err := utils.UploadImg(src, fileName, global.Cloudinary)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 5: Create book

	newBook := &models.Book{
		Title:       book.Title,
		ImageURL:    uploadURL,
		Price:       book.Price,
		Description: book.Description,
		Authors:     authors,
		Stock:       book.Stock,
		CategoryID:  book.CategoryID,
		Category:    *category,
		PublishedAt: pushlishAtParsed,
	}

	if err := b.bookRepo.CreateBook(newBook, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	// Step 6: Commit
	tx.Commit()

	data := &response.BookData{}
	copier.Copy(data, newBook)

	// End transaction

	return data, nil
}

func NewBookService(br repositories.IBookRepository, as IAuthorService, cs ICategoryService) IBookService {
	return &bookService{
		bookRepo:        br,
		authorService:   as,
		categoryService: cs,
	}
}
