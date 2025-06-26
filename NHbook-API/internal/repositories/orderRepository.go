package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(tx *gorm.DB, order *models.Order) error
	CreateOrderItem(tx *gorm.DB, orderItem *models.OrderItem) error
	GetOrderByID(id string) (*models.Order, error)
	GetOrderItemByID(id string) (*models.OrderItem, error)
}

type orderRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

// CreateOrderItem implements IOrderRepository.
func (o *orderRepository) CreateOrderItem(tx *gorm.DB, orderItem *models.OrderItem) error {
	return tx.Create(&orderItem).Error

}

// GetOrderByID implements IOrderRepository.
func (o *orderRepository) GetOrderByID(id string) (*models.Order, error) {
	panic("unimplemented")
}

// GetOrderItemByID implements IOrderRepository.
func (o *orderRepository) GetOrderItemByID(id string) (*models.OrderItem, error) {
	panic("unimplemented")
}

// CreateOrder implements IOrderRepository.
func (o *orderRepository) CreateOrder(tx *gorm.DB, order *models.Order) error {
	return tx.Create(&order).Error
}

func NewOrderRepository(db *gorm.DB, rd *redis.Client) IOrderRepository {
	return &orderRepository{
		db:          db,
		redisClient: rd,
	}
}
