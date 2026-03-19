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
	fmt.Println("  fix <tool>      attempt to fix an issue")
	fmt.Println("  port <port>     check a specific port")
	fmt.Println("  list            list available checks")
	fmt.Println("  version         show doctor version")
	fmt.Println()
	fmt.Println("EXAMPLES")
	fmt.Println("  doctor diagnose")
	fmt.Println("  doctor check docker")
	fmt.Println("  doctor fix docker")
	fmt.Println("  doctor fix ports 3000")
}
