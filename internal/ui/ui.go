package ui

import (
	"fmt"

	"github.com/paszed/doctor/internal/model"
)

func PrintResult(r model.Result) {
	switch r.Status {
	case model.OK:
		fmt.Printf("%s %s\n", ColorGreen("✓"), r.Message)

	case model.Missing:
		fmt.Printf("%s %s\n", ColorRed("✗"), r.Message)

	case model.Warning:
		fmt.Printf("%s %s\n", ColorYellow("!"), r.Message)

	default:
		fmt.Printf("? %s\n", r.Message)
	}
}

func PrintResults(results []model.Result) {
	for _, r := range results {
		PrintResult(r)
	}
}
