package fix

import (
	"fmt"
	"os/exec"
	"runtime"
)

// FixPostgres ensures PostgreSQL is installed and available in PATH
func FixPostgres(args []string) error {
	fmt.Println("[FIX] postgres")

	// Check if PostgreSQL is installed
	cmd := exec.Command("psql", "--version")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ PostgreSQL is already installed")
		return nil
	}

	// Install depending on OS
	if runtime.GOOS == "darwin" { // macOS
		fmt.Println("→ PostgreSQL not found. Installing via Homebrew...")
		err := exec.Command("brew", "install", "postgresql").Run()
		if err != nil {
			return fmt.Errorf("failed to install PostgreSQL via Homebrew: %w", err)
		}
	} else if runtime.GOOS == "linux" { // Linux
		fmt.Println("→ PostgreSQL not found. Installing via apt...")
		err := exec.Command("sudo", "apt", "install", "-y", "postgresql", "postgresql-client").Run()
		if err != nil {
			return fmt.Errorf("failed to install PostgreSQL via apt: %w", err)
		}
	} else {
		fmt.Println("→ Please install PostgreSQL manually: https://www.postgresql.org/download/")
	}

	fmt.Println("✓ fix completed")
	return nil
}

// Register the PostgreSQL fix
func init() {
	Register("postgres", FixPostgres)
}

