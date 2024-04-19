package runner

import (
	"io"
	"os"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	updateutils "github.com/projectdiscovery/utils/update"
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
	Verbose              bool
	DisableUpdateCheck   bool
	OnResult             func(result interface{})
}

// ParseOptions parses the command line options for application
func ParseOptions() *Options {
	opts := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.Marshal = true
	flagSet.StringVar(&opts.APIKey, "key", "", "Chaos key for API")
	flagSet.StringVar(&opts.Domain, "d", "", "Domain to search for subdomains")
	flagSet.BoolVar(&opts.Count, "count", false, "Show statistics for the specified domain")
	flagSet.BoolVar(&opts.Silent, "silent", false, "Make the output silent")
	flagSet.StringVar(&opts.Output, "o", "", "File to write output to (optional)")
	flagSet.StringVar(&opts.DomainsFile, "dL", "", "File containing domains to search for subdomains (optional)")
	flagSet.BoolVar(&opts.JSONOutput, "json", false, "Print output as json")
	flagSet.BoolVar(&opts.Version, "version", false, "Show version of chaos")
	flagSet.BoolVarP(&opts.Verbose, "verbose", "v", false, "Verbose")
	flagSet.CallbackVarP(GetUpdateCallback(), "update", "up", "update chaos to latest version")
	flagSet.BoolVarP(&opts.DisableUpdateCheck, "disable-update-check", "duc", false, "disable automatic chaos update check")

	_ = flagSet.Parse()

	if opts.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
	showBanner()

	if opts.Version {
		gologger.Info().Msgf("Current Version: %s\n", chaos.Version)
		os.Exit(0)
	}

	if !opts.DisableUpdateCheck {
		latestVersion, err := updateutils.GetToolVersionCallback("chaos-client", chaos.Version)()
		if err != nil {
			if opts.Verbose {
				gologger.Error().Msgf("chaos version check failed: %v", err.Error())
			}
		} else {
			gologger.Info().Msgf("Current chaos version %v %v", chaos.Version, updateutils.GetVersionDescription(chaos.Version, latestVersion))
		}
	}

	// is this sdk
	chaos.IsSDK = false

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
