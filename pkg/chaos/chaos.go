package chaos

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// Client is a client for making requests to chaos API
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// New creates a new client for chaos API communication
func New(apiKey string) *Client {
	httpclient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxIdleConns:        100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(600) * time.Second, // 10 minutes - uploads may take long
	}
	return &Client{httpClient: httpclient, apiKey: apiKey}
}

// GetStatisticsRequest is the request for a domain statistics
type GetStatisticsRequest struct {
	Domain string
}

// GetStatisticsResponse is the response for a statistics request
type GetStatisticsResponse struct {
	Subdomains uint64 `json:"subdomains"`
}

// GetStatistics returns the statistics for a given domain.
func (c *Client) GetStatistics(req *GetStatisticsRequest) (*GetStatisticsResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://dns.projectdiscovery.io/dns/%s", req.Domain), nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	request.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "could not make request")
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read response")
		}
		return nil, fmt.Errorf("invalid status code received: %d - %s", resp.StatusCode, string(body))
	}

	response := GetStatisticsResponse{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal results")
	}
	return &response, nil
}

// GetSubdomainsRequest is the request for a host subdomains.
type GetSubdomainsRequest struct {
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
func (c *Client) GetSubdomains(req *GetSubdomainsRequest) chan *Result {
	results := make(chan *Result)
	go func(results chan *Result) {
		defer close(results)

		request, err := http.NewRequest("GET", fmt.Sprintf("https://dns.projectdiscovery.io/dns/%s/subdomains", req.Domain), nil)
		if err != nil {
			results <- &Result{Error: errors.Wrap(err, "could not create request")}
			return
		}
		request.Header.Set("Authorization", c.apiKey)

		resp, err := c.httpClient.Do(request)
		if err != nil {
			results <- &Result{Error: errors.Wrap(err, "could not make request")}
			return
		}

		if resp.StatusCode != 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				results <- &Result{Error: errors.Wrap(err, "could not read response")}
				return
			}
			results <- &Result{Error: fmt.Errorf("invalid status code received: %d - %s", resp.StatusCode, string(body))}
			return
		}

		switch req.OutputFormat {
		case "json":
			results <- &Result{Reader: &resp.Body}
		default:
			d := json.NewDecoder(resp.Body)
			d.Token()
			// first 4 token should be skipped
			skip := 0
			for d.More() {
				token, _ := d.Token()
				skip++
				if skip <= 4 {
					continue
				}
				results <- &Result{Subdomain: fmt.Sprintf("%s", token)}
			}
			d.Token()
		}
	}(results)

	return results
}

// PutSubdomainsRequest is the request for uploading subdomains.
type PutSubdomainsRequest struct {
	Contents io.Reader
}

// PutSubdomainsResponse is the response for a subdomains upload request.
type PutSubdomainsResponse struct{}

// PutSubdomains uploads the subdomains to Chaos API.
func (c *Client) PutSubdomains(req *PutSubdomainsRequest) (*PutSubdomainsResponse, error) {
	request, err := http.NewRequest("POST", "https://dns.projectdiscovery.io/dns/add", req.Contents)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	request.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "could not make request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read response")
		}
		return nil, fmt.Errorf("invalid status code received: %d - %s", resp.StatusCode, string(body))
	}
	io.Copy(ioutil.Discard, resp.Body)
	return &PutSubdomainsResponse{}, nil
}
