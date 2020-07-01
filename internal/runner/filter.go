package runner

import (
	"fmt"
	"strings"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
)

type DNSStatusCode int

const (
	NOERROR DNSStatusCode = iota
	NXDOMAIN
	SERVFAIL
	REFUSED
	ANYDNSCODE
)

type DNSRecordType int

const (
	A DNSRecordType = iota
	AAAA
	CNAME
	NS
	ANYRECORDTYPE
)

type Filter struct {
	DNSStatusCode     DNSStatusCode
	DNSRecordType     DNSRecordType
	FilterWildcard    bool
	Response          bool
	ResponseOnly 		bool
	HTTPUrl           bool
	HTTPTitle         bool
	HTTPStatusCode    bool
	HTTPStatusCodeValue    int
	HTTPContentLength bool
}

func (f *Filter) isHTTPRequested() bool {
	return f.HTTPUrl || f.HTTPTitle || f.HTTPStatusCode || f.HTTPContentLength
}

func applyFilter(data *chaos.BBQData, filter *Filter) bool {
	// wildcard
	if filter.FilterWildcard && data.Wildcard {
		return false
	}
	// dns status code
	if filter.DNSStatusCode == NOERROR && data.StatusCode != "NOERROR" {
		return false
	}
	if filter.DNSStatusCode == NXDOMAIN && data.StatusCode != "NXDOMAIN" {
		return false
	}
	if filter.DNSStatusCode == SERVFAIL && data.StatusCode != "SERVFAIL" {
		return false
	}
	if filter.DNSStatusCode == REFUSED && data.StatusCode != "REFUSED" {
		return false
	}

	// dns record type
	if filter.DNSRecordType == A && len(data.A) == 0 {
		return false
	}
	if filter.DNSRecordType == AAAA && len(data.AAAA) == 0 {
		return false
	}
	if filter.DNSRecordType == CNAME && len(data.CNAME) == 0 {
		return false
	}
	if filter.DNSRecordType == NS && len(data.NS) == 0 {
		return false
	}

	// http status code
	if filter.HTTPStatusCodeValue > 0 && filter.HTTPStatusCodeValue != data.HTTPStatusCode {
		return false
	}

	return true
}

func extractOutput(data *chaos.BBQData, filter *Filter) string {
	// dns - response
	if filter.Response {
		switch filter.DNSRecordType {
		case A:
			return strings.Join(prefixWith(data.A, data.Domain), "\n")
		case AAAA:
			return strings.Join(prefixWith(data.AAAA, data.Domain), "\n")
		case CNAME:
			return strings.Join(prefixWith(data.CNAME, data.Domain), "\n")
		case NS:
			return strings.Join(prefixWith(data.NS, data.Domain), "\n")
		}
	}

	if filter.ResponseOnly {
		switch filter.DNSRecordType {
		case A:
			return strings.Join(data.A, "\n")
		case AAAA:
			return strings.Join(data.AAAA, "\n")
		case CNAME:
			return strings.Join(data.CNAME, "\n")
		case NS:
			return strings.Join(data.NS, "\n")
		}
	}

	if filter.isHTTPRequested() {
		// http - flags
		httpbuf := data.HTTPUrl
		if filter.HTTPStatusCode {
			httpbuf += fmt.Sprintf(" [%d]", data.HTTPStatusCode)
		}
		if filter.HTTPContentLength {
			httpbuf += fmt.Sprintf(" [%d]", data.HTTPContentLength)
		}
		if filter.HTTPTitle {
			httpbuf += fmt.Sprintf(" [%s]", data.HTTPTitle)
		}
		// if the url has been requested or some data added to the base one
		if  (filter.HTTPUrl || len(httpbuf) != len(data.HTTPUrl)) && len(data.HTTPUrl)>0 {
			return httpbuf
		}
		return ""
	}
	
	
	// default - print subdomain
	return fmt.Sprintf("%s.%s", data.Subdomain, data.Domain)
}

func prefixWith(s []string, prefix string) []string {
	for i,ss := range s {
		s[i] = fmt.Sprintf("%s %s", prefix, ss )
	}

	return s
}