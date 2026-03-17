package cli

import (
	"fmt"
	"os"
)

func Run() {

	if len(os.Args) < 2 ||
		os.Args[1] == "help" ||
		os.Args[1] == "-h" ||
		os.Args[1] == "--help" {

		RunHelp()
		return
	}
	switch os.Args[1] {

	case "diagnose":
		RunDiagnose()

	case "check":
		RunCheck()

	case "list":
		RunList()

	case "port":
		RunPort()

	case "version":
		RunVersion()

	case "fix":
		RunFix()

	case "help":
		RunHelp()

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}
}

func printHelp() {

	fmt.Println("Doctor — system diagnostics tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  doctor diagnose")
	fmt.Println("  doctor check <tool>")
	fmt.Println("  doctor port <port>")
	fmt.Println("  doctor fix")
	fmt.Println("  doctor version")
}
