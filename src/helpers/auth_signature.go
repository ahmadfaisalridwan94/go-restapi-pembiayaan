package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type (
	AuthComponent struct {
		HttpMethod   string
		RequestPath  string
		AccessToken  string
		ClientSecret string
		RequestBody  string
		Signature    string
		Timestamp    string
	}
)

// GetAuthHeader generates the authentication header for making API requests.
// It takes the HTTP method, request URL, access token, secret key, and request body as input.
// It returns a map containing the timestamp, payload, and HMAC signature for the authentication header.
func GetAuthHeader(httpMethod, requestUrl, accessToken, secretKey, requestBody string) map[string]string {
	requestPath := getPath(requestUrl)
	//queryString := getQueryString(requestUrl)

	if httpMethod == "GET" || requestBody == "" {
		requestBody = ""
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	payload := fmt.Sprintf("path=%s&verb=%s&token=Bearer %s&timestamp=%s&body=%s",
		requestPath, httpMethod, accessToken, timestamp, requestBody)

	// Debug prints
	// fmt.Println("Payload: " + payload)

	hmacSignature := generateHMACSignature(payload, secretKey)

	return map[string]string{
		"timestamp":     timestamp,
		"payload":       payload,
		"hmacSignature": hmacSignature,
	}
}

func CheckAuthSignature(a *AuthComponent) bool {

	payload := fmt.Sprintf("path=%s&verb=%s&token=Bearer %s&timestamp=%s&body=%s",
		a.RequestPath, a.HttpMethod, a.AccessToken, a.Timestamp, a.RequestBody)

	hmacSignature := generateHMACSignature(payload, a.ClientSecret)

	// fmt.Println("PAYLOAD: ", payload)
	// fmt.Println("REQUEST HEADER SIGNATURE: ", a.Signature)
	// fmt.Println("GENERATED SIGNATURE: ", hmacSignature)
	// fmt.Println("AUTH SIGNATURE: ")
	// fmt.Println(a)

	if a.Signature == hmacSignature {
		return true
	}

	return false
}

func getPath(url string) string {
	pathRegex := regexp.MustCompile(`.+?\:\/\/.+?(/.+?)(?:#|\?|$)`)
	result := pathRegex.FindStringSubmatch(url)
	if result != nil && len(result) > 1 {
		return result[1]
	}
	return ""
}

func getQueryString(url string) string {
	arrSplit := strings.Split(url, "?")
	if len(arrSplit) > 1 {
		return url[strings.Index(url, "?")+1:]
	}
	return ""
}

func generateHMACSignature(data, key string) string {
	hmacHash := hmac.New(sha256.New, []byte(key))
	hmacHash.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(hmacHash.Sum(nil))
}
