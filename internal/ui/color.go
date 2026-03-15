package ui

const (
	Green = "\033[32m"
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func Success(msg string) string {
	return Green + msg + Reset
}

func Error(msg string) string {
	return Red + msg + Reset
}
