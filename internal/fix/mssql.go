package fix

import (
	"fmt"
)

func FixMSSQL(args []string) error {
	fmt.Println("→ sqlcmd not found. Please install manually:")
	fmt.Println("  Visit: https://learn.microsoft.com/en-us/sql/tools/sqlcmd-utility")
	return nil
}

func init() {
	Register("mssql", FixMSSQL)
}
