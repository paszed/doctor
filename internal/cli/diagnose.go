package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/state"
	"github.com/paszed/doctor/internal/ui"
)

func RunDiagnose() {

	results := checks.Diagnose()

	// JSON mode
	if len(os.Args) > 2 && os.Args[2] == "--json" {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("failed to encode json:", err)
			os.Exit(1)
		}

		fmt.Println(string(data))
		return
	}

	// Human readable output
	ui.PrintResults(results)

	// Exit code calculation
	_, warn, missing := state.Summary(results)
	// Exit codes
	if missing > 0 {
		os.Exit(2)
	}

	if warn > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
