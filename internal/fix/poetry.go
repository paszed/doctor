package fix

import (
	"fmt"
	"os"
	"os/exec"
)

func FixPoetry(args []string) error {
	cmd := exec.Command("pipx", "install", "poetry")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install poetry: %w", err)
	}
	return nil
}

func init() {
	Register("poetry", FixPoetry)
}

