package application

// Request DTOs
type CreateOrderFromCartDTO struct {
	UserID string `json:"user_id" validate:"required"`
}

type CancelOrderDTO struct {
	OrderID string `json:"order_id" validate:"required"`
}

// Response DTOs
type OrderItemDTO struct {
	ID                string `json:"id"`
	BookID            string `json:"book_id"`
	BookTitleSnapshot string `json:"book_title"`
	PriceSnapshot     string `json:"price"`
	Quantity          int    `json:"quantity"`
	CreatedAt         string `json:"created_at"`
}

type OrderDTO struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Items       []OrderItemDTO `json:"items"`
	TotalAmount string         `json:"total_amount"`
	Status      string         `json:"status"`
	ItemCount   int            `json:"item_count"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

type CreateOrderResponseDTO struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Items       []OrderItemDTO `json:"items"`
	TotalAmount string         `json:"total_amount"`
	Status      string         `json:"status"`
	ItemCount   int            `json:"item_count"`
	CreatedAt   string         `json:"created_at"`
}

type ListOrdersDTO struct {
	Orders     []OrderDTO `json:"orders"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
}
