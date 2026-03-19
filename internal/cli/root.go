package cli

import (
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) < 2 {
		RunHelp()
		return
	}

	cmd := os.Args[1]

	switch cmd {

	case "help", "-h", "--help":
		RunHelp()
		return

	case "diagnose":
		RunDiagnose()
		return

	case "check":
		RunCheck()
		return

	case "list":
		RunList()
		return

	case "port":
		RunPort()
		return

	case "version":
		RunVersion()
		return

	case "fix":
		RunFix()
		return

	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println()
		RunHelp()
		return
	}
}
