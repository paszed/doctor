package main

import (
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/cli"
)

func main() {
	// Register all available checks
	checks.RegisterAll()

	// Run the CLI
	cli.Run(os.Args[1:])
}
