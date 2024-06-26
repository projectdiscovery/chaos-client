<h1 align="center">
Chaos Client
</h1>
<h4 align="center">Go client to communicate with Chaos dataset API.</h4>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases"><img src="https://img.shields.io/github/downloads/projectdiscovery/chaos-client/total">
<a href="https://github.com/projectdiscovery/chaos-client/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/projectdiscovery/chaos-client">
<a href="https://github.com/projectdiscovery/chaos-client/releases/"><img src="https://img.shields.io/github/release/projectdiscovery/chaos-client">
<a href="https://discord.gg/projectdiscovery"><img src="https://img.shields.io/discord/695645237418131507.svg?logo=discord"></a>
<a href="https://twitter.com/pdchaos"><img src="https://img.shields.io/twitter/follow/pdchaos.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="https://github.com/projectdiscovery/chaos-client/blob/main/README.md">English</a> •
  <a href="https://github.com/projectdiscovery/chaos-client/blob/main/README_CN.md">中文</a> 
</p>

## Installation

```bash
go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
```

## Usage

```bash
chaos -h
```

This will display help for the tool. Here are all the switches it supports.

```console
   -key string                  projectdiscovery cloud (pdcp) api key
   -d string                    domain to search for subdomains
   -count                       show statistics for the specified domain
   -silent                      make the output silent
   -o string                    file to write output to (optional)
   -dL string                   file containing domains to search for subdomains (optional)
   -json                        print output as json
   -version                     show version of chaos
   -v, -verbose                 verbose mode
   -up, -update                 update chaos to latest version
   -duc, -disable-update-check  disable automatic chaos update check
```

You can also set the API key as an environment variable in your bash profile. 

```bash
export PDCP_API_KEY=xxxxx
```

### How to avail `PDCP_API_KEY`

You can get your API key by either signing up or logging in at [cloud.projectdiscovery.io](https://cloud.projectdiscovery.io?ref=api_key).

## Running chaos

In order to get subdomains for a domain, use the following command.

```bash
chaos -d uber.com -silent

restaurants.uber.com
testcdn.uber.com
approvalservice.uber.com
zoom-logs.uber.com
eastwood.uber.com
meh.uber.com
webview.uber.com
kiosk-api.uber.com
utmbeta-staging.uber.com
getmatched-staging.uber.com
logs.uber.com
dca1.cfe.uber.com
cn-staging.uber.com
frontends-primary.uber.com
eng.uber.com
guest.uber.com
kiosk-home-staging.uber.com
```

💡 Notes
-----

- **The API is rate-limited to 60 request / min / ip**
- Chaos API **only** supports domain name to query.

## Chaos as a library
`Chaos` can be utilized as a library for subdomain enumeration by instantiating the `Options` struct and populating it with the same options that would be specified via CLI.

### Example
```go
package main

import (
	"os"
	"github.com/projectdiscovery/chaos-client/internal/runner"
	"github.com/projectdiscovery/chaos-client/pkg/chaos"
)

func main() {
	var results []chaos.Result
	opts := &runner.Options{
		Domain: "projectdiscovery.io",
		APIKey: os.Getenv("PDCP_API_KEY"),
		OnResult: func(result interface{}) {
			if val, ok := result.(chaos.Result); !ok {
				results = append(results, val)
			}
		},
	}

	runner.RunEnumeration(opts)
}

```
💡 Note

To run the program, you need to set the `PDCP_API_KEY` environment variable to your Chaos API key.

👨‍💻 Community
-----

You are welcomed to join our [Discord Community](https://discord.gg/projectdiscovery). You can also follow us on [Twitter](https://twitter.com/pdchaos) to keep up with everything related to chaos project.


Thanks again for your contribution and keeping the community vibrant. :heart:
