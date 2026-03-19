package fix

import (
	"fmt"
	"os/exec"
)

func FixNode() error {
	fmt.Println("[FIX] node")

	// check if node exists
	if _, err := exec.LookPath("node"); err == nil {
		fmt.Println("✓ already installed")
		return nil
	}

	fmt.Println("→ node not found")

	// suggest install (DO NOT auto-install yet)
	fmt.Println("→ install via:")
	fmt.Println("   brew install node")
	fmt.Println("   or https://nodejs.org")

	return nil
}

func init() {
	Register("node", FixNode)
}
