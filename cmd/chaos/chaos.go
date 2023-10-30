package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/projectdiscovery/chaos-client/internal"
	"github.com/projectdiscovery/chaos-client/internal/subdomains"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	updateutils "github.com/projectdiscovery/utils/update"
	"os"
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
		DL     string `help:"File containing domains to search for subdomains (optional)" aliases:"dL"`
		JSON   bool   `help:"Print output as json"`
	} `cmd:"" help:"List subdomains"`

	DNS struct {
		Paths []string `arg:"" optional:"" help:"Paths to list." type:"path"`
	} `cmd:"" help:"Get DNS record"`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("Chaos"),
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
			APIKey:             cli.Key,
			Domain:             cli.Subdomains.Domain,
			Count:              cli.Subdomains.Count,
			Silent:             cli.Silent,
			Output:             cli.Output,
			DomainsFile:        cli.Subdomains.DL,
			JSONOutput:         cli.Subdomains.JSON,
			Version:            cli.Version,
			Verbose:            cli.Verbose,
			DisableUpdateCheck: cli.DisableUpdateCheck,
		}
		opts.ValidateOptions()
		subdomains.RunEnumeration(&opts)
	default:
		fmt.Printf("unexpected command")
		os.Exit(1)
	}
}
