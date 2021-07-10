package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt serve caller to given password to hash and salt encrypt
func HashAndSalt(pwd string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
