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

// TTL for token
const (
	ACCESS_TOKEN_TTL  = 15 * 60          // 15 minutes
	REFRESH_TOKEN_TTL = 7 * 24 * 60 * 60 // 7 days
)

// Key cookie
const (
	REFRESH_TOKEN_COOKIE = "refresh-token"
)
