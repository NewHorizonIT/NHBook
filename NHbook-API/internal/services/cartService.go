package services

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
)

type ICartService interface {
	AddItemCartByID(userID string, item *models.CartItem) error
	GetCartByID(userID string) ([]models.CartItem, error)
	DeleteCartByID(userID string) ([]models.CartItem, error)
	RemoveItemCart(userID string, productID int) ([]models.CartItem, error)
}

type cartService struct {
	cartRepo repositories.ICartRepository
}

// RemoveItemCart implements ICartService.
func (c *cartService) RemoveItemCart(userID string, productID int) ([]models.CartItem, error) {
	var cart []models.CartItem

	if err := c.cartRepo.RemoveItemInCart(userID, productID); err != nil {
		return nil, err
	}

	cart, err := c.GetCartByID(userID)

	if err != nil {
		return nil, err
	}

	return cart, nil

}

// DeleteCartByID implements ICartService.
func (c *cartService) DeleteCartByID(userID string) ([]models.CartItem, error) {
	// Delete all item in cart
	err := c.cartRepo.DeleteCart(userID)

	if err != nil {
		return nil, err
	}

	// Get new cart
	cart, err := c.GetCartByID(userID)

	if err != nil {
		return nil, err
	}
	return cart, nil
}

// GetCartByID implements ICartService.
func (c *cartService) GetCartByID(userID string) ([]models.CartItem, error) {
	cart, err := c.cartRepo.GetCart(userID)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

// AddItemCartByID implements ICartService.
func (c *cartService) AddItemCartByID(userID string, item *models.CartItem) error {

	err := c.cartRepo.AddToCart(userID, item)

	if err != nil {
		return err
	}

	return nil
}

func NewCartService(cr repositories.ICartRepository) ICartService {
	return &cartService{
		cartRepo: cr,
	}
}
