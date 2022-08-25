package runner

import (
	"flag"
	"io"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

// Options contains configuration options for chaos client.
type Options struct {
	APIKey               string
	Domain               string
	Count                bool
	Silent               bool
	Output               string
	DomainsFile          string
	JSONOutput           bool
	DNSStatusCode        string
	DNSRecordType        string
	FilterWildcard       bool
	Response             bool
	ResponseOnly         bool
	HTTPUrl              bool
	HTTPTitle            bool
	HTTPStatusCode       bool
	HTTPStatusCodeFilter int
	HTTPContentLength    bool
	BBQ                  bool
	Version              bool
	outputFile           *os.File
	outputWriter         io.Writer
	filter               *Filter
}

// ParseOptions parses the command line options for application
func ParseOptions() *Options {
	opts := &Options{}

	flag.StringVar(&opts.APIKey, "key", "", "Chaos key for API")
	flag.StringVar(&opts.Domain, "d", "", "Domain contains domain to find subs for")
	flag.BoolVar(&opts.Count, "count", false, "Show statistics for the specified domain")
	flag.BoolVar(&opts.Silent, "silent", false, "Make the output silent")
	flag.StringVar(&opts.Output, "o", "", "File to write output to (optional)")
	flag.StringVar(&opts.DomainsFile, "dL", "", "File containing subdomains to query (optional)")
	flag.BoolVar(&opts.JSONOutput, "json", false, "Print output as json")
	flag.BoolVar(&opts.Version, "version", false, "Show version of chaos")

	flag.Parse()

	if opts.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
	showBanner()

	if opts.Version {
		gologger.Info().Msgf("Current Version: %s\n", Version)
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
		gologger.Fatal().Msgf("Authorization token not specified\n")
	}

	if opts.Domain == "" && opts.DomainsFile == "" && !opts.hasStdin() {
		gologger.Fatal().Msgf("No input specified for the API\n")
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
	filter.HTTPUrl = opts.HTTPUrl
	filter.Response = opts.Response
	filter.ResponseOnly = opts.ResponseOnly

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
