package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunPort(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: doctor port <port>")
		os.Exit(1)
	}

	port := args[0]

	result := checks.CheckPort(port)

	// Consistent UI rendering (same as diagnose/check)
	ui.RenderResults([]model.Result{result})

	// Exit code behavior
	if result.Status != model.OK {
		os.Exit(1)
	}
}
