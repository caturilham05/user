package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func PasswordVerify(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
