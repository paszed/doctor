package cli

import (
	"fmt"
)

func Run(args []string) {
	if len(args) == 0 {
		RunHelp()
		return
	}

	switch args[0] {

	case "diagnose":
		RunDiagnose()

	case "check":
		RunCheck()

	case "fix":
		RunFix()

	case "port":
		RunPort(args[1:])

	case "list":
		RunList()

	case "version":
		RunVersion()

	case "help", "-h", "--help":
		RunHelp()

	default:
		fmt.Printf("Unknown command: %s\n\n", args[0])
		RunHelp()
	}
}
