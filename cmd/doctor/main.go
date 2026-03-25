package main

import (
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/cli"
	"github.com/paszed/doctor/internal/runtime"
)

func main() {

	runtime.FixPath()
	// Register all available checks
	checks.RegisterAll()

	// Run the CLI
	cli.Run(os.Args[1:])
}
