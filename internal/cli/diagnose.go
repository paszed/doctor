package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
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
	ok := 0
	warn := 0
	missing := 0

	for _, r := range results {

		switch r.Status {

		case model.OK:
			ok++

		case model.Warning:
			warn++

		case model.Missing:
			missing++

		}
	}

	// Exit codes
	if missing > 0 {
		os.Exit(2)
	}

	if warn > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
