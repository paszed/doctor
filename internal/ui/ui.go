package ui

import (
	"fmt"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

// --- status printers ---

func OK(name string, msg string) {
	msg = strings.TrimSpace(msg)
	fmt.Printf("%-16s %s %s\n", name, Success("✓"), msg)
}

func Fail(name string, msg string) {
	msg = strings.TrimSpace(msg)
	fmt.Printf("%-16s %s %s\n", name, Error("✗"), msg)
}

func Warn(name string, msg string) {
	msg = strings.TrimSpace(msg)
	fmt.Printf("%-16s %s %s\n", name, Warning("!"), msg)
}

// --- single result ---

func PrintResult(r model.Result) {
	switch r.Status {
	case model.OK:
		OK(r.Name, r.Message)

	case model.Missing:
		Fail(r.Name, r.Message)

	case model.Warning:
		Warn(r.Name, r.Message)

	default:
		msg := strings.TrimSpace(r.Message)
		fmt.Printf("%-16s %s\n", r.Name, msg)
	}
}

// --- section header ---

func Section(title string) {
	fmt.Printf("\n%s\n", title)
	fmt.Println(strings.Repeat("-", len(title)))
}
