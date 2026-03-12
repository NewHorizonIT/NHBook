package domain

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID
	CartID    uuid.UUID
	BookID    uuid.UUID
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCartItem(cartID uuid.UUID, bookID uuid.UUID, quantity int) (*CartItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	return &CartItem{
		ID:        uuid.New(),
		CartID:    cartID,
		BookID:    bookID,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (ci *CartItem) UpdateQuantity(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	ci.Quantity = quantity
	ci.UpdatedAt = time.Now()
	return nil
}

type Cart struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Items     []*CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCart(userID uuid.UUID) *Cart {
	return &Cart{
		ID:        uuid.New(),
		UserID:    userID,
		Items:     make([]*CartItem, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Cart) AddItem(item *CartItem) error {
	// Check if item already exists
	for _, existingItem := range c.Items {
		if existingItem.BookID == item.BookID {
			// Update quantity instead
			return existingItem.UpdateQuantity(existingItem.Quantity + item.Quantity)
		}
	}

	c.Items = append(c.Items, item)
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Cart) RemoveItem(bookID uuid.UUID) error {
	for i, item := range c.Items {
		if item.BookID == bookID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			c.UpdatedAt = time.Now()
			return nil
		}
	}
	return ErrItemNotFound
}

func (c *Cart) UpdateItemQuantity(bookID uuid.UUID, quantity int) error {
	for _, item := range c.Items {
		if item.BookID == bookID {
			return item.UpdateQuantity(quantity)
		}
	}
	return ErrItemNotFound
}

func (c *Cart) GetItemCount() int {
	return len(c.Items)
}

func (c *Cart) GetTotalQuantity() int {
	total := 0
	for _, item := range c.Items {
		total += item.Quantity
	}
	return total
}

func (c *Cart) Clear() {
	c.Items = make([]*CartItem, 0)
	c.UpdatedAt = time.Now()
}

func (c *Cart) IsEmpty() bool {
	return len(c.Items) == 0
}
