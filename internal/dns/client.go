package dns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/retryablehttp-go"
)

// Client is an http client for DNS
type Client struct {
	cl *internal.HTTPClient
}

// NewClient creates a new client for chaos API communication
func NewClient(cl *internal.HTTPClient) *Client {
	return &Client{cl: cl}
}

// LookupResponse is an api response
type LookupResponse struct {
	A          []string `json:"a,omitempty"`
	Aaaa       []string `json:"aaaa,omitempty"`
	All        []string `json:"all,omitempty"`
	Caa        []string `json:"caa,omitempty"`
	Cname      []string `json:"cname,omitempty"`
	Mx         []string `json:"mx,omitempty"`
	Ns         []string `json:"ns,omitempty"`
	Resolver   []string `json:"resolver,omitempty"`
	Soa        []string `json:"soa,omitempty"`
	Srv        []string `json:"srv,omitempty"`
	StatusCode string   `json:"statusCode,omitempty"`
	TTL        int      `json:"ttl,omitempty"`
	Txt        []string `json:"txt,omitempty"`
}

// Lookup requests a domain resolution
func (c *Client) Lookup(domain string, types []RecordType) (*LookupResponse, error) {
	request, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s/dnsx/%s", internal.APIAddress, domain), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	resp, err := c.cl.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not make request: wrong status %d", resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	res := LookupResponse{}
	err = d.Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("could not decode request: %w", err)
	}
	if types == nil {
		return &res, nil
	}

	out := LookupResponse{}
	for _, t := range types {
		switch t {
		case RecordTypeA:
			out.A = res.A
		case RecordTypeAAAA:
			out.Aaaa = res.Aaaa
		case RecordTypeCAA:
			out.Caa = res.Caa
		case RecordTypeCNAME:
			out.Cname = res.Cname
		case RecordTypeMX:
			out.Mx = res.Mx
		case RecordTypeNS:
			out.Ns = res.Ns
		case RecordTypeSOA:
			out.Soa = res.Soa
		case RecordTypeSRV:
			out.Srv = res.Srv
		case RecordTypeTXT:
			out.Txt = res.Txt
		}
	}
	return &out, nil
}
