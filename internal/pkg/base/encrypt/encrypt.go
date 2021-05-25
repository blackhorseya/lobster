package encrypt

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// EncPWD serve caller to encrypt password to irent format
func EncPWD(pwd string) string {
	return fmt.Sprintf("0x%x", sha256.Sum256([]byte(pwd)))
}

// HashAndSalt serve caller to given password to hash and salt encrypt
func HashAndSalt(pwd string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
