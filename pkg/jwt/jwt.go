package jwt

import (
	"errors"
	"os"
	"strings"
	"time"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID uint 	`json:"user_id"`
	Email string 	`json:"email"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, email string) (string, error) {
	env := godotenv.Load()
	if env != nil {
		return "", errors.New("failed to load .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	expiryStr := os.Getenv("JWT_ACCESS_EXPIRY")
	expiry, _ := time.ParseDuration(expiryStr)

	claims := &Claims {
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(userID uint, email string) (string, error) {
	env := godotenv.Load()
	if env != nil {
		return "", errors.New("failed to load .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	expiryStr := os.Getenv("JWT_REFRESH_EXPIRY")
	expiry, _ := time.ParseDuration(expiryStr)

	claims := &Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (*Claims, error) {
	env := godotenv.Load()
	if env != nil {
		return nil, errors.New("failed to load .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractToken(r *http.Request) string {
	bearerTolen := r.Header.Get("Authorization")

	if len(strings.Split(bearerTolen, "")) == 2 {
		return strings.Split(bearerTolen, "")[1]
	}

	return ""
}

func ExtractUserID(r *http.Request) (uint, error) {
	tokenString := ExtractToken(r)

	if tokenString == "" {
		return 0, errors.New("no token provided")
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}