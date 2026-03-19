package cli

import (
	"fmt"

	"github.com/paszed/doctor/internal/checks"
)

func RunPort(args []string) {
	if len(args) == 0 {
		fmt.Println("port required (e.g. doctor port 3000)")
		return
	}

	port := args[0]

	result := checks.CheckPort(port)

	fmt.Printf("%s\n", result.Message)

	if result.Status != 0 && result.Fix != "" {
		fmt.Printf("→ fix: %s\n", result.Fix)
	}
}
