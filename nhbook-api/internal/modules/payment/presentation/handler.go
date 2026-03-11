package presentation

import (
	"net/http"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/middlewares"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/gin-gonic/gin"
	validatorpkg "github.com/go-playground/validator/v10"
)

type PaymentHandler struct {
	createPaymentUsecase      *usecase.CreatePaymentUsecase
	processPaymentUsecase     *usecase.ProcessPaymentUsecase
	getPaymentStatusUsecase   *usecase.GetPaymentStatusUsecase
	validator                 *validatorpkg.Validate
}

func NewPaymentHandler(
	createPaymentUsecase *usecase.CreatePaymentUsecase,
	processPaymentUsecase *usecase.ProcessPaymentUsecase,
	getPaymentStatusUsecase *usecase.GetPaymentStatusUsecase,
) *PaymentHandler {
	return &PaymentHandler{
		createPaymentUsecase:    createPaymentUsecase,
		processPaymentUsecase:   processPaymentUsecase,
		getPaymentStatusUsecase: getPaymentStatusUsecase,
		validator:               validatorpkg.New(),
	}
}

// CreatePayment godoc
// @Summary      Create payment for order
// @Description  Create a payment record for an order
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        request body application.CreatePaymentDTO true "Payment creation data"
// @Success      201 {object} application.CreatePaymentResponseDTO
// @Router       /payments [post]
// @Security     BearerAuth
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req application.CreatePaymentDTO

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		).WithError(err)
		middlewares.HandleError(c, domainErr)
		return
	}

	// Validate
	if err := h.validator.Struct(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeValidationError,
			"Validation failed",
			http.StatusBadRequest,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createPaymentUsecase.Execute(ctx, &req)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}

// ProcessPayment godoc
// @Summary      Process payment
// @Description  Process a payment through payment gateway
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        request body application.ProcessPaymentDTO true "Payment ID to process"
// @Success      200 {object} application.PaymentStatusDTO
// @Router       /payments/process [post]
// @Security     BearerAuth
func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req application.ProcessPaymentDTO

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		).WithError(err)
		middlewares.HandleError(c, domainErr)
		return
	}

	// Validate
	if err := h.validator.Struct(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeValidationError,
			"Validation failed",
			http.StatusBadRequest,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	result, err := h.processPaymentUsecase.Execute(ctx, &req)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// GetPaymentStatus godoc
// @Summary      Get payment status
// @Description  Retrieve payment status by payment ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        payment_id path string true "Payment ID"
// @Success      200 {object} application.PaymentStatusDTO
// @Router       /payments/{payment_id} [get]
// @Security     BearerAuth
func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	paymentID := c.Param("payment_id")

	ctx := c.Request.Context()
	result, err := h.getPaymentStatusUsecase.Execute(ctx, paymentID)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}
