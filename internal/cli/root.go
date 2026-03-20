package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/config"
	"github.com/paszed/doctor/internal/state"
	"github.com/paszed/doctor/internal/ui"
)

func Run(args []string) {

	config.Current = config.Load()
	if len(args) == 0 {
		RunInteractive()
		return
	}

	switch args[0] {
	case "help", "-h", "--help":
		RunHelp(args[1:])
	case "diagnose":
		RunDiagnose(args[1:])
	case "check":
		RunCheck(args[1:])
	case "fix":
		RunFix(args[1:])
	case "port":
		RunPort(args[1:])
	case "list":
		RunList(args[1:])
	case "version":
		RunVersion(args[1:])
	default:
		fmt.Printf("Unknown command: %s\n\n", args[0])
		RunHelp(nil)
		os.Exit(1)
	}
}

func RunInteractive() {
	fmt.Println("doctor - interactive mode\n")

	results := checks.RunAll()
	ui.RenderResults(results)
	state.PrintSummary(results)

	// collect actionable issues (Missing with fix)
	var issues []string
	for _, r := range results {
		if r.Status == 1 && r.Fix != "" {
			issues = append(issues, r.Name)
		}
	}

	if len(issues) == 0 {
		exitCode := state.ExitCode(results)
		os.Exit(exitCode)
	}

	for _, issue := range issues {
		fmt.Printf("\n---\n✗ %s\n", issue)
		fmt.Printf("→ fix %s now? (y/n): ", issue)

		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			RunFix([]string{issue})
		}
	}

	fmt.Println("\n✓ all actionable issues processed")

	fmt.Println("\n→ re-running diagnose...")
	results = checks.RunAll()
	ui.RenderResults(results)
	state.PrintSummary(results)

	exitCode := state.ExitCode(results)
	os.Exit(exitCode)
}
