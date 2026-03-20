package ui

import (
	"fmt"
	"sort"

	"github.com/paszed/doctor/internal/model"
)

func RenderResults(results []model.Result) {
	// Sort results by name
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	fmt.Println("TOOLS")
	fmt.Println("-----")

	for _, r := range results {
		name := fmt.Sprintf("%-10s", r.Name)

		switch r.Status {
		case model.OK:
			fmt.Printf("%s %s %s\n",
				name,
				ColorGreen("✓"),
				r.Message,
			)

		case model.Missing:
			fmt.Printf("%s %s %s\n",
				name,
				ColorRed("✗"),
				r.Message,
			)

			if r.Fix != "" {
				fmt.Printf("            fix → %s\n", ColorCyan(r.Fix))
			}

		case model.Warning:
			fmt.Printf("%s %s %s\n",
				name,
				ColorYellow("!"),
				r.Message,
			)
		}
	}

	fmt.Println()
}
