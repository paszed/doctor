package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/state"
	"github.com/paszed/doctor/internal/ui"
)

func RunDiagnose() {
	args := os.Args

	jsonMode := false
	var selected []string

	for _, arg := range args[2:] {
		if arg == "--json" {
			jsonMode = true
		} else {
			selected = append(selected, arg)
		}
	}

	var results []model.Result

	if len(selected) == 0 {
		results = checks.RunAll()
	} else {
		for _, name := range selected {
			r, ok := checks.RunOne(name)
			if !ok {
				fmt.Printf("unknown tool: %s\n", name)
				printAvailable()
				os.Exit(1)
			}
			results = append(results, r)
		}
	}

	if jsonMode {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("failed to encode json:", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
		return
	}

	ui.PrintResults(results)

	_, warn, missing := state.Summary(results)

	if missing > 0 {
		os.Exit(2)
	}
	if warn > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}

func printAvailable() {
	fmt.Println("\navailable checks:")
	for _, name := range checks.List() {
		fmt.Printf("  %s\n", name)
	}
}
