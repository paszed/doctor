package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/fix"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/state"
	"github.com/paszed/doctor/internal/ui"
)

func Run(args []string) {
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
		fmt.Printf("unknown command: %s\n\n", args[0])
		RunHelp(nil)
		os.Exit(1)
	}
}

func RunInteractive() {
	fmt.Println("doctor - interactive mode\n")

	results := checks.RunAll()
	ui.RenderResults(results)

	// collect only real issues (Missing)
	var issues []model.Result
	for _, r := range results {
		if r.Status == model.Missing && r.Fix != "" {
			issues = append(issues, r)
		}
	}

	if len(issues) == 0 {
		fmt.Println("\n✓ no issues found")
		return
	}

	fmt.Println("\n---")

	reader := bufio.NewReader(os.Stdin)

	for _, issue := range issues {
		fmt.Printf("✗ %s\n", issue.Name)
		fmt.Printf("→ fix now? (y/n): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" {
			parts := strings.Fields(issue.Fix)

			if len(parts) >= 3 {
				tool := parts[2]

				fn, ok := fix.Get(tool)
				if ok {
					fmt.Printf("\n[FIX] %s\n", tool)

					var fixArgs []string
					if len(parts) > 3 {
						fixArgs = parts[3:]
					}

					if err := fn(fixArgs); err != nil {
						fmt.Printf("✗ failed: %v\n", err)
					} else {
						fmt.Println("✓ fixed")
					}
				} else {
					fmt.Println("✗ no fix available")
				}
			}
		}

		fmt.Println()
	}

	fmt.Println("→ re-running diagnose...")

	results = checks.RunAll()
	ui.RenderResults(results)

	_, _, exitCode := state.Summary(results)
	os.Exit(exitCode)
}
