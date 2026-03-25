package fix

import (
	"fmt"
)

func FixSQLite(args []string) error {
	fmt.Println("→ sqlite3 not found. Please install manually:")
	fmt.Println("  On macOS: brew install sqlite")
	fmt.Println("  On Linux: sudo apt install sqlite3")
	return nil
}

func init() {
	Register("sqlite3", FixSQLite)
}
