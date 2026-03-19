package ui

import (
	"fmt"

	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/state"
)

func RenderResults(results []model.Result) {
	fmt.Println("TOOLS")
	fmt.Println("-----")

	maxWidth := maxNameWidth(results)

	for _, r := range results {
		name := padRight(r.Name, maxWidth)

		switch r.Status {
		case model.OK:
			fmt.Printf("%s   %s %s\n",
				name,
				ColorGreen("✓"),
				r.Message,
			)

		case model.Missing:
			fmt.Printf("%s   %s %s\n",
				name,
				ColorRed("✗"),
				r.Message,
			)

			if r.Fix != "" {
				fmt.Printf("%s   %s\n",
					padRight("", maxWidth),
					ColorCyan("fix → "+r.Fix),
				)
				fmt.Println()
			}

		case model.Warning:
			fmt.Printf("%s   %s %s\n",
				name,
				ColorYellow("!"),
				r.Message,
			)
		}
	}

	fmt.Println()
	renderSummary(results)
}

func renderSummary(results []model.Result) {
	ok, warn, missing := state.Summary(results)
	total := len(results)

	fmt.Println("SUMMARY")
	fmt.Println("-------")

	switch {
	case missing > 0:
		fmt.Printf("%s %d/%d checks passed\n",
			ColorRed("✗"),
			ok,
			total,
		)
	case warn > 0:
		fmt.Printf("%s %d/%d checks passed\n",
			ColorYellow("!"),
			ok,
			total,
		)
	default:
		fmt.Printf("%s %d/%d checks passed\n",
			ColorGreen("✓"),
			ok,
			total,
		)
	}
}

func maxNameWidth(results []model.Result) int {
	max := 0
	for _, r := range results {
		if len(r.Name) > max {
			max = len(r.Name)
		}
	}
	return max
}

func padRight(s string, width int) string {
	format := fmt.Sprintf("%%-%ds", width)
	return fmt.Sprintf(format, s)
}
