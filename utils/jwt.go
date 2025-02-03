package utils

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWTSecret() []byte {
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) == 0 {
		panic("JWT_SECRET environment variable not set")
	}
	return []byte(jwtSecret)
}

// GenerateAccessToken creates a signed JWT access token with a 15-minute expiry.
func GenerateAccessToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString(GetJWTSecret())
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32) // 256-bit token
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
