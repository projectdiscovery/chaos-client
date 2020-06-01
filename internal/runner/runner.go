package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	"github.com/projectdiscovery/gologger"
)

// RunEnumeration runs the enumeration for Chaos client
func RunEnumeration(opts *Options) {
	client := chaos.New(opts.APIKey)
	if opts.Update {
		var buf = &bytes.Buffer{}
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			buf.Write(in.Bytes())
			buf.WriteString("\n")
		}
		_, err := client.PutSubdomains(&chaos.PutSubdomainsRequest{
			Contents: buf,
		})
		if err != nil {
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

		_, err = client.PutSubdomains(&chaos.PutSubdomainsRequest{
			Contents: file,
		})
		if err != nil {
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

	var (
		file      *os.File
		bufwriter *bufio.Writer
	)

	if opts.Output != "" {
		var err error
		file, err = os.Create(opts.Output)
		if err != nil {
			gologger.Fatalf("Could not create file %s for %s: %s\n", opts.Output, opts.Domain, err)
		}
		bufwriter = bufio.NewWriter(file)
	}

	for item := range client.GetSubdomains(&chaos.GetSubdomainsRequest{Domain: opts.Domain}) {
		if item.Error != nil {
			gologger.Fatalf("Could not get subdomains: %s\n", item.Error)
		}
		if item.Subdomain != "" {
			gologger.Silentf("%s.%s\n", item.Subdomain, opts.Domain)
		}

		if opts.Output != "" {
			_, err := bufwriter.WriteString(fmt.Sprintf("%s.%s\n", item.Subdomain, opts.Domain))
			if err != nil {
				gologger.Fatalf("Could not write results to file %s for %s: %s\n", opts.Output, opts.Domain, err)
			}
		}
	}

	if file != nil {
		bufwriter.Flush()
		file.Close()
	}

	return
}
