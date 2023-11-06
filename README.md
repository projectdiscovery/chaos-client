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

## Installation

```bash
go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
```

## General usage

```bash
chaos -h
```

This will display help for the tool. Here are all the switches it supports.

```bash
chaos <command>
```

| Flag                     | Description                           | Example                        |
|--------------------------|---------------------------------------|--------------------------------|
| `-h`,`--help`            | Show help                             | `chaos -h`                     |
| `--key`                  | Chaos key for API                     | `chaos <command> -key API_KEY` |
| `--silent`               | Make the output silent                | `chaos <command> -silent`      |
| `-o`,`--output`          | File to write output to (optional)    | `chaos <command> -o uber.txt`  |
| `-version`               | Print current version of chaos client | `chaos -version`               |
| `-v`,`--verbose`         | Show verbose output                   | `chaos -verbose`               |
| `--update`               | updates to latest version             | `chaos -up`                    | 
| `--disable-update-check` | disables automatic update check       | `chaos -duc`                   |

### Subdomains

Get subdomains for a domain.

```bash
chaos subdomains <domain>
```

Take a domain name and return subdomains for it.

| Flag      | Description                              | Example                               |
|-----------|------------------------------------------|---------------------------------------|
| `--count` | Show statistics for the specified domain | `chaos subdomains google.com --count` |
| `--json`  | Print output as json                     | `chaos subdomains google.com --json`  |

Example:

```bash
chaos subdomains uber.com -silent

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

### Batch subdomains

Get subdomains from list of domains.

```bash
chaos subdomains-batch <path-to-files>
```

| Flag     | Description          | Example                                  |
|----------|----------------------|------------------------------------------|
| `--json` | Print output as json | `chaos subdomains-batch subs.txt --json` |

Take domains from a file and return subdomains for each domain.

### DNS Records

Lookup DNS records for a domain.

```bash
chaos dns <domain>
```

| Flag                      | Description                                             | Example                             |
|---------------------------|---------------------------------------------------------|-------------------------------------|
| `-t`, `--types=TYPES,...` | DNS record type(s) (a,aaaa,caa,cname,mx,ns,soa,srv,txt) | `chaos dns hackerone.com -t a,aaaa` |

Example:

```bash
chaos dns uber.com -t a,soa | jq

{
  "a": [
    "34.98.127.226"
  ],
  "soa": [
    "edns126.ultradns.com",
    "edns126.ultradns.com",
    "edns126.ultradns.com",
    "edns126.ultradns.com",
    "edns126.ultradns.com"
  ]
}
```

### How to avail `API_KEY`

Chaos DNS API is in beta and only available to people who have been invited to use it. You can request an invite
at [chaos.projectdiscovery.io](https://chaos.projectdiscovery.io).

You can also set the API key as an environment variable in your bash profile.

```bash
export CHAOS_KEY=CHAOS_API_KEY
```

💡 Notes
-----

- **The API is rate-limited to 60 request / min / ip**
- Chaos API **only** supports domain name to query.

👨‍💻 Community
-----

You are welcomed to join our [Discord Community](https://discord.gg/projectdiscovery). You can also follow us
on [Twitter](https://twitter.com/pdchaos) to keep up with everything related to chaos project.

Thanks again for your contribution and keeping the community vibrant. :heart:
