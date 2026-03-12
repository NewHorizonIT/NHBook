package errs

type ErrorCode string

const (
	// Error codes
	CodeUnauthorized      ErrorCode = "UNAUTHORIZED"
	CodeInvalidInput      ErrorCode = "INVALID_INPUT"
	CodeValidationError   ErrorCode = "VALIDATION_ERROR"
	CodeInternalServer    ErrorCode = "INTERNAL_SERVER_ERROR"
	CodeNotFound          ErrorCode = "NOT_FOUND"
	CodeConflict          ErrorCode = "CONFLICT"
	CodeTooManyRequests   ErrorCode = "TOO_MANY_REQUESTS"
	CodeInvalidPassword   ErrorCode = "INVALID_PASSWORD"
	CodeEmailAlreadyExists ErrorCode = "EMAIL_ALREADY_EXISTS"
)

type Severity string

const (
	SeverityInfo    Severity = "INFO"
	SeverityWarning Severity = "WARNING"
	SeverityError   Severity = "ERROR"
)

type DomainError struct {
	Code       ErrorCode      `json:"code"`
	Message    string         `json:"message"`
	StatusCode int            `json:"-"`
	Severity   Severity       `json:"severity"`
	Details    map[string]any `json:"details,omitempty"`
	Err        error          `json:"-"`
}

func (e *DomainError) Error() string {
	return e.Message
}

// NewDomainError creates a new domain error
func NewDomainError(code ErrorCode, message string, statusCode int) *DomainError {
	return &DomainError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Severity:   SeverityError,
		Details:    make(map[string]any),
	}
}

// WithDetail adds detail information
func (e *DomainError) WithDetail(key string, value any) *DomainError {
	e.Details[key] = value
	return e
}

// WithError adds internal error for logging
func (e *DomainError) WithError(err error) *DomainError {
	e.Err = err
	return e
}

// WithSeverity sets severity level
func (e *DomainError) WithSeverity(severity Severity) *DomainError {
	e.Severity = severity
	return e
}
