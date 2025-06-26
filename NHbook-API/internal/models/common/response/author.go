package response

import "time"

type AuthorResponse struct {
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	BirthDate time.Time `json:"birthDate"`
}
