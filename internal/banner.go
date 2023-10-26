package internal

import (
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

// Version is the current Version of chaos
const Version = `0.5.1`

// ShowBanner is used to show the banner to the user
func ShowBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tchaos.projectdiscovery.io\n\n")
}

// GetUpdateCallback returns a callback function that updates chaos
func GetUpdateCallback() func() {
	return func() {
		ShowBanner()
		updateutils.GetUpdateToolCallback("chaos-client", Version)()
	}
}
