package main

import (
	"os"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	"github.com/projectdiscovery/utils/errkit"
)

var libraryTestcases = map[string]TestCase{
	"chaos-client as library": &goIntegrationTest{},
}

type goIntegrationTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *goIntegrationTest) Execute() error {
	var got []string
	client := chaos.New(os.Getenv("CHAOS_KEY"))
	for result := range client.GetSubdomains(&chaos.SubdomainsRequest{
		Domain: "projectdiscovery.io",
	}) {
		if result.Error != nil {
			return result.Error
		}
		if result.Subdomain != "" {
			got = append(got, result.Subdomain)
		}
	}
	if len(got) < 1 {
		return errkit.New("failed")
	}
	return nil
}

