package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash compares a given password with a hash and returns true if they match, false otherwise.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
// The password parameter is the plain text password to be checked.
// The hash parameter is the bcrypt hash to compare against.
// It returns a boolean value indicating whether the password matches the hash.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
