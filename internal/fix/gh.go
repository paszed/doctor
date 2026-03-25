package fix

import (
	"fmt"
	"os/exec"
)

// FixGH tries to install GitHub CLI
func FixGH(args []string) error {
	// Check if already installed
	cmd := exec.Command("gh", "--version")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ GitHub CLI is already installed")
		return nil
	}

	fmt.Println("→ gh not found. Installing via Homebrew...")
	installCmd := exec.Command("brew", "install", "gh")
	installCmd.Stdout = nil
	installCmd.Stderr = nil

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install GitHub CLI: %w", err)
	}

	fmt.Println("✓ GitHub CLI installed")
	return nil
}

func init() {
	// Register as a fix, not a check
	Register("gh", FixGH)
}
