package runner

import (
	"flag"
	"io"
	"os"

	"github.com/projectdiscovery/gologger"
)

// Options contains configuration options for chaos client.
type Options struct {
	Update            string
	APIKey            string
	Domain            string
	Count             bool
	Silent            bool
	Output            string
	DomainsFile       string
	JSONOutput        bool
	DNSStatusCode     string
	DNSRecordType     string
	FilterWildcard    bool
	Response          bool
	HTTPUrl           bool
	HTTPTitle         bool
	HTTPStatusCode    bool
	HTTPStatusCodeFilter   int
	HTTPContentLength bool
	BBQ               bool
	Version           bool
	outputFile        *os.File
	outputWriter      io.Writer
	filter            *Filter
}

// ParseOptions parses the command line options for application
func ParseOptions() *Options {
	opts := &Options{}

	flag.StringVar(&opts.APIKey, "key", "", "Chaos key for API")
	flag.StringVar(&opts.Domain, "d", "", "Domain contains domain to find subs for")
	flag.BoolVar(&opts.Count, "count", false, "Show statistics for the specified domain")
	flag.StringVar(&opts.Update, "update", "-", "Upload subdomains from stdin or filename")
	flag.BoolVar(&opts.Silent, "silent", false, "Make the output silent")
	flag.StringVar(&opts.Output, "o", "", "File to write output to (optional)")
	flag.StringVar(&opts.DomainsFile, "dL", "", "File containing subdomains to query (optional)")
	flag.BoolVar(&opts.JSONOutput, "json", false, "Print output as json")
	flag.BoolVar(&opts.BBQ, "bbq", false, "Public bugbounty recon data")
	flag.StringVar(&opts.DNSStatusCode, "dns-status-code", "", "Filter by dns status code")
	flag.StringVar(&opts.DNSRecordType, "dns-record-type", "", "Filter by dns record type")
	flag.BoolVar(&opts.FilterWildcard, "filter-wildcard", false, "Filter wildcards")
	flag.BoolVar(&opts.Response, "resp", false, "Print record response")
	flag.BoolVar(&opts.HTTPUrl, "http-url", false, "Print http url if the fqdn exposes a web server")
	flag.BoolVar(&opts.HTTPTitle, "http-title", false, "Print http homepage title if the fqdn exposes a web server")
	flag.BoolVar(&opts.HTTPStatusCode, "http-status-code", false, "Print http status code if the fqdn exposes a web server")
	flag.IntVar(&opts.HTTPStatusCodeFilter, "http-status-code-filter", -1, "Print http status code if the value equals the specified one")
	flag.BoolVar(&opts.HTTPContentLength, "http-content-length", false, "Print http content length if the fqdn exposes a web server")
	flag.BoolVar(&opts.Version, "version", false, "Show version of chaos")

	flag.Parse()

	if opts.Silent {
		gologger.MaxLevel = gologger.Silent
	}
	showBanner()

	if opts.Version {
		gologger.Infof("Current Version: %s\n", Version)
		os.Exit(0)
	}

	opts.validateOptions()

	return opts
}
func (opts *Options) validateOptions() {
	// If empty try to retrieve the key from env variables
	if opts.APIKey == "" {
		opts.APIKey = os.Getenv("CHAOS_KEY")
	}

	if opts.APIKey == "" {
		gologger.Fatalf("Authorization token not specified\n")
	}

	if opts.Update == "-" && opts.Domain == "" && opts.DomainsFile == "" && !opts.hasStdin() {
		gologger.Fatalf("No input specified for the API\n")
	}

	var filter Filter
	switch opts.DNSStatusCode {
	case "noerror":
		filter.DNSStatusCode = NOERROR
	case "nxdomain":
		filter.DNSStatusCode = NXDOMAIN
	case "servfail":
		filter.DNSStatusCode = SERVFAIL
	case "refused":
		filter.DNSStatusCode = REFUSED
	default:
		filter.DNSStatusCode = ANYDNSCODE
	}

	switch opts.DNSRecordType {
	case "a":
		filter.DNSRecordType = A
	case "aaaa":
		filter.DNSRecordType = AAAA
	case "cname":
		filter.DNSRecordType = CNAME
	case "ns":
		filter.DNSRecordType = NS
	default:
		filter.DNSRecordType = ANYRECORDTYPE
	}

	filter.FilterWildcard = opts.FilterWildcard
	filter.HTTPContentLength = opts.HTTPContentLength
	filter.HTTPStatusCode = opts.HTTPStatusCode
	filter.HTTPStatusCodeValue = opts.HTTPStatusCodeFilter
	filter.HTTPTitle = opts.HTTPTitle
	filter.Response = opts.Response

	opts.filter = &filter
}

func (opts *Options) hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}
