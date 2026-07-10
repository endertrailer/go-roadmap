package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}



func GenerateToken(userID string) (string, error) {
	expireTime := time.Now().Add(2 * time.Hour)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET variable is not set")
	}
	parsed, err := uuid.FromString(userID)
	if err != nil {
		return "", fmt.Errorf("failed parsing the uuid probably a invalid uuid")
	}

	claims := &CustomClaims{UserID: parsed, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "endertrailer", IssuedAt: jwt.NewNumericDate(time.Now()),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		// TODO: Replace informal error message with proper error wrapping (e.g., fmt.Errorf("failed to sign token: %w", err))
		return "", fmt.Errorf("you cooked lil bro")
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET variable is not set")
	}

	// Parse takes the token string, a pointer to your claims struct object,
	// and a callback function that returns the secret key used to sign it.
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (HS256) before validating the signature
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Verify the token is valid (signature matches and claims like exp are verified)
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
