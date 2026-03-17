package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunCheck() {

	if len(os.Args) < 3 {
		fmt.Println("usage: doctor check <tool>")
		os.Exit(1)
	}

	name := os.Args[2]

	result, found := checks.RunOne(name)

	if !found {

		fmt.Printf("unknown tool: %s\n\n", name)

		fmt.Println("available checks:")

		for _, n := range checks.Names() {
			fmt.Printf("  %s\n", n)
		}

		os.Exit(1)
	}
	// handle flags
	if len(os.Args) > 3 {

		flag := os.Args[3]

		switch flag {

		case "--json":
			data, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				fmt.Println("json error:", err)
				os.Exit(1)
			}

			fmt.Println(string(data))
			return

		case "--version":
			fmt.Println(result.Message)
			return

		case "--path":
			fmt.Println(result.Path)
			return

		}
	}

	ui.PrintResults([]model.Result{result})
}
