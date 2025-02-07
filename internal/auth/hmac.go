package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Credentials holds the API credentials
type Credentials struct {
	APIKey    string
	APISecret string
}

// SignRequest creates an HMAC signature for the request
func SignRequest(creds Credentials, method, path string, body []byte) (string, string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	
	// Create raw signature string
	rawSignature := fmt.Sprintf("%s\r\n%s\r\n%s\r\n\r\n%s",
		timestamp,
		strings.ToUpper(method),
		path,
		string(body),
	)

	// Create HMAC SHA256 hash
	h := hmac.New(sha256.New, []byte(creds.APISecret))
	h.Write([]byte(rawSignature))
	signature := hex.EncodeToString(h.Sum(nil))

	// Create authorization token
	token := fmt.Sprintf("%s:%s:%s", creds.APIKey, timestamp, signature)

	return token, timestamp, nil
}

// ValidateAPIKey validates the API key format
func ValidateAPIKey(apiKey string) error {
	if !strings.HasPrefix(apiKey, "pk_test_") && !strings.HasPrefix(apiKey, "pk_prod_") {
		return fmt.Errorf("invalid API key format: must start with pk_test_ or pk_prod_")
	}
	return nil
}

// ValidateAPISecret validates the API secret format
func ValidateAPISecret(apiSecret string) error {
	if !strings.HasPrefix(apiSecret, "sk_test_") && !strings.HasPrefix(apiSecret, "sk_prod_") {
		return fmt.Errorf("invalid API secret format: must start with sk_test_ or sk_prod_")
	}
	return nil
} 