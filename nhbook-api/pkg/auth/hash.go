package auth

import "golang.org/x/crypto/bcrypt"

// Hash passwords
func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

// Compare passwords
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
