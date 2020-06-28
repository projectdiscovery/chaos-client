package runner

import "github.com/projectdiscovery/chaos-client/pkg/chaos"

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
	HTTPUrl           bool
	HTTPTitle         bool
	HTTPStatusCode    int
	HTTPContentLength bool
}

func applyFilter(data *chaos.BBQData, filter *Filter) bool {
	// wildcard
	if filter.FilterWildcard && data.Wildcard {
		return false
	}

	// dns status code
	if filter.DNSStatusCode == NOERROR && data.StatusCode != "noerror" {
		return false
	}
	if filter.DNSStatusCode == NXDOMAIN && data.StatusCode != "nxdomain" {
		return false
	}
	if filter.DNSStatusCode == SERVFAIL && data.StatusCode != "servfail" {
		return false
	}
	if filter.DNSStatusCode == REFUSED && data.StatusCode != "refused" {
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
	if filter.HTTPStatusCode > 0 && filter.HTTPStatusCode != data.HTTPStatusCode {
		return false
	}

	return true
}
