package subdomains

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/projectdiscovery/chaos-client/internal"

	"github.com/projectdiscovery/gologger"
)

// Run runs the enumeration for Chaos client
func Run(opts *Options) error {
	httpCl := internal.NewHTTPClient(opts.APIKey)
	client := NewClient(httpCl)
	if opts.Count {
		resp, err := client.GetStatistics(&GetStatisticsRequest{
			Domain: opts.Domain,
		})
		if err != nil {
			return fmt.Errorf("could not get statistics: %w", err)
		}
		if opts.JSONOutput {
			gologger.Silent().Msgf(`{"count":%d}\n`, resp.Subdomains)
		} else {
			gologger.Silent().Msgf("%d\n", resp.Subdomains)
		}

		return nil
	}

	outputWriters := []io.Writer{os.Stdout}
	if opts.Output != "" {
		var err error
		opts.outputFile, err = os.Create(opts.Output)
		if err != nil {
			return fmt.Errorf("could not create file %s for %s: %w", opts.Output, opts.Domain, err)
		}
		defer opts.outputFile.Close()
		outputWriters = append(outputWriters, opts.outputFile)
	}
	opts.outputWriter = io.MultiWriter(outputWriters...)

	if opts.Domain != "" {
		err := processDomain(client, opts)
		if err != nil {
			return err
		}
	}

	if opts.hasStdin() || opts.DomainsFile != "" {
		err := processList(client, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

func processDomain(client *Client, opts *Options) error {
	req := &Request{Domain: opts.Domain}
	if opts.JSONOutput {
		req.OutputFormat = "json"
	}
	for item := range client.GetSubdomains(req) {
		if item.Error != nil {
			return fmt.Errorf("could not get subdomains for %s: %w", opts.Domain, item.Error)
		}
		if opts.JSONOutput {
			_, _ = io.Copy(opts.outputWriter, *item.Reader)
		} else {
			if item.Subdomain != "" {
				gologger.Silent().Msgf("%s.%s\n", item.Subdomain, opts.Domain)
			}
			if opts.Output != "" {
				_, err := fmt.Fprintf(opts.outputWriter, "%s.%s\n", item.Subdomain, opts.Domain)
				if err != nil {
					return fmt.Errorf("could not write results to file %s for %s: %w", opts.Output, opts.Domain, err)
				}
			}
		}
	}

	return nil
}

func processList(client *Client, opts *Options) error {
	var file *os.File
	var err error

	if opts.hasStdin() {
		file = os.Stdin
	}
	if opts.DomainsFile != "" {
		file, err = os.Open(opts.DomainsFile)
		if err != nil {
			return fmt.Errorf("could not open file %s: %w", opts.DomainsFile, err)
		}
	}

	in := bufio.NewScanner(file)
	for in.Scan() {
		opts.Domain = in.Text()
		err = processDomain(client, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
