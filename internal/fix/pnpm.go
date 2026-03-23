package fix

import (
	"fmt"
	"os/exec"
	"strings"
)

// FixPNPM checks if pnpm is installed and provides install instructions if not
func FixPNPM(args []string) error {
	fmt.Println("[FIX] pnpm")

	// Check if pnpm exists
	cmd := exec.Command("pnpm", "--version")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Printf("✓ pnpm is already installed: %s\n", strings.TrimSpace(string(out)))
		return nil
	}

	// Otherwise, print instructions
	fmt.Println("→ pnpm not found.")
	fmt.Println("  Install pnpm via npm: npm install -g pnpm")
	return fmt.Errorf("pnpm not installed")
}

func init() {
	Register("pnpm", FixPNPM)
}
