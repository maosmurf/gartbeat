package main

import (
	"os"

	"github.com/maosmurf/gartbeat/cmd"

	_ "github.com/maosmurf/gartbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
