package fix

import (
	"fmt"
	"os/exec"
)

// FixPostgres attempts to check and guide installation of PostgreSQL
func FixPostgres(args []string) error {
	fmt.Println("[FIX] postgres")

	// Check if psql is available
	_, err := exec.LookPath("psql")
	if err == nil {
		fmt.Println("✓ PostgreSQL is already installed")
		return nil
	}

	// Not found, print OS-specific instructions
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

