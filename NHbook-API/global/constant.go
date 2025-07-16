package global

// Defined Header key
const (
	HEADER_API_KEY       = "X-Api-Key"
	HEADER_AUTHORIZATION = "Authorization"
)

// Defined variable status order

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusShipped   = "shipped"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

// Define variable payment method order

const (
	OrderPaymentMethodCOD  = "cod"
	OrderPaymentMethodMOMO = "momo"
	OrderPaymentMethodBank = "bank"
)
