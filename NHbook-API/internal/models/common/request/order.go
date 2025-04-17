package request

type OrderRequest struct {
	UserID        string `json:"user_id"`
	PaymentMethod string `json:"payment_method" binding:"required,oneof=cod momo bank"`
}
