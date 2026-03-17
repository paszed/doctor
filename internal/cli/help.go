package cli

import "fmt"

func RunHelp() {

	fmt.Println("doctor - developer environment diagnostics")
	fmt.Println()

	fmt.Println("USAGE")
	fmt.Println("  doctor <command> [options]")
	fmt.Println()

	fmt.Println("COMMANDS")
	fmt.Println("  diagnose        run full environment checks")
	fmt.Println("  check <tool>    run a single tool check")
	fmt.Println("  list            list available checks")
	fmt.Println("  version         show doctor version")
	fmt.Println()

	fmt.Println("EXAMPLES")
	fmt.Println("  doctor diagnose")
	fmt.Println("  doctor diagnose --json")
	fmt.Println("  doctor check docker")
	fmt.Println("  doctor check docker --version")
	fmt.Println("  doctor check docker --path")
	fmt.Println()
}
