package ui

import (
	"fmt"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func OK(name string, msg string) {
	fmt.Printf("%-16s %s✓%s %s\n", name, Green, Reset, msg)
}

func Fail(name string, msg string) {
	fmt.Printf("%-16s %s✗%s %s\n", name, Red, Reset, msg)
}

func Warn(name string, msg string) {
	fmt.Printf("%-16s %s!%s %s\n", name, Yellow, Reset, msg)
}

func PrintResult(r model.Result) {
	switch r.Status {

	case model.OK:
		OK(r.Name, r.Message)

	case model.Missing:
		Fail(r.Name, r.Message)

	case model.Warning:
		Warn(r.Name, r.Message)

	default:
		fmt.Printf("%-16s ? %s\n", r.Name, r.Message)
	}
}

func Section(title string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)+4))
}
