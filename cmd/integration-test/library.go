package main

import (
	"os"

	"github.com/projectdiscovery/chaos-client/internal/runner"
	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	errorutil "github.com/projectdiscovery/utils/errors"
)

var libraryTestcases = map[string]TestCase{
	"chaos-client as library": &goIntegrationTest{},
}

type goIntegrationTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *goIntegrationTest) Execute() error {
	var got []string
	opts := &runner.Options{
		Domain: "projectdiscovery.io",
		APIKey: os.Getenv("CHAOS_KEY"),
		OnResult: func(result interface{}) {
			if val, ok := result.(chaos.Result); !ok {
				got = append(got, val.Subdomain)
			}
		},
	}
	runner.RunEnumeration(opts)
	if len(got) < 1 {
		return errorutil.New("Failed")
	}
	return nil
}

