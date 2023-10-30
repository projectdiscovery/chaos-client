package subdomains

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/retryablehttp-go"
	pdhttputil "github.com/projectdiscovery/utils/http"
)

// GetStatisticsRequest is the request for a domain statistics
type GetStatisticsRequest struct {
	Domain string
}

// GetStatisticsResponse is the response for a statistics request
type GetStatisticsResponse struct {
	Subdomains uint64 `json:"subdomains"`
}

type Client struct {
	cl *internal.HTTPClient
}

// NewClient creates a new client for chaos API communication
func NewClient(cl *internal.HTTPClient) *Client {
	return &Client{cl: cl}
}

// GetStatistics returns the statistics for a given domain.
func (c *Client) GetStatistics(req *GetStatisticsRequest) (*GetStatisticsResponse, error) {
	request, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s/dns/%s", internal.APIAddress, req.Domain), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	resp, err := c.cl.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read response: %w", err)
		}
		return nil, internal.InvalidStatusCodeError{StatusCode: resp.StatusCode, Err: string(body)}
	}

	defer pdhttputil.DrainResponseBody(resp)

	response := GetStatisticsResponse{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal results: %w", err)
	}

	return &response, nil
}

// Request is the request for a host subdomains.
type Request struct {
	Domain       string
	OutputFormat string
}

// Response is the response for a host subdomains.
type Response struct {
	Subdomain string
	Reader    *io.ReadCloser
	Error     error
}

// GetSubdomains returns the subdomains for a given domain.
func (c *Client) GetSubdomains(req *Request) chan *Response {
	results := make(chan *Response)
	go func(results chan *Response) {
		defer close(results)

		request, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s/dns/%s/subdomains", internal.APIAddress, req.Domain), nil)
		if err != nil {
			results <- &Response{Error: fmt.Errorf("could not create request: %w", err)}
			return
		}

		resp, err := c.cl.Do(request)
		if err != nil {
			results <- &Response{Error: fmt.Errorf("could not make request: %w", err)}
			return
		}

		if resp.StatusCode != http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				results <- &Response{Error: fmt.Errorf("could not read response: %w", err)}
				return
			}
			pdhttputil.DrainResponseBody(resp)
			results <- &Response{Error: internal.InvalidStatusCodeError{StatusCode: resp.StatusCode, Err: string(body)}}
			return
		}

		switch req.OutputFormat {
		case "json":
			results <- &Response{Reader: &resp.Body}
		default:
			defer pdhttputil.DrainResponseBody(resp)
			d := json.NewDecoder(resp.Body)
			if !internal.CheckToken(d, "{") {
				return
			}
			if !internal.CheckToken(d, "domain") {
				return
			}
			if !internal.CheckToken(d, req.Domain) {
				return
			}
			if !internal.CheckToken(d, "subdomains") {
				return
			}
			if !internal.CheckToken(d, "[") {
				return
			}

			for d.More() {
				// process all the tokens within the list
				token, err := d.Token()
				if token == nil || err != nil {
					break
				}
				results <- &Response{Subdomain: fmt.Sprintf("%s", token)}
			}
		}
	}(results)

	return results
}
