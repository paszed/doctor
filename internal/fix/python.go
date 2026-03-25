package fix

import (
	"fmt"
	"os/exec"
)

// FixPython ensures Python is installed and functional.
func FixPython(args []string) error {
	cmd := exec.Command("python3", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("python3 not found or not working: %w", err)
	}
	fmt.Println("✓ Python is already installed")
	return nil
}

func init() {
	Register("python3", FixPython)
}

