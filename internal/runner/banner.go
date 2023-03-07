package runner

import "github.com/projectdiscovery/gologger"

const banner = `
        __                    
  _____/ /_  ____ _____  _____
 / ___/ __ \/ __  / __ \/ ___/
/ /__/ / / / /_/ / /_/ (__  ) 
\___/_/ /_/\__,_/\____/____/  v0.5.0
`

// Version is the current version of chaos
const Version = `0.5.0`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tchaos.projectdiscovery.io\n\n")
}
