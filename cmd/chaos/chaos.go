package main

import "github.com/projectdiscovery/chaos-client/chaos/internal/runner"

func main() {
	opts := runner.ParseOptions()
	runner.RunEnumeration(opts)
}
