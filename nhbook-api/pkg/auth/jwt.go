package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define claims JWT
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Generate JWT token
func GenerateJWT(userID, role, secretKey string, expiration time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// Verify and parse JWT
func ParseJWT(tokenStr, secretKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
