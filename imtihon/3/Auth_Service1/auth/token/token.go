package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenstr string) (bool, error) {
	_, err := ExtractClaims(tokenstr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenstr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("SIGNING_KEY"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func GenerateJWTToken(userId, email string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Hour).Unix(), // Expires in 1 hour
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the token
	return token.SignedString([]byte("SIGNING_KEY"))
}
func GenerateRefreshToken(userId, email string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Expires in 24 hours
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the token
	return token.SignedString([]byte("SIGNING_KEY"))
}
