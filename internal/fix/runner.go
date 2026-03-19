package fix

import (
	"fmt"
)

// Run executes a fix by name with arguments
func Run(name string, args []string) error {
	fn, ok := Get(name)
	if !ok {
		return fmt.Errorf("unknown fix: %s", name)
	}

	return fn(args)
}

// RunAll executes all registered fixes (auto-fix mode)
func RunAll() {
	for name, fn := range registry {
		fmt.Printf("[FIX] %s\n", name)

		err := fn([]string{}) // no args for auto-fix

		if err != nil {
			fmt.Printf("✗ %v\n", err)
		} else {
			fmt.Println("✓ fix completed")
		}

		fmt.Println()
	}
}
