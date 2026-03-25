package fix

import (
	"fmt"
	"os/exec"
)

// FixNVM installs nvm automatically via Homebrew (macOS)
func FixNVM(args []string) error {
	cmd := exec.Command("brew", "install", "nvm")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install nvm: %v", err)
	}
	return nil
}

func init() {
	Register("nvm", FixNVM)
}
