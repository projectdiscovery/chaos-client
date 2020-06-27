package runner

import (
	"bufio"
	"fmt"
	"os"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	"github.com/projectdiscovery/gologger"
)

// RunEnumeration runs the enumeration for Chaos client
func RunEnumeration(opts *Options) {
	client := chaos.New(opts.APIKey)
	if opts.Update {
		if _, err := client.PutSubdomains(&chaos.PutSubdomainsRequest{Contents: os.Stdin}); err != nil {
			gologger.Fatalf("Could not upload subdomains: %s\n", err)
		}
		gologger.Infof("Input processed successfully and subdomains with valid records will be updated to chaos dataset.")
		return
	}
	if opts.UploadFilename != "" {
		file, err := os.Open(opts.UploadFilename)
		if err != nil {
			gologger.Fatalf("Could not open input file: %s\n", err)
		}
		defer file.Close()

		if _, err = client.PutSubdomains(&chaos.PutSubdomainsRequest{Contents: file}); err != nil {
			gologger.Fatalf("Could not upload subdomains: %s\n", err)
		}
		gologger.Infof("File processed successfully and subdomains with valid records will be updated to chaos dataset.")
		return
	}

	if opts.Count {
		resp, err := client.GetStatistics(&chaos.GetStatisticsRequest{
			Domain: opts.Domain,
		})
		if err != nil {
			gologger.Fatalf("Could not get statistics: %s\n", err)
		}
		gologger.Silentf("%d\n", resp.Subdomains)
		return
	}

	if opts.Output != "" {
		var err error
		opts.outputFile, err = os.Create(opts.Output)
		if err != nil {
			gologger.Fatalf("Could not create file %s for %s: %s\n", opts.Output, opts.Domain, err)
		}
		defer opts.outputFile.Close()
		opts.outputFileWriter = bufio.NewWriter(opts.outputFile)
		defer opts.outputFileWriter.Flush()
	}

	if opts.Domain != "" {
		processDomain(client, opts)
	}

	if opts.hasStdin() || opts.DomainsFile != "" {
		processList(client, opts)
	}
}

func processDomain(client *chaos.Client, opts *Options) {
	for item := range client.GetSubdomains(&chaos.GetSubdomainsRequest{Domain: opts.Domain}) {
		if item.Error != nil {
			gologger.Fatalf("Could not get subdomains: %s\n", item.Error)
		}
		if item.Subdomain != "" {
			gologger.Silentf("%s.%s\n", item.Subdomain, opts.Domain)
		}

		if opts.Output != "" {
			_, err := opts.outputFileWriter.WriteString(fmt.Sprintf("%s.%s\n", item.Subdomain, opts.Domain))
			if err != nil {
				gologger.Fatalf("Could not write results to file %s for %s: %s\n", opts.Output, opts.Domain, err)
			}
		}
	}
}

func processList(client *chaos.Client, opts *Options) {
	var (
		file *os.File
		err  error
	)
	if opts.hasStdin() {
		file = os.Stdin
	} else if opts.UploadFilename != "" {
		file, err = os.Open(opts.UploadFilename)
		if err != nil {
			gologger.Fatalf("Could not open input file: %s\n", err)
		}
		defer file.Close()
	}

	in := bufio.NewScanner(file)
	for in.Scan() {
		opts.Domain = in.Text()
		processDomain(client, opts)
	}
}
