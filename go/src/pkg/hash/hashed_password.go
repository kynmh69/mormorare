package hash

import (
	"github.com/kynmh69/mormorare/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	logger := logging.GetLogger()
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	strHashedPass := string(hashedPassword)
	logger.Debug("Hashed password: ", strHashedPass)
	return strHashedPass, nil
}

func ComparePassword(hashedPassword, password string) error {
	// Compare password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
