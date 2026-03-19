package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunCheck(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: doctor check <tool> [--json|--version|--path]")
		os.Exit(1)
	}

	tool := args[0]

	checkFunc, ok := checks.Get(tool)
	if !ok {
		fmt.Printf("unknown tool: %s\n\n", tool)
		fmt.Println("available checks:")
		for _, name := range checks.List() {
			fmt.Printf("  %s\n", name)
		}
		os.Exit(1)
	}

	result := checkFunc()

	// Flags
	if len(args) > 1 {
		flag := args[1]

		switch flag {
		case "--json":
			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return

		case "--version":
			fmt.Println(result.Message)
			return

		case "--path":
			if result.Path != "" {
				fmt.Println(result.Path)
			}
			return
		}
	}

	// Default → unified UI
	ui.RenderResults([]model.Result{result})

	// Exit code
	if result.Status != model.OK {
		os.Exit(1)
	}
}
