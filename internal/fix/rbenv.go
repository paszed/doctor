package fix

import (
	"fmt"
	"os/exec"
)

// FixRbenv installs rbenv automatically on macOS
func FixRbenv(args []string) error {
	cmd := exec.Command("brew", "install", "rbenv")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install rbenv: %v", err)
	}
	return nil
}

func init() {
	Register("rbenv", FixRbenv)
}
