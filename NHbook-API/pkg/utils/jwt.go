package utils

import (
	"fmt"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/golang-jwt/jwt/v5"
)

// Claim of token

type Claim struct {
	UserID string
	Email  string
	jwt.RegisteredClaims
}

func CreateTokenPair(userID string, email string) (string, string, error) {
	jc := global.Config.JWT
	secretKey := []byte(jc.Secret)
	// CREATE ACCESS TOKEN
	accessExp, _ := time.ParseDuration(jc.AccessTokenExpiry)
	// Step 1: Create Claim
	accessClaims := Claim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// Step 2: Create Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	// Step 3: Create String access token
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// CREATE REFRESH TOKEN
	refreshExp, _ := time.ParseDuration(jc.RefreshTokenExpiry)
	// Step 1: Create Claim refresh token
	refreshClaim := Claim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// Step 2: Create RefreshToken
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	// Step 3: Create string refresh token
	refreshTokenString, err := refreshToken.SignedString(secretKey)

	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string) (*Claim, error) {
	jc := global.Config.JWT
	secretKey := []byte(jc.Secret)
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
