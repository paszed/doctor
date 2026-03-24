package fix

import (
	"fmt"
	"os"
	"os/exec"
)

// FixYarn checks if Yarn is installed and attempts to install it via npm if missing.
func FixYarn(args []string) error {
	// Check if Yarn is already installed
	if _, err := exec.LookPath("yarn"); err == nil {
		fmt.Println("✓ yarn installed")
		return nil
	}

	// Not found, attempt installation via npm
	fmt.Println("→ yarn not found. Installing via npm...")
	cmd := exec.Command("npm", "install", "-g", "yarn")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("✗ failed: yarn not installed")
		return err
	}

	fmt.Println("✓ fix completed")
	return nil
}

func init() {
	Register("yarn", FixYarn)
}

