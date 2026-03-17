package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunCheck() {

	if len(os.Args) < 3 {
		fmt.Println("usage: doctor check <tool>")
		os.Exit(1)
	}

	name := os.Args[2]

	result, found := checks.RunOne(name)

	if !found {
		fmt.Printf("unknown tool: %s\n", name)
		os.Exit(1)
	}

	ui.PrintResults([]model.Result{result})
}
