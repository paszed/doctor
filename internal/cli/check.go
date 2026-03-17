package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
)

func RunCheck() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: doctor check <tool>")
		return
	}

	tool := os.Args[2]

	if tool == "docker" {
		checks.CheckDocker()
		return
	}

	checks.ToolVersion(tool)
}
