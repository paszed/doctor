package fix

import (
	"fmt"
	"os/exec"
)

// FixPNPM attempts to install pnpm globally via npm if not present
func FixPNPM(args []string) error {
	// Check if pnpm is already installed
	_, err := exec.LookPath("pnpm")
	if err == nil {
		fmt.Println("✓ pnpm already installed")
		return nil
	}

	fmt.Println("→ pnpm not found. Installing via npm...")
	cmd := exec.Command("npm", "install", "-g", "pnpm")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("✗ failed to install pnpm: %v", err)
	}

	fmt.Println("✓ pnpm installed")
	return nil
}

func init() {
	Register("pnpm", FixPNPM)
}
