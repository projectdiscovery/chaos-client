package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	pdcp "github.com/projectdiscovery/utils/auth/pdcp"
)

type TestCase interface {
	// Execute executes a test case and returns any errors if occurred
	Execute() error
}

var (
	customTest = os.Getenv("TEST")

	errored = false
	success = aurora.Green("[✓]").String()
	failed  = aurora.Red("[✘]").String()

	tests = map[string]map[string]TestCase{
		"code": libraryTestcases,
	}
)

func main() {
	// skip if creds are not given
	creds := pdcp.PDCPCredHandler{}
	got, err := creds.GetCreds()
	if err != nil || got == nil || got.APIKey == "" {
		fmt.Printf("Skipping integration tests as creds are not provided: %s\n", err)
		os.Exit(0)
	}

	for name, tests := range tests {
		fmt.Printf("Running test cases for \"%s\"\n", aurora.Blue(name))
		if customTest != "" && !strings.Contains(name, customTest) {
			continue // only run tests user asked
		}
		for name, test := range tests {
			err := test.Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s Test \"%s\" failed: %s\n", failed, name, err)
				errored = true
			} else {
				fmt.Printf("%s Test \"%s\" passed!\n", success, name)
			}
		}
	}
	if errored {
		os.Exit(1)
	}
}
