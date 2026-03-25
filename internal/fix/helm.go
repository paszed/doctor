package fix

import (
	"fmt"
	"os/exec"
)

func FixHelm(args []string) error {
	cmd := exec.Command("helm", "version", "--short")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ Helm is already installed")
		return nil
	}

	fmt.Println("→ helm not found. Installing via Homebrew...")
	install := exec.Command("brew", "install", "helm")
	install.Stdout = nil
	install.Stderr = nil
	if err := install.Run(); err != nil {
		return fmt.Errorf("✗ failed to install Helm: %v", err)
	}

	fmt.Println("✓ fix completed")
	return nil
}

func init() {
	Register("helm", FixHelm)
}
