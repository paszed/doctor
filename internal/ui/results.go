package ui

import (
	"fmt"
	"sort"

	"github.com/paszed/doctor/internal/model"
)

func RenderResults(results []model.Result) {
	if len(results) == 0 {
		fmt.Println("No results")
		return
	}

	// sort results alphabetically
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	// dynamic column width
	maxWidth := 0
	for _, r := range results {
		if len(r.Name) > maxWidth {
			maxWidth = len(r.Name)
		}
	}
	maxWidth += 2

	fmt.Println(ColorGray("TOOLS"))
	fmt.Println(ColorGray("-----"))

	for _, r := range results {
		name := padRight(r.Name, maxWidth)

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
				fmt.Printf("   fix → %s\n\n", ColorCyan(r.Fix))
			}

		case model.Warning:
			fmt.Printf("%s %s %s\n\n",
				name,
				ColorYellow("!"),
				r.Message,
			)
		}
	}

	printSummary(results)
}

func printSummary(results []model.Result) {
	total := len(results)
	passed := 0
	warn := 0
	missing := 0

	for _, r := range results {
		switch r.Status {
		case model.OK:
			passed++
		case model.Warning:
			warn++
		case model.Missing:
			missing++
		}
	}

	fmt.Println(ColorGray("SUMMARY"))
	fmt.Println(ColorGray("-------"))

	if missing > 0 {
		fmt.Printf("%s %d/%d checks passed\n", ColorRed("✗"), passed, total)
	} else if warn > 0 {
		fmt.Printf("%s %d/%d checks passed\n", ColorYellow("!"), passed, total)
	} else {
		fmt.Printf("%s %d/%d checks passed\n", ColorGreen("✓"), passed, total)
	}
}

func padRight(str string, length int) string {
	for len(str) < length {
		str += " "
	}
	return str
}
