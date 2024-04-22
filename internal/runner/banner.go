package runner

import (
	"github.com/projectdiscovery/chaos-client/pkg/chaos"
	"github.com/projectdiscovery/gologger"
	updateutils "github.com/projectdiscovery/utils/update"
)

const banner = `
        __                    
  _____/ /_  ____ _____  _____
 / ___/ __ \/ __  / __ \/ ___/
/ /__/ / / / /_/ / /_/ (__  ) 
\___/_/ /_/\__,_/\____/____/
`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tchaos.projectdiscovery.io\n\n")
}

// GetUpdateCallback returns a callback function that updates chaos
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("chaos-client", chaos.Version)()
	}
}
