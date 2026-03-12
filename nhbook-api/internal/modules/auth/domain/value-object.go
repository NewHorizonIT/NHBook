package domain

import "regexp"

// Email Value Object
type Email string

func NewEmail(email string) (Email, error) {
	// Check empty
	if email == "" {
		return "", ErrEmailIsRequired
	}

	// Check format
	const emailPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	var emailRegex = regexp.MustCompile(emailPattern)
	if emailRegex.MatchString(email) == false {
		return "", ErrInvalidEmailFormat
	}

	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

// Status Value Object
type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)

func NewStatus(status string) Status {
	switch status {
	case string(StatusInactive):
		return StatusInactive
	default:
		return StatusActive
	}
}

func (s Status) String() string {
	return string(s)
}

// Role Value Object
type Role string

const (
	RoleUser  Role = "CUSTOMER"
	RoleAdmin Role = "ADMIN"
)

func NewRole(role string) Role {
	switch role {
	case string(RoleAdmin):
		return RoleAdmin
	default:
		return RoleUser
	}
}

func (r Role) String() string {
	return string(r)
}
