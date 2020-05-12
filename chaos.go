package main

import (
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/projectdiscovery/gologger"
)

var (
	chaosKey       = flag.String("key", "", "Chaos key for API")
	domain         = flag.String("d", "", "Domain contains domain to find subs for")
	count          = flag.Bool("count", false, "Show statistics for the specified domain")
	uploadfilename = flag.String("f", "", "File containing subdomains to upload")
	silent         = flag.Bool("silent", false, "Make the output silent")
)

var httpclient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
		MaxIdleConns:        100,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
	Timeout: time.Duration(30) * time.Second,
}

const banner = `
        __                    
  _____/ /_  ____ _____  _____
 / ___/ __ \/ __  / __ \/ ___/
/ /__/ / / / /_/ / /_/ (__  ) 
\___/_/ /_/\__,_/\____/____/  v1
`

// Version is the current version of chaos
const Version = `1.0.0`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Printf("%s\n", banner)
	gologger.Printf("\t\tprojectdiscovery.io\n\n")

	gologger.Labelf("Use with caution. You are responsible for your actions\n")
	gologger.Labelf("Developers assume no liability and are not responsible for any misuse or damage.\n")
}

func main() {
	flag.Parse()

	if *silent {
		gologger.MaxLevel = gologger.Silent
	}

	showBanner()

	// If empty try to retrieve the key from env variables
	if *chaosKey == "" {
		*chaosKey = os.Getenv("CHAOS_KEY")
	}

	if *chaosKey == "" {
		gologger.Fatalf("Authorization token not specified\n")
	}

	if *uploadfilename != "" {
		uploadFile()
		return
	}

	if *domain == "" {
		gologger.Fatalf("Domain not specified\n")
	}

	// Only domain stats
	if *count {
		getDomainStats()
		return
	}

	getSubdomains()
}

func getDomainStats() {
	req, err := http.NewRequest("GET", "https://dns.projectdiscovery.io/dns/"+*domain, nil)
	if err != nil {
		gologger.Fatalf("Could not make request: %s\n", err)
	}

	req.Header.Set("Authorization", *chaosKey)

	resp, err := httpclient.Do(req)
	if err != nil {
		gologger.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		gologger.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	var r map[string]interface{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		gologger.Fatalf("Could not unmarshal result: %s\n", err)
	}
	gologger.Silentf("%d\n", int(r["subdomains"].(float64)))
}

type result struct {
	Subdomains []string `json:"subdomains"`
}

func getSubdomains() {
	req, err := http.NewRequest("GET", "https://dns.projectdiscovery.io/dns/"+*domain+"/subdomains", nil)
	if err != nil {
		gologger.Fatalf("Could not make request: %s\n", err)
	}
	req.Header.Set("Authorization", *chaosKey)

	resp, err := httpclient.Do(req)
	if err != nil {
		gologger.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		gologger.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	r := result{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		gologger.Fatalf("Could not unmarshal result: %s\n", err)
	}

	for _, subdomain := range r.Subdomains {
		if subdomain != "" {
			gologger.Silentf("%s.%s\n", subdomain, *domain)
		}
	}
}

func uploadFile() {
	file, err := os.Open(*uploadfilename)
	if err != nil {
		gologger.Fatalf("Could not open subdomains file: %s\n", err)
	}

	req, err := http.NewRequest("POST", "https://dns.projectdiscovery.io/dns/add", file)
	if err != nil {
		gologger.Fatalf("Could not make request: %s\n", err)
	}
	req.Header.Set("Authorization", *chaosKey)

	resp, err := httpclient.Do(req)
	if err != nil {
		gologger.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		gologger.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	gologger.Infof("File processed successfully and subdomains with valid records will be updated to chaos dataset.")
}
