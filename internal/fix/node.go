package fix

import (
	"fmt"
	"os/exec"
)

func init() {
	Register("node", FixNode)
}

func FixNode(args []string) error {
	_ = args

	// check if node exists
	_, err := exec.LookPath("node")
	if err == nil {
		fmt.Println("✓ already installed")
		return nil
	}

	fmt.Println("→ installing Node.js...")

	// macOS (Homebrew)
	cmd := exec.Command("brew", "install", "node")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install node: %v", err)
	}

	fmt.Println("✓ Node installed")
	return nil
}
