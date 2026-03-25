package fix

import (
	"fmt"
	"os"
	"os/exec"
)

// FixPipenv installs pipenv
func FixPipenv(args []string) error {
	cmd := exec.Command("pipx", "install", "pipenv")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install pipenv: %w", err)
	}
	return nil
}

func init() {
	Register("pipenv", FixPipenv)
}
