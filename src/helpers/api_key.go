package helpers

import (
	"os"
)

// GenerateApiKey generates an API key by encrypting the client secret using AES encryption.
// It retrieves the secret key from the environment variable ACCESS_TOKEN_SECRET_KEY.
// The first 16 characters of the secret key are used for AES encryption.
//
// Parameters:
// - clientSecret: The client secret to be encrypted.
//
// Returns:
// - The encrypted API key.
// - An error if encryption fails.
func GenerateApiKey(clientSecret string) (string, error) {
	secretKey := os.Getenv("ACCESS_TOKEN_SECRET_KEY")

	//ambil 16 karakter dari secretKey untuk kebutuhan AES Encryption
	sKey := secretKey[:16]

	encryptedText, err := AESEncrypt(clientSecret, sKey)

	if err != nil {
		return "", err
	}

	return encryptedText, nil
}

// DecryptApiKey decrypts the given encrypted API key using AES encryption.
// It retrieves the secret key from the environment variable ACCESS_TOKEN_SECRET_KEY.
// The first 16 characters of the secret key are used for AES encryption.
// If the decryption is successful, it returns the plaintext API key.
// Otherwise, it returns an error.
func DecryptApiKey(encryptedText string) (string, error) {
	secretKey := os.Getenv("ACCESS_TOKEN_SECRET_KEY")

	//ambil 16 karakter dari secretKey untuk kebutuhan AES Encryption
	sKey := secretKey[:16]

	plaintText, err := AESDecrypt(encryptedText, sKey)
	if err != nil {
		return "", err
	}

	return plaintText, nil
}
