package runtime

import (
	"os"
	"strings"
)

// FixPath ensures that extra binary directories are added to the PATH
func FixPath() {
	// Extra paths to prepend
	extra := []string{
		"/usr/local/bin",
		"/opt/homebrew/bin",
	}

	// Get current PATH
	current := os.Getenv("PATH")

	// Prepend extra paths that are not already in PATH
	for _, p := range extra {
		if !strings.Contains(current, p) {
			current = p + ":" + current
		}
	}

	// Update the PATH environment variable
	os.Setenv("PATH", current)
}

