package subdomains

import (
	"fmt"
	"io"
	"os"
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

// ValidateOptions validates the options
func (opts *Options) ValidateOptions() error {
	// If empty try to retrieve the key from env variables
	if opts.APIKey == "" {
		opts.APIKey = os.Getenv("CHAOS_KEY")
	}

	if opts.APIKey == "" {
		return fmt.Errorf("Authorization token not specified")
	}

	if opts.Domain == "" && opts.DomainsFile == "" && !opts.hasStdin() {
		return fmt.Errorf("No input specified for the API")
	}

	return nil
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
