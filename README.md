# Chaos

Go client to communicate with Chaos dataset API. 

## Installation:- 

```bash
> GO111MODULE=on go get -u github.com/projectdiscovery/chaos-client/cmd/chaos
```

## Usage:- 

```bash
> chaos -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag    | Description                              | Example                   |
|---------|------------------------------------------|---------------------------|
| -d      | Domain to find subdomains for          | chaos -d uber.com         |
| -bbq    | Bugbounty recon data query            | chaos -bbq -d uber.com         |
| -count  | Show statistics for the specified domain | chaos -d uber.com -count  |
| -o      | File to write output to (optional)       | chaos -d uber.com -o uber.txt  |
| -json | Print output as json                  | chaos -d uber.com -json |
| -update | Upload subdomains from stdin or filename     | chaos -update subdomains.txt   |
| -key    | Chaos key for API                        | chaos -key API_KEY        |
| -dL | File containing subdomains to query (optional)      | chaos -dL domains.txt |
| -dns-record-type | Filter by dns record type                   | chaos -bbq -d uber.com -dns-record-type cname |
| -dns-status-code | Filter by dns status code                   | chaos -bbq -d uber.com -dns-status-code noerror |
| -filter-wildcard | Filter DNS wildcards                   | chaos -bbq -d uber.com -filter-wildcard |
| -http-url | Print URL of the subdomains                  | chaos -bbq -d uber.com -http-url |
| -http-title | Print title of the URL                   | chaos -bbq -d uber.com -http-title |
| -http-status-code | Print http status code                   | chaos -bbq -d uber.com -http-status-code |
| -http-status-code-filter | Filter http status code                   | chaos -bbq -d uber.com -http-status-code -http-status-code-filter 200 |
| -resp | Print DNS record with response                 | chaos -bbq -d uber.com -resp |
| -resp-only | Print the response of DNS record                   | chaos -bbq -d uber.com -dns-record-type cname -resp-only |
| -silent | Make the output silent                   | chaos -d uber.com -silent |
| -version | Print current version of chaos client                  | chaos -version |


You can also set the API key as environment variable in your bash profile. 

```bash
export CHAOS_KEY="CHAOS_API_KEY"
```

### How to avail `API_KEY`

As of now Chaos dataset is in beta for testing and API endpoint access available to invited users only, you can request an invite for yourself [here](https://forms.gle/LkHUjoxAiHE6djtU6), we are sending out invites in FIFO manner, so we have no ETA.  

# Running chaos

In order to get subdomains for a domain, use the following command.

```bash
> chaos -d uber.com -silent 
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

NOTE:- 

1. **Chaos dataset endpoint supports "domain" name as input, "string" or "subdomain" based searches are not supported.**   
2. **Chaos recon data can be retrieved with `bbq` flag only.**   

In order to get URLs for a domain, use the following command.

```bash
> chaos -d uber.com -silent -bbq -http-url
https://free.uber.com
https://jobs.uber.com
https://join.uber.com
https://amp.uber.com
https://restaurant-onboarding-staging.uber.com
https://tc.uber.com
https://wallet.uber.com
https://brand.uber.com
https://autor.uber.com
https://vouchers.uber.com
https://survey.uber.com
https://drive.uber.com
https://spotlight.uber.com
https://cn-sjc1.uber.com
https://patagonia.uber.com
https://cn-sjc1.cfe.uber.com
https://lite.uber.com
https://freight.uber.com
https://ar.uber.com
https://freightbonjour.uber.com
https://azkaban.uber.com
https://voice.uber.com
https://messages-staging.uber.com
```


To get the number of subdomains without getting actual results, you can use the `count` flag.

```bash
> chaos -d uber.com -count -silent 
10640320
```

Additional subdomains can also be uploaded to the Chaos dataset using the `update` flag. Uploads are limited to 10 MB as of now. The uploaded data will be added to the public dataset and is completely voluntary. 

NOTE:- 

**Only subdomains with valid record gets added to dataset, subdomains with dead records gets eliminated**   


```bash
> cat subs.txt | chaos -update

        __                    
  _____/ /_  ____ _____  _____
 / ___/ __ \/ __  / __ \/ ___/
/ /__/ / / / /_/ / /_/ (__  ) 
\___/_/ /_/\__,_/\____/____/  v1

		projectdiscovery.io

[WRN] Use with caution. You are responsible for your actions
[WRN] Developers assume no liability and are not responsible for any misuse or damage.
[INF] Input processed successfully and subdomains with valid records will be updated to chaos dataset.
```
Subfinder also supports updating data to chaos dataset and can be queried later on the go. 

```bash
> cat domains.txt | subfinder -cd

        __                    
  _____/ /_  ____ _____  _____
 / ___/ __ \/ __  / __ \/ ___/
/ /__/ / / / /_/ / /_/ (__  ) 
\___/_/ /_/\__,_/\____/____/  v1

		projectdiscovery.io

[WRN] Use with caution. You are responsible for your actions
[WRN] Developers assume no liability and are not responsible for any misuse or damage.
[INF] Input processed successfully and subdomains with valid records will be updated to chaos dataset.
```

NOTE: 

**The API is rate-limited to 1 request at time per token (you can issue the next request only when the previous one is finished).**
