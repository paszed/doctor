package ui

import "fmt"

func OK(name string, version string) {
	fmt.Printf("%-12s ✓ %s\n", name, version)
}

func Fail(name string, msg string) {
	fmt.Printf("%-12s ✗ %s\n", name, msg)
}

