package ui

import (
	"fmt"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func RenderResults(results []model.Result) {
	if len(results) == 0 {
		fmt.Println("No results")
		return
	}

	fmt.Println("TOOLS")
	fmt.Println("-----")

	for _, r := range results {
		printResult(r)
	}

	fmt.Println()
	printSummary(results)
}

func printResult(r model.Result) {
	name := padRight(r.Name, 15)

	var statusSymbol string
	var message string

	switch r.Status {
	case 0:
		statusSymbol = ColorGreen("✓")
		message = r.Message
	case 1:
		statusSymbol = ColorRed("✗")
		message = r.Message
	default:
		statusSymbol = ColorYellow("!")
		message = r.Message
	}

	// main line
	fmt.Printf("%s %s %s\n", name, statusSymbol, message)

	// fix suggestion
	if r.Status != 0 && r.Fix != "" {
		fmt.Printf("   → fix: %s\n", ColorCyan(r.Fix))
	}
}

func printSummary(results []model.Result) {
	total := len(results)
	passed := 0

	for _, r := range results {
		if r.Status == 0 {
			passed++
		}
	}

	fmt.Println("SUMMARY")
	fmt.Println("-------")

	if passed == total {
		fmt.Printf("%s %d/%d checks passed\n",
			ColorGreen("✓"),
			passed,
			total,
		)
	} else {
		fmt.Printf("%s %d/%d checks passed\n",
			ColorRed("!"),
			passed,
			total,
		)
	}
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
