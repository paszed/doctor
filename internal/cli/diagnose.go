package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/ui"
)

func RunDiagnose() {

	results := checks.Diagnose()

	if len(os.Args) > 2 && os.Args[2] == "--json" {

		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))
		return
	}

	ui.PrintResults(results)
}
