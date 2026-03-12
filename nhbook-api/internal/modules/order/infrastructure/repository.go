package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/infrastructure/sqlc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewOrderRepository(db *sql.DB) domain.IOrderRepository {
	return &OrderRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

// CreateOrder creates a new order
func (r *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	// Start transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queries := r.q.WithTx(tx)

	// Create order
	err = queries.CreateOrder(ctx, sqlc.CreateOrderParams{
		ID:          order.ID,
		UserID:      uuid.NullUUID{UUID: order.UserID, Valid: true},
		TotalAmount: order.TotalAmount.String(),
		Status:      string(order.Status),
		CreatedAt:   sql.NullTime{Time: order.CreatedAt, Valid: true},
		UpdatedAt:   sql.NullTime{Time: order.UpdatedAt, Valid: true},
	})
	if err != nil {
		return err
	}

	// Create order items
	for _, item := range order.Items {
		err = queries.CreateOrderItem(ctx, sqlc.CreateOrderItemParams{
			ID:                 item.ID,
			OrderID:            uuid.NullUUID{UUID: order.ID, Valid: true},
			BookID:             uuid.NullUUID{UUID: item.BookID, Valid: true},
			BookTitleSnapshot:  sql.NullString{String: item.BookTitleSnapshot, Valid: true},
			PriceSnapshot:      sql.NullString{String: item.PriceSnapshot.String(), Valid: true},
			Quantity:           int32(item.Quantity),
			CreatedAt:          sql.NullTime{Time: item.CreatedAt, Valid: true},
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func parseDecimal(s string) (interface{}, error) {
	return s, nil
}

// GetOrderByID retrieves order by ID with all items
func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	// Get order
	dbOrder, err := r.q.GetOrderByID(ctx, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	// Get order items
	dbItems, err := r.q.GetOrderItems(ctx, uuid.NullUUID{UUID: orderID, Valid: true})
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Convert to domain entities
	totalAmount, _ := decimal.NewFromString(dbOrder.TotalAmount)
	order := &domain.Order{
		ID:          dbOrder.ID,
		UserID:      dbOrder.UserID.UUID,
		Items:       make([]*domain.OrderItem, 0),
		TotalAmount: totalAmount,
		Status:      domain.OrderStatus(dbOrder.Status),
		CreatedAt:   dbOrder.CreatedAt.Time,
		UpdatedAt:   dbOrder.UpdatedAt.Time,
	}

	for _, dbItem := range dbItems {
		priceSnapshot, _ := decimal.NewFromString(dbItem.PriceSnapshot.String)
		item := &domain.OrderItem{
			ID:                dbItem.ID,
			OrderID:           dbItem.OrderID.UUID,
			BookID:            dbItem.BookID.UUID,
			BookTitleSnapshot: dbItem.BookTitleSnapshot.String,
			PriceSnapshot:     priceSnapshot,
			Quantity:          int(dbItem.Quantity),
			CreatedAt:         dbItem.CreatedAt.Time,
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

// GetOrdersByUserID retrieves all orders for a user
func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Order, int, error) {
	// Get total count
	countResult, err := r.q.GetOrdersCountByUserID(ctx, uuid.NullUUID{UUID: userID, Valid: true})
	if err != nil {
		return nil, 0, err
	}

	// Get orders
	dbOrders, err := r.q.GetOrdersByUserID(ctx, sqlc.GetOrdersByUserIDParams{
		UserID: uuid.NullUUID{UUID: userID, Valid: true},
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil && err != sql.ErrNoRows {
		return nil, 0, err
	}

	orders := make([]*domain.Order, 0)
	for _, dbOrder := range dbOrders {
		// Get items for each order
		dbItems, err := r.q.GetOrderItems(ctx, uuid.NullUUID{UUID: dbOrder.ID, Valid: true})
		if err != nil && err != sql.ErrNoRows {
			return nil, 0, err
		}

		totalAmount, _ := decimal.NewFromString(dbOrder.TotalAmount)
		order := &domain.Order{
			ID:          dbOrder.ID,
			UserID:      dbOrder.UserID.UUID,
			Items:       make([]*domain.OrderItem, 0),
			TotalAmount: totalAmount,
			Status:      domain.OrderStatus(dbOrder.Status),
			CreatedAt:   dbOrder.CreatedAt.Time,
			UpdatedAt:   dbOrder.UpdatedAt.Time,
		}

		for _, dbItem := range dbItems {
			priceSnapshot, _ := decimal.NewFromString(dbItem.PriceSnapshot.String)
			item := &domain.OrderItem{
				ID:                dbItem.ID,
				OrderID:           dbItem.OrderID.UUID,
				BookID:            dbItem.BookID.UUID,
				BookTitleSnapshot: dbItem.BookTitleSnapshot.String,
				PriceSnapshot:     priceSnapshot,
				Quantity:          int(dbItem.Quantity),
				CreatedAt:         dbItem.CreatedAt.Time,
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, int(countResult), nil
}

// UpdateOrderStatus updates order status
func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status domain.OrderStatus) error {
	return r.q.UpdateOrderStatus(ctx, sqlc.UpdateOrderStatusParams{
		Status:    string(status),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:        orderID,
	})
}

// AddOrderItem adds item to order
func (r *OrderRepository) AddOrderItem(ctx context.Context, orderID uuid.UUID, item *domain.OrderItem) error {
	return r.q.CreateOrderItem(ctx, sqlc.CreateOrderItemParams{
		ID:                item.ID,
		OrderID:           uuid.NullUUID{UUID: orderID, Valid: true},
		BookID:            uuid.NullUUID{UUID: item.BookID, Valid: true},
		BookTitleSnapshot: sql.NullString{String: item.BookTitleSnapshot, Valid: true},
		PriceSnapshot:     sql.NullString{String: item.PriceSnapshot.String(), Valid: true},
		Quantity:          int32(item.Quantity),
		CreatedAt:         sql.NullTime{Time: item.CreatedAt, Valid: true},
	})
}

// CancelOrder cancels an order
func (r *OrderRepository) CancelOrder(ctx context.Context, orderID uuid.UUID) error {
	return r.q.UpdateOrderStatus(ctx, sqlc.UpdateOrderStatusParams{
		Status:    string(domain.OrderStatusCancelled),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:        orderID,
	})
}
