package utils

import (
	"fmt"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/golang-jwt/jwt/v5"
)

// Convert key to slice byte
var (
	jwtConfig = global.JWT
	secretKey = []byte(jwtConfig.Secret)
)

// Claim of token

type Claim struct {
	UserID string
	Email  string
	jwt.RegisteredClaims
}

func CreateTokenPair(userID string, email string) (string, string, error) {

	accessExp, _ := time.ParseDuration(jwtConfig.AccessTokenExpiry)

	accessClaims := Claim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, accessClaims)
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	refreshExp, _ := time.ParseDuration(jwtConfig.RefreshTokenExpiry)

	refreshClaim := Claim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(secretKey)

	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
