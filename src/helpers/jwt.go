package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type (
	ParamsGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       string
		ClientID        string
		ApplicationID   string
		ApplicationName string
		ApiKey          string
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}
	Claims struct {
		jwt.StandardClaims
		ClientID        string `json:"clientID,omitempty"`
		ApplicationID   string `json:"applicationID,omitempty"`
		ApplicationName string `json:"applicationName,omitempty"`
		ApiKey          string `json:"apiKey,omitempty"`
	}
)

// GenerateJWT generates a new JWT token.
//
// The function takes a pointer to a ParamsGenerateJWT struct as an argument, which contains the necessary parameters for generating the JWT.
// These parameters include the expiration time in minutes, the secret key, client ID, application ID, application name, and API key.
//
// The function returns three values: the generated JWT as a string, the Unix timestamp of when the token will expire, and an error value.
// If the token is generated successfully, the error value will be nil. Otherwise, it will contain information about what went wrong.
func GenerateJWT(p *ParamsGenerateJWT) (string, int64, error) {
	expiredAt := time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix()
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
		ClientID:        p.ClientID,
		ApplicationID:   p.ApplicationID,
		ApplicationName: p.ApplicationName,
		ApiKey:          p.ApiKey,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(p.SecretKey))

	return signedToken, expiredAt, err
}

// ValidateJWT validates a JSON Web Token (JWT) using the provided parameters.
// It parses the token, verifies the signing method, and returns the claims if the token is valid.
// If the token is invalid or an error occurs during parsing or verification, an error is returned.
func ValidateJWT(p *ParamsValidateJWT) (jwt.MapClaims, error) {
	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != JWT_SIGNING_METHOD {
			return nil, errors.New("Invalid Token")
		}

		return []byte(p.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
