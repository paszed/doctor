package state

import (
	"fmt"

	"github.com/paszed/doctor/internal/model"
)

// Summary aggregates results into counts
func Summary(results []model.Result) (ok, warn, missing int) {
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
	return
}

// Status returns overall status priority
func Status(results []model.Result) model.Status {
	ok, warn, missing := Summary(results)

	if missing > 0 {
		return model.Missing
	}
	if warn > 0 {
		return model.Warning
	}
	if ok > 0 {
		return model.OK
	}

	return model.OK
}

// ExitCode defines CLI exit behavior
func ExitCode(results []model.Result) int {
	ok, warn, missing := Summary(results)

	if missing > 0 {
		return 2 // errors
	}
	if warn > 0 {
		return 1 // warnings only
	}
	if ok > 0 {
		return 0 // all good
	}

	return 0
}

// PrintSummary renders CLI summary output
func PrintSummary(results []model.Result) {
	ok, warn, missing := Summary(results)
	total := ok + warn + missing

	symbol := "✓"
	if missing > 0 {
		symbol = "✗"
	} else if warn > 0 {
		symbol = "!"
	}

	fmt.Println("SUMMARY")
	fmt.Println("-------")
	fmt.Printf("%s %d/%d checks passed\n\n", symbol, ok, total)
}
