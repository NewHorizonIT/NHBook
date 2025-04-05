package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/redis/go-redis/v9"
)

type ICartRepository interface {
	GetCart(userID string) ([]models.CartItem, error)
	AddToCart(userID string, item *models.CartItem) error
	RemoveItemInCart(userID string, productID int) error
	DeleteCart(userID string) error
}

type cartRepository struct {
	redisClient *redis.Client
}

// AddCart implements ICartRepository.
func (c *cartRepository) AddToCart(userID string, item *models.CartItem) error {
	ctx := context.Background()
	redisKey := fmt.Sprintf("cart:%v", userID)
	// Check IsExists
	cartData, err := c.redisClient.Get(ctx, redisKey).Result()
	var cart []models.CartItem

	if err != nil {
		newCart, _ := json.Marshal(cart)
		c.redisClient.Set(ctx, redisKey, newCart, 7*24*time.Hour)
		return err
	}

	json.Unmarshal([]byte(cartData), &cart)

	// Check items exists in cart. If it exits then update
	updated := false

	for i, v := range cart {
		if v.ID == item.ID {
			cart[i].Quantity += item.Quantity
			cart[i].Total += float64(item.Quantity) * item.Price
			updated = true
		}
	}

	if !updated {
		cart = append(cart, *item)
	}

	newCart, _ := json.Marshal(cart)
	// Set cart into redis
	c.redisClient.Set(ctx, redisKey, newCart, 7*24*time.Hour)

	return nil

}

// DeleteCart implements ICartRepository.
func (c *cartRepository) DeleteCart(userID string) error {
	ctx := context.Background()
	keyRedis := fmt.Sprintf("cart:%v", userID)

	cartData, err := c.redisClient.Get(ctx, keyRedis).Result()

	if err != nil {
		return err
	}
	cartData = ""

	c.redisClient.Set(ctx, keyRedis, &cartData, 7*24*time.Hour)
	return nil

}

// GetCart implements ICartRepository.
func (c *cartRepository) GetCart(userID string) ([]models.CartItem, error) {
	ctx := context.Background()
	keyRedis := fmt.Sprintf("cart:%v", userID)
	var cart []models.CartItem

	cartData, err := c.redisClient.Get(ctx, keyRedis).Result()

	if err != nil {
		return nil, err
	}

	if cartData == "" {
		return nil, err
	}

	if err := json.Unmarshal([]byte(cartData), &cart); err != nil {
		return nil, err
	}

	return cart, nil

}

// RemoveCart implements ICartRepository.
func (c *cartRepository) RemoveItemInCart(userID string, productID int) error {
	ctx := context.Background()
	keyRedis := fmt.Sprintf("cart:%s", userID)

	cartData, err := c.redisClient.Get(ctx, keyRedis).Result()

	if err != nil {
		return err
	}

	var cart []models.CartItem
	if err := json.Unmarshal([]byte(cartData), &cart); err != nil {
		return err
	}

	for i, v := range cart {
		if v.ID == uint(productID) {
			cart = append(cart[:i], cart[i+1:]...)
		}
	}

	updatedCart, err := json.Marshal(cart)

	if err != nil {
		return err
	}

	result := c.redisClient.Set(ctx, keyRedis, updatedCart, 7*24*time.Hour)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}

func NewCartRepository(rd *redis.Client) ICartRepository {
	return &cartRepository{
		redisClient: rd,
	}
}
