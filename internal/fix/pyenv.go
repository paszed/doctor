package fix

import (
	"fmt"
	"os/exec"
)

// FixPyenv installs pyenv automatically on macOS
func FixPyenv(args []string) error {
	cmd := exec.Command("brew", "install", "pyenv")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install pyenv: %v", err)
	}
	return nil
}

func init() {
	Register("pyenv", FixPyenv)
}
