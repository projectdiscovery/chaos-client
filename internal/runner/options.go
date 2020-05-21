package runner

import (
	"flag"
	"os"

	"github.com/projectdiscovery/gologger"
)

// Options contains configuration options for chaos client.
type Options struct {
	Update         bool
	APIKey         string
	Domain         string
	Count          bool
	UploadFilename string
	Silent         bool
	Output         string
}

// ParseOptions parses the command line options for application
func ParseOptions() *Options {
	opts := &Options{}

	flag.StringVar(&opts.APIKey, "key", "", "Chaos key for API")
	flag.StringVar(&opts.Domain, "d", "", "Domain contains domain to find subs for")
	flag.BoolVar(&opts.Count, "count", false, "Show statistics for the specified domain")
	flag.StringVar(&opts.UploadFilename, "f", "", "File containing subdomains to upload")
	flag.BoolVar(&opts.Update, "update", false, "Upload subdomains from stdin")
	flag.BoolVar(&opts.Silent, "silent", false, "Make the output silent")
	flag.StringVar(&opts.Output, "o", "", "File to write output to (optional)")
	flag.Parse()

	if opts.Silent {
		gologger.MaxLevel = gologger.Silent
	}
	showBanner()

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

	if !opts.Update && opts.UploadFilename == "" && opts.Domain == "" {
		gologger.Fatalf("No input specified for the API\n")
	}
}
