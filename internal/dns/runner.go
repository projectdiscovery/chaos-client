package dns

import (
	"encoding/json"
	"fmt"

	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/gologger"
)

// Run runs the chaos client and lookup DNS
func Run(opts *Options) error {
	httpCl := internal.NewHTTPClient(opts.APIKey)
	client := NewClient(httpCl)
	resp, err := client.Lookup(opts.Domain, opts.Types)
	if err != nil {
		return fmt.Errorf("could not lookup domain: %w", err)
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}
	gologger.Silent().Msg(string(out))

	return nil
}
