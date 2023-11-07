package main

import "github.com/projectdiscovery/chaos-client/internal/runner"

func main() {
	opts := runner.ParseOptions()
	runner.RunEnumeration(opts)
}
