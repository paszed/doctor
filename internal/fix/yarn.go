package fix

import (
	"fmt"
	"os/exec"
	"strings"
)

// FixYarn checks if Yarn is installed and provides install instructions if not
func FixYarn(args []string) error {
	fmt.Println("[FIX] yarn")

	// Check if yarn exists
	cmd := exec.Command("yarn", "--version")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Printf("✓ yarn is already installed: %s\n", strings.TrimSpace(string(out)))
		return nil
	}

	// Otherwise, print instructions
	fmt.Println("→ yarn not found.")
	fmt.Println("  Install yarn via npm: npm install -g yarn")
	return fmt.Errorf("yarn not installed")
}

func init() {
	Register("yarn", FixYarn)
}
