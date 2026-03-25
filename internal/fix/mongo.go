package fix

import (
	"fmt"
	"os/exec"
)

// FixMongo installs MongoDB CLI if missing.
func FixMongo(args []string) error {
	fmt.Println("[FIX] mongo")

	cmd := exec.Command("mongo", "--version")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ MongoDB CLI is already installed")
		return nil
	}

	// Attempt auto-install on macOS via Homebrew
	install := exec.Command("brew", "install", "mongodb/brew/mongodb-community-shell")
	install.Stdout = nil
	install.Stderr = nil
	if err := install.Run(); err != nil {
		fmt.Println("→ MongoDB CLI not found. Please install manually:")
		fmt.Println("   https://www.mongodb.com/docs/mongodb-shell/install/")
		return nil
	}

	fmt.Println("✓ fix completed")
	return nil
}

func init() {
	Register("mongo", FixMongo)
}
