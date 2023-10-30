package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/chaos-client/internal/dns"
	"github.com/projectdiscovery/chaos-client/internal/subdomains"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	updateutils "github.com/projectdiscovery/utils/update"
)

var cli struct {
	Key                string `help:"Chaos key for API" short:"k"`
	Silent             bool   `help:"Make the output silent"`
	Output             string `help:"File to write output to (optional)" short:"o"`
	Version            bool   `help:"Show version of chaos"`
	Verbose            bool   `help:"Verbose" short:"v"`
	Update             bool   `help:"update Chaos to latest version" aliases:"up"`
	DisableUpdateCheck bool   `help:"disable automatic Chaos update check" aliases:"duc"`
	Subdomains         struct {
		Domain string `arg:"" help:"Domain to search for subdomains"`
		Count  bool   `help:"Show statistics for the specified domain"`
		JSON   bool   `help:"Print output as json"`
	} `cmd:"" help:"List subdomains"`
	SubdomainsBatch struct {
		File string `arg:"" help:"File containing domains to search for subdomains"`
		JSON bool   `help:"Print output as json"`
	} `cmd:"" help:"List subdomains from file"`
	DNS struct {
		Domain string           `arg:"" help:"Domain to search for DNS record"`
		Types  []dns.RecordType `help:"DNS record type(s) (a,aaaa,caa,cname,mx,ns,soa,srv,txt)" short:"t"`
	} `cmd:"" help:"Get DNS record"`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("chaos"),
		kong.Description("Chaos client"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	if cli.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
	internal.ShowBanner()

	if cli.Version {
		gologger.Info().Msgf("Current Version: %s\n", internal.Version)
		os.Exit(0)
	}

	if !cli.DisableUpdateCheck {
		latestVersion, err := updateutils.GetVersionCheckCallback("chaos-client")()
		if err != nil {
			if cli.Verbose {
				gologger.Error().Msgf("chaos version check failed: %v", err.Error())
			}
		} else {
			gologger.Info().Msgf("Current chaos version %v %v", internal.Version, updateutils.GetVersionDescription(internal.Version, latestVersion))
		}
	}

	switch ctx.Command() {
	case "subdomains <domain>":
		opts := subdomains.Options{
			APIKey:     cli.Key,
			Count:      cli.Subdomains.Count,
			Silent:     cli.Silent,
			Output:     cli.Output,
			JSONOutput: cli.Subdomains.JSON,
			Version:    cli.Version,
			Verbose:    cli.Verbose,
		}
		err := opts.ValidateOptions()
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		err = subdomains.Run(&opts)
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	case "subdomains-batch <file>":
		opts := subdomains.Options{
			APIKey:      cli.Key,
			DomainsFile: cli.SubdomainsBatch.File,
			Silent:      cli.Silent,
			Output:      cli.Output,
			JSONOutput:  cli.SubdomainsBatch.JSON,
			Version:     cli.Version,
			Verbose:     cli.Verbose,
		}
		err := opts.ValidateOptions()
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
		err = subdomains.Run(&opts)
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	case "dns <domain>":
		opts := dns.Options{
			APIKey: cli.Key,
			Domain: cli.DNS.Domain,
			Types:  cli.DNS.Types,
		}
		err := opts.ValidateOptions()
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}

		err = dns.Run(&opts)
		if err != nil {
			gologger.Fatal().Msg(err.Error())
		}
	default:
		gologger.Fatal().Msgf("unexpected command: %s", ctx.Command())
	}
}
