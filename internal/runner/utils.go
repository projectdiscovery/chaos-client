package runner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// WriteOutput writes the output list of subdomains to an io.Writer
func WriteOutput(results []string, domain string, writer io.Writer) error {
	bufwriter := bufio.NewWriter(writer)
	sb := &strings.Builder{}

	for _, host := range results {
		subdomain := fmt.Sprintf("%s.%s\n", host, domain)
		sb.WriteString(subdomain)
		_, err := bufwriter.WriteString(sb.String())
		if err != nil {
			bufwriter.Flush()
			return err
		}
		sb.Reset()
	}
	return bufwriter.Flush()
}
