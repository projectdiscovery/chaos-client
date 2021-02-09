[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/projectdiscovery/chaos-client)](https://goreportcard.com/report/github.com/projectdiscovery/chaos-client)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/projectdiscovery/chaos-client/issues)
[![GitHub Release](https://img.shields.io/github/release/projectdiscovery/chaos-client)](https://github.com/projectdiscovery/chaos-client/releases)
[![Follow on Twitter](https://img.shields.io/twitter/follow/pdchaos.svg?logo=twitter)](https://twitter.com/pdchaos)
[![Chat on Discord](https://img.shields.io/discord/695645237418131507.svg?logo=discord)](https://discord.gg/KECAGdH)

# Chaos

Go client to communicate with Chaos dataset API. 

## Installation:- 

```bash
▶ GO111MODULE=on go get -v github.com/projectdiscovery/chaos-client/cmd/chaos
```

## Usage:- 

```bash
▶ chaos -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag                     | Description                              | Example                                                  |
| ------------------------ | ---------------------------------------- | -------------------------------------------------------- |
| -d                       | Domain to find subdomains for            | chaos -d uber.com                                        |
| -count                   | Show statistics for the specified domain | chaos -d uber.com -count                                 |
| -o                       | File to write output to (optional)       | chaos -d uber.com -o uber.txt                            |
| -json                    | Print output as json                     | chaos -d uber.com -json                                  |
| -key                     | Chaos key for API                        | chaos -key API_KEY                                       |
| -dL                      | File with list of domains (optional)     | chaos -dL domains.txt                                    |
| -dns-record-type         | Filter by dns record type                | chaos -bbq -d uber.com -dns-record-type cname            |
| -dns-status-code         | Filter by dns status code                | chaos -bbq -d uber.com -dns-status-code noerror          |
| -filter-wildcard         | Filter DNS wildcards                     | chaos -bbq -d uber.com -filter-wildcard                  |
| -http-url                | Print URL of the subdomains              | chaos -bbq -d uber.com -http-url                         |
| -http-title              | Print title of the URL                   | chaos -bbq -d uber.com -http-title                       |
| -http-status-code        | Print http status code                   | chaos -bbq -d uber.com -http-status-code                 |
| -http-status-code-filter | Filter http status code                  | chaos -bbq -d uber.com -http-status-code-filter 200      |
| -resp                    | Print DNS record with response           | chaos -bbq -d uber.com -resp                             |
| -resp-only               | Print the response of DNS record         | chaos -bbq -d uber.com -dns-record-type cname -resp-only |
| -silent                  | Make the output silent                   | chaos -d uber.com -silent                                |
| -version                 | Print current version of chaos client    | chaos -version                                           |


You can also set the API key as environment variable in your bash profile. 

```bash
export CHAOS_KEY=CHAOS_API_KEY
```

### How to avail `API_KEY`

As of now Chaos dataset is in beta for testing and API endpoint access available to invited users only, you can request an invite for yourself [here](https://forms.gle/GP5nTamxJPfiMaBn9), we are sending out new invites every 2nd monday of the month, singed up and still missing the invites? feel free to shoot us a DM is our [Discord](https://discord.gg/KECAGdH) server.

# Running chaos

In order to get subdomains for a domain, use the following command.

```bash
▶ chaos -d uber.com -silent

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

To get the number of subdomains count, you can use the `count` flag.

```bash
▶ chaos -d uber.com -count -silent

7685
```


👨‍💻 Community
-----

You are welcomed to join our [Discord Community](https://discord.gg/KECAGdH). You can also follow us on [Twitter](https://twitter.com/pdchaos) to keep up with everything related to chaos project.

💡 Notes
-----

- The API is rate-limited to 1 request at a time per token.
- Chaos API **only** supports domain name to query.
- Chaos recon data can be retrieved using `bbq` flag.

📌 Reference
-----

- [Chaos Recon Data](https://blog.projectdiscovery.io/post/chaos-recon-data/)


Thanks again for your contribution and keeping the community vibrant. :heart:
