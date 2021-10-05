package main

import (
	"os"
	"runtime"

	"blacklemur.com/datanerd/pkg/cli"
	"blacklemur.com/datanerd/pkg/config/system"

	"github.com/rs/zerolog/log"
)

// Main function
func main() {
	log.Debug().Msgf("datanerd version: %s", system.GetVersion())
	log.Debug().Msgf("go runtime version: %s", runtime.Version())
	log.Debug().Msgf("CLI args: %#v", os.Args)
	cli.Execute()
}
