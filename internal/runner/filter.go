package runner

type DNSStatusCode int

const (
	NOERROR DNSStatusCode = iota
	NXDOMAIN
	SERVFAIL
	REFUSED
)

type DNSRecordType int

const (
	A DNSRecordType = iota
	AAAA
	CNAME
	NS
)

type Filter struct {
	DNSStatusCode     DNSStatusCode
	DNSRecordType     DNSRecordType
	FilterWildcard    bool
	Response          bool
	HTTPUrl           bool
	HTTPTitle         bool
	HTTPStatusCode    bool
	HTTPContentLength bool
}

func applyFilter(item interface{}, filter *Filter) string {
	return ""
}
