package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	authorizationToken = flag.String("token", "", "Authorization token for API")
	domain             = flag.String("domain", "", "Domain contains domain to find subs for")
	statsTotal         = flag.Bool("stats-total", false, "Show total statistics")
	stats              = flag.Bool("stats", false, "Only show statistics for the specified domain")
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

func main() {
	flag.Parse()

	// Only total stats
	if *statsTotal {
		getTotalStats()
		return
	}

	if *domain == "" {
		log.Fatal("Domain not specified")
	}
	if *authorizationToken == "" {
		log.Fatal("Authorization token not specified")
	}

	// Only domain stats
	if *stats {
		getDomainStats()
		return
	}

	getSubdomains()
}

func getTotalStats() {
	req, err := http.NewRequest("GET", "https://dns.projectdiscovery.io/dns/stats", nil)
	if err != nil {
		log.Fatalf("Could not make request: %s\n", err)
	}
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	var r map[string]interface{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalf("Could not unmarshal result: %s\n", err)
	}

	fmt.Printf("Total: %.0f\n", r["total"])
	fmt.Printf("New Last Week: %.0f\n", r["new_last_week"])
	fmt.Printf("New Last 24 Hours: %.0f\n", r["new_last_24hour"])
	fmt.Printf("New Last Hour: %.0f\n", r["new_last_hour"])
}

func getDomainStats() {
	req, err := http.NewRequest("GET", "https://dns.projectdiscovery.io/dns/"+*domain, nil)
	if err != nil {
		log.Fatalf("Could not make request: %s\n", err)
	}

	req.Header.Set("Authorization", *authorizationToken)

	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	var r map[string]interface{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalf("Could not unmarshal result: %s\n", err)
	}

	fmt.Println("Subdomains", r["subdomains"])
}

type result struct {
	Subdomains []string `json:"subdomains"`
}

func getSubdomains() {
	req, err := http.NewRequest("GET", "https://dns.projectdiscovery.io/dns/"+*domain+"/subdomains", nil)
	if err != nil {
		log.Fatalf("Could not make request: %s\n", err)
	}
	req.Header.Set("Authorization", *authorizationToken)

	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatalf("Could not send request: %s\n", err)
	}

	if resp.StatusCode != 200 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Fatalf("Could not finish request: %d statuscode\n", resp.StatusCode)
	}

	r := result{}
	err = jsoniter.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalf("Could not unmarshal result: %s\n", err)
	}

	for _, subdomain := range r.Subdomains {
		if subdomain != "" {
			fmt.Println(subdomain + "." + *domain)
		}
	}
}
