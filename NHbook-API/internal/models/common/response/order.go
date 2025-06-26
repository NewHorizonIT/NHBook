package response

type OrderItemResponse struct {
	BookID   uint   `json:"book_id"`
	Price    int    `json:"price"`
	BookName string `json:"book_name"`
	Quantity int    `json:"quantity"`
	Total    int    `json:"total"`
}

type OrderResponse struct {
	UserID        string              `json:"user_id"`
	PaymentMethod string              `json:"payment_method"`
	TotalAmount   int                 `json:"total_amount"`
	Status        string              `json:"status"`
	CreatedAt     string              `json:"created_at"`
	OrderItems    []OrderItemResponse `json:"order_items"`
}
