package fix

import (
	"fmt"
	"os/exec"
	"strings"
)

// FixPostgres attempts to detect PostgreSQL and provide guidance to install it
func FixPostgres(args []string) error {
	fmt.Println("[FIX] postgres")

	// Check if psql exists
	cmd := exec.Command("psql", "--version")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Printf("✓ PostgreSQL is already installed: %s\n", strings.TrimSpace(string(out)))
		return nil
	}

	// Otherwise, print instructions
	fmt.Println("→ PostgreSQL not found.")
	fmt.Println("  Please install PostgreSQL and ensure `psql` is in your PATH.")
	fmt.Println("  On macOS: brew install postgresql")
	fmt.Println("  On Linux (Debian/Ubuntu): sudo apt install postgresql postgresql-client")
	fmt.Println("  On Windows: https://www.postgresql.org/download/windows/")
	return fmt.Errorf("postgres not installed")
}

func init() {
	Register("postgres", FixPostgres)
}
