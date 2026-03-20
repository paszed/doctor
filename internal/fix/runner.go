package fix

import (
	"fmt"
	"time"

	"github.com/paszed/doctor/internal/checks"
)

const maxRetries = 2

func Run(name string, args []string) error {
	fn, ok := registry[name]
	if !ok {
		return fmt.Errorf("unknown fix: %s", name)
	}

	fmt.Printf("[FIX] %s\n", name)

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {

		// 🔧 run fix
		err := fn(args)
		if err != nil {
			lastErr = err
			fmt.Printf("  ✗ attempt %d failed: %v\n", attempt, err)

			if attempt < maxRetries {
				fmt.Println("→ retrying...")
			}
			continue
		}

		// 🔍 verify via check registry
		checkFn, exists := checks.Get(name)
		if exists {
			time.Sleep(2 * time.Second)

			result := checkFn()

			if result.Status == 0 {
				fmt.Println("✓ fixed")
				return nil
			}

			fmt.Printf("  ! still failing: %s\n", result.Message)

			lastErr = fmt.Errorf(result.Message)

			if attempt < maxRetries {
				fmt.Println("→ retrying...")
			}

			continue
		}

		// fallback (no check available)
		fmt.Println("✓ fix completed")
		return nil
	}

	return fmt.Errorf("fix failed after %d attempts: %v", maxRetries, lastErr)
}

