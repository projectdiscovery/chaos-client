package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/projectdiscovery/ratelimit"
	"github.com/projectdiscovery/retryablehttp-go"
)

var APIAddress = "https://api.chaosdb.sh"

// HTTPClient is a client for making requests to chaos API
type HTTPClient struct {
	apiKey     string
	httpClient *retryablehttp.Client
	ratelimit  *ratelimit.Limiter
}

// NewHTTPClient creates a new client for chaos API communication
func NewHTTPClient(apiKey string) *HTTPClient {
	httpclient := retryablehttp.NewClient(retryablehttp.DefaultOptionsSingle)
	return &HTTPClient{httpClient: httpclient, apiKey: apiKey}
}

// do adds apiKey and implements rate limit
func (c *HTTPClient) Do(request *retryablehttp.Request) (*http.Response, error) {
	request.Header.Set("Authorization", c.apiKey)
	if c.ratelimit != nil {
		c.ratelimit.Take()
	}
	resp, err := c.httpClient.Do(request)
	if resp != nil && c.ratelimit == nil {
		rl := resp.Header.Get("X-Ratelimit-Limit")
		rlMax, err := strconv.Atoi(rl)
		if err == nil && rlMax > 0 {
			// if er then ratelimit header is not present. Hence, no rate limit
			c.ratelimit = ratelimit.New(context.Background(), uint(rlMax), time.Minute)
		}
	}
	return resp, err
}

type InvalidStatusCodeError struct {
	StatusCode int
	Err        string `json:"error"`
}

func (e InvalidStatusCodeError) Error() string {
	return fmt.Sprintf("invalid status code received: %d - %s", e.StatusCode, e.Err)
}

func CheckToken(d *json.Decoder, value string) bool {
	token, err := d.Token()
	return strings.EqualFold(fmt.Sprint(token), value) && err == nil
}
