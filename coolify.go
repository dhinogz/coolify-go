package coolifygo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client .
type Client struct {
	apiKey     string
	baseURL    string
	HTTPClient *http.Client
}

// NewClient creates new Coolify client with given API key
func NewClient(coolifyAddr, apiKey string) *Client {
	baseUrl := fmt.Sprintf("%s/api/v1", coolifyAddr)
	return &Client{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: baseUrl,
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

// Content-type and body should be already added to req
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
