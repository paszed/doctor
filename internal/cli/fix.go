package cli

import (
	"fmt"
	"strings"

	"github.com/paszed/doctor/internal/checks"
	"github.com/paszed/doctor/internal/fix"
	"github.com/paszed/doctor/internal/model"
	"github.com/paszed/doctor/internal/ui"
)

func RunFix(args []string) {
	// AUTO MODE: doctor fix
	if len(args) == 0 {
		runAutoFix()
		return
	}

	// SPECIFIC FIX: doctor fix <tool>
	name := args[0]

	fn, ok := fix.Get(name)
	if !ok {
		fmt.Printf("unknown fix: %s\n", name)
		return
	}

	fmt.Printf("[FIX] %s\n", name)

	err := fn()
	if err != nil {
		fmt.Printf("✗ fix failed: %v\n", err)
		return
	}

	fmt.Println("✓ fix completed")
}

func runAutoFix() {
	fmt.Println("[AUTO FIX] running diagnose...")

	results := checks.RunAll()

	var failed []model.Result

	for _, r := range results {
		if r.Status == model.Missing {
			failed = append(failed, r)
		}
	}

	if len(failed) == 0 {
		fmt.Println("✓ nothing to fix")
		return
	}

	fmt.Printf("→ found %d issues\n\n", len(failed))

	for _, r := range failed {
		fmt.Printf("[FIX] %s\n", r.Name)

		// map check name → fix name
		name := r.Name

		if strings.HasPrefix(name, "port:") {
			name = "ports"
		}

		fn, ok := fix.Get(name)
		if !ok {
			fmt.Println("  → no fix available\n")
			continue
		}

		err := fn()
		if err != nil {
			fmt.Printf("  ✗ failed: %v\n\n", err)
			continue
		}

		fmt.Println("  ✓ fixed\n")
	}

	fmt.Println("→ re-running diagnose...\n")

	results = checks.RunAll()
	ui.RenderResults(results)
}
