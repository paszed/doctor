package cli

import (
	"fmt"
	"os"
)

func Run() {

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {

	case "diagnose":
		RunDiagnose()

	case "check":
		RunCheck()

	case "port":
		RunPort()

	case "version":
		RunVersion()

	case "fix":
		RunFix()

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
