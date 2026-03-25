package fix

import (
	"fmt"
	"os/exec"
)

// FixYarn attempts to install yarn globally via npm if not present
func FixYarn(args []string) error {
	// Check if yarn is already installed
	_, err := exec.LookPath("yarn")
	if err == nil {
		fmt.Println("✓ yarn already installed")
		return nil
	}

	fmt.Println("→ yarn not found. Installing via npm...")
	cmd := exec.Command("npm", "install", "-g", "yarn")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("✗ failed to install yarn: %v", err)
	}

	fmt.Println("✓ yarn installed")
	return nil
}

func init() {
	Register("yarn", FixYarn)
}
