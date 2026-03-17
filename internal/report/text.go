package report

import (
	"fmt"
	"github.com/paszed/doctor/internal/model"
)

func Print(results []model.CheckResult) {

	fmt.Println()
	fmt.Println("TOOLS")
	fmt.Println("-----")

	ok := 0

	for _, r := range results {

		if r.Status {

			fmt.Printf(
				"%-15s ✓ %s (%s)\n",
				r.Name,
				r.Version,
				r.Path,
			)

			ok++

		} else {

			fmt.Printf(
				"%-15s ✗ %s\n",
				r.Name,
				r.Message,
			)
		}
	}

	fmt.Println()
	fmt.Println("SUMMARY")
	fmt.Println("-------")

	fmt.Printf("✓ %d/%d checks passed\n", ok, len(results))
}
