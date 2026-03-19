package cli

import (
	"fmt"
)

func Run(args []string) {
	if len(args) == 0 {
		RunHelp(nil)
		return
	}

	switch args[0] {

	case "diagnose":
		RunDiagnose(args[1:])

	case "check":
		RunCheck(args[1:])

	case "fix":
		RunFix(args[1:])

	case "port":
		RunPort(args[1:])

	case "list":
		RunList(args[1:])

	case "version":
		RunVersion(args[1:])

	case "help", "-h", "--help":
		RunHelp(args[1:])

	default:
		fmt.Printf("Unknown command: %s\n\n", args[0])
		RunHelp(nil)
	}
}
