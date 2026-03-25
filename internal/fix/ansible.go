package fix

import (
	"fmt"
	"os/exec"
)

func FixAnsible(args []string) error {
	fmt.Println("[FIX] ansible")

	// Try to detect if already installed
	cmdCheck := exec.Command("ansible", "--version")
	if err := cmdCheck.Run(); err == nil {
		fmt.Println("✓ Ansible is already installed")
		return nil
	}

	// Attempt to install via brew (macOS)
	cmdInstall := exec.Command("brew", "install", "ansible")
	cmdInstall.Stdout = nil
	cmdInstall.Stderr = nil
	fmt.Println("→ ansible not found. Installing via Homebrew...")
	if err := cmdInstall.Run(); err != nil {
		fmt.Println("✗ failed: could not install Ansible automatically")
		fmt.Println("  Please install manually: brew install ansible or pip install ansible")
		return err
	}

	fmt.Println("✓ fix completed")
	return nil
}

func init() {
	Register("ansible", FixAnsible)
}
