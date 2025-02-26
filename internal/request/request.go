package request

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/internal/auth"
)

// Request handles HTTP requests to the Lalamove API
type Request struct {
	client  *http.Client
	creds   auth.Credentials
	baseURL string
	market  string
}

// NewRequest creates a new request handler
func NewRequest(client *http.Client, creds auth.Credentials, baseURL, market string) *Request {
	return &Request{
		client:  client,
		creds:   creds,
		baseURL: baseURL,
		market:  market,
	}
}

func (r *Request) Do(ctx context.Context, method, path string, body, response interface{}) error {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	var bodyStr string
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyStr = string(b)
	}

	// Generate signature
	rawSignature := fmt.Sprintf("%s\r\n%s\r\n%s\r\n\r\n%s", timestamp, method, path, bodyStr)
	h := hmac.New(sha256.New, []byte(r.creds.APISecret))
	h.Write([]byte(rawSignature))
	signature := hex.EncodeToString(h.Sum(nil))

	// Create request
	var bodyReader *bytes.Reader
	if body != nil {
		bodyReader = bytes.NewReader([]byte(bodyStr))
	}

	req, err := http.NewRequestWithContext(ctx, method, r.baseURL+path, bodyReader)
	if err != nil {
		return err
	}

	// Set headers
	token := fmt.Sprintf("%s:%s:%s", r.creds.APIKey, timestamp, signature)
	req.Header.Set("Authorization", fmt.Sprintf("hmac %s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Market", r.market)

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}
	return nil
}
