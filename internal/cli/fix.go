package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/fix"
)

func RunFix() {
	args := os.Args[2:]

	if len(args) == 0 {
		fmt.Println("usage: doctor fix <tool> [args]")
		os.Exit(1)
	}

	name := args[0]
	rest := args[1:]

	// special case: ports (needs args)
	if name == "ports" {
		err := fix.FixPorts(rest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// normal fixes
	err := fix.Run(name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("\navailable fixes:")
		for _, f := range fix.List() {
			fmt.Printf("  %s\n", f)
		}
		os.Exit(1)
	}
}
