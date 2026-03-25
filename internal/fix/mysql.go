package fix

import (
	"fmt"
	"os/exec"
)

// FixMySQL attempts to install MySQL client automatically on macOS.
func FixMySQL(args []string) error {
	cmd := exec.Command("brew", "install", "mysql")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install MySQL: %v", err)
	}
	return nil
}

func init() {
	Register("mysql", FixMySQL)
}
