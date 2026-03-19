package ui

import "os"

var NoColor = false

func init() {
	if os.Getenv("NO_COLOR") != "" {
		NoColor = true
	}
}

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	gray   = "\033[90m"
)

// base color function
func colorize(color, s string) string {
	if NoColor {
		return s
	}
	return color + s + reset
}

// public helpers

func ColorGreen(s string) string {
	return colorize(green, s)
}

func ColorRed(s string) string {
	return colorize(red, s)
}

func ColorYellow(s string) string {
	return colorize(yellow, s)
}

func ColorCyan(s string) string {
	return colorize(cyan, s)
}

func ColorGray(s string) string {
	return colorize(gray, s)
}
