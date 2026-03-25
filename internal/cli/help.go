package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
)

// RunHelp prints a detailed help message listing all available checks and fixes.
func RunHelp(args []string) {
	fmt.Println("doctor - developer environment diagnostics\n")
	fmt.Println("USAGE")
	fmt.Println("  doctor <command> [options]\n")
	fmt.Println("COMMANDS")
	fmt.Println("  diagnose        run full environment checks")
	fmt.Println("  check <tool>    run a single tool check")
	fmt.Println("  fix [tool]      fix all issues or a specific tool")
	fmt.Println("  port <port>     check a specific port")
	fmt.Println("  list            list available checks")
	fmt.Println("  version         show doctor version\n")

	fmt.Println("EXAMPLES")
	fmt.Println("  doctor diagnose")
	fmt.Println("  doctor check docker")
	fmt.Println("  doctor check java")
	fmt.Println("  doctor fix")
	fmt.Println("  doctor fix docker")
	fmt.Println("  doctor fix java")
	fmt.Println("  doctor fix ports 3000\n")

	// List all available checks dynamically
	fmt.Println("AVAILABLE CHECKS")
	fmt.Println("----------------")
	for _, name := range checks.List() {
		fmt.Printf("  %s\n", name)
	}

	fmt.Println()
	if len(args) > 0 {
		fmt.Printf("Unknown argument(s): %v\n", args)
		os.Exit(1)
	}
}

