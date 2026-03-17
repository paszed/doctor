package ui

import (
	"fmt"

	"github.com/paszed/doctor/internal/model"
)

func PrintResults(results []model.Result) {

	Section("TOOLS")

	for _, r := range results {
		PrintResult(r)
	}

	ok := 0

	for _, r := range results {
		if r.Status == model.OK {
			ok++
		}
	}

	Section("SUMMARY")

	fmt.Printf("✓ %d/%d checks passed\n", ok, len(results))
}
