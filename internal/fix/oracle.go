package fix

import (
	"fmt"
	"os/exec"
)

// FixOracle attempts to check if Oracle CLI is available
func FixOracle(args []string) error {
	fmt.Println("[FIX] oracle")

	cmd := exec.Command("sqlplus", "-V")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ Oracle CLI is already installed")
		return nil
	}

	fmt.Println("→ Oracle CLI not found. Please install it manually:")
	fmt.Println("   https://www.oracle.com/database/technologies/instant-client.html")
	return nil
}

// init registers the fix
func init() {
	Register("oracle", FixOracle)
}
