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

## Usage

```bash
chaos -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag                       | Description                              | Example                                                    |
|----------------------------|------------------------------------------|------------------------------------------------------------|
| `-d`                       | Domain to find subdomains for            | `chaos -d uber.com`                                        |
| `-count`                   | Show statistics for the specified domain | `chaos -d uber.com -count`                                 |
| `-o`                       | File to write output to (optional)       | `chaos -d uber.com -o uber.txt`                            |
| `-json`                    | Print output as json                     | `chaos -d uber.com -json`                                  |
| `-key`                     | Chaos key for API                        | `chaos -key API_KEY`                                       |
| `-dL`                      | File with list of domains (optional)     | `chaos -dL domains.txt`                                    |
| `-silent`                  | Make the output silent                   | `chaos -d uber.com -silent`                                |
| `-version`                 | Print current version of chaos client    | `chaos -version`                                           |
| `-verbose`                 | Show verbose output                      | `chaos -verbose`                                           |
| `-update`                  | updates to latest version                | `chaos -up`                                                | 
| `-disable-update-check`    | disables automatic update check          | `chaos -duc`                                               |

You can also set the API key as environment variable in your bash profile. 

```bash
export CHAOS_KEY=CHAOS_API_KEY
```

### How to avail `API_KEY`

Chaos DNS API is in beta and only available to people who have been invited to use it. You can request an invite for yourself by filling [Google form](https://forms.gle/GP5nTamxJPfiMaBn9). We send out new invites every second Monday of the month. Please send us a DM on our [Discord](https://discord.gg/projectdiscovery) server.

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

👨‍💻 Community
-----

You are welcomed to join our [Discord Community](https://discord.gg/projectdiscovery). You can also follow us on [Twitter](https://twitter.com/pdchaos) to keep up with everything related to chaos project.


Thanks again for your contribution and keeping the community vibrant. :heart:
