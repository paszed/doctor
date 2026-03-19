package ui

// --- ANSI color codes ---
const (
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

// --- global toggle ---
var UseColor = true

// --- helpers ---

func Colorize(color string, text string) string {
	if !UseColor {
		return text
	}
	return color + text + Reset
}

// --- semantic helpers (cleaner usage) ---

func Success(text string) string {
	return Colorize(Green, text)
}

func Error(text string) string {
	return Colorize(Red, text)
}

func Warning(text string) string {
	return Colorize(Yellow, text)
}
