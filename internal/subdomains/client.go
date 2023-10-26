package subdomains

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/retryablehttp-go"
	pdhttputil "github.com/projectdiscovery/utils/http"
	"io"
	"net/http"
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
	request, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("https://api.chaosdb.sh/dns/%s", req.Domain), nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request.")
	}

	resp, err := c.cl.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "could not make request.")
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read response.")
		}
		return nil, internal.InvalidStatusCodeError{StatusCode: resp.StatusCode, Message: body}
	}

	defer pdhttputil.DrainResponseBody(resp)

	response := GetStatisticsResponse{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal results.")
	}

	return &response, nil
}

// SubdomainsRequest is the request for a host subdomains.
type SubdomainsRequest struct {
	Domain       string
	OutputFormat string
}

// Result is the response for a host subdomains.
type Result struct {
	Subdomain string
	Reader    *io.ReadCloser
	Error     error
}

// GetSubdomains returns the subdomains for a given domain.
func (c *Client) GetSubdomains(req *SubdomainsRequest) chan *Result {
	results := make(chan *Result)
	go func(results chan *Result) {
		defer close(results)

		request, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("https://api.chaosdb.sh/dns/%s/subdomains", req.Domain), nil)
		if err != nil {
			results <- &Result{Error: errors.Wrap(err, "could not create request.")}
			return
		}

		resp, err := c.cl.Do(request)
		if err != nil {
			results <- &Result{Error: errors.Wrap(err, "could not make request.")}
			return
		}

		if resp.StatusCode != http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				results <- &Result{Error: errors.Wrap(err, "could not read response.")}
				return
			}
			pdhttputil.DrainResponseBody(resp)
			results <- &Result{Error: internal.InvalidStatusCodeError{StatusCode: resp.StatusCode, Message: body}}
			return
		}

		switch req.OutputFormat {
		case "json":
			results <- &Result{Reader: &resp.Body}
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
				results <- &Result{Subdomain: fmt.Sprintf("%s", token)}
			}
		}
	}(results)

	return results
}
