package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secret []byte

func init() {
	// Load .env saat package helper diinisialisasi
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Ambil SECRET_KEY dari .env
	secret = []byte(os.Getenv("ENCRYPTED_SECRET_KEY"))
}

type JWTClaims struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := JWTClaims{
		Id:        id,
		Username:  username,
		IsRefresh: false, // Menandakan ini adalah token akses, bukan refresh token
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key dan kembalikan string token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	// Kembalikan token dan waktu kedaluwarsa
	return tokenString, expirationTime, nil
}

// RefreshToken digunakan untuk membuat token refresh
func RefreshToken(tokenString string) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token refresh berlaku selama 24 jam
	claims := JWTClaims{
		IsRefresh: true, // Menandakan ini adalah token refresh
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key dan kembalikan string token
	tokenStringReturn, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	// Kembalikan token dan waktu kedaluwarsa
	return tokenStringReturn, expirationTime, nil
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, errors.New("unexpected signing method")
		// }
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	if claims.IsRefresh {
		return nil, errors.New("token is a refresh token, not an access token")
	}

	return claims, nil
}
