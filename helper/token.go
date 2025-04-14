package helper

import (
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
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := JWTClaims{
		Id:       id,
		Username: username,
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

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
