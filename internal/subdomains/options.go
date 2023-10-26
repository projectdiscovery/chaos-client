package subdomains

import (
	"io"
	"os"

	"github.com/projectdiscovery/gologger"
)

// Options contains configuration options for chaos client.
type Options struct {
	APIKey             string
	Domain             string
	Count              bool
	Silent             bool
	Output             string
	DomainsFile        string
	JSONOutput         bool
	Version            bool
	outputFile         *os.File
	outputWriter       io.Writer
	Verbose            bool
	DisableUpdateCheck bool
}

func (opts *Options) ValidateOptions() {
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
