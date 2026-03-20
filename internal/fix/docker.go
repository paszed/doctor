package fix

import (
	"fmt"
	"os/exec"
	"time"
)

func FixDocker(args []string) error {
	fmt.Println("→ starting Docker Desktop (this may take ~30–60s)...")

	// start Docker Desktop (macOS)
	err := exec.Command("open", "-a", "Docker").Start()
	if err != nil {
		return fmt.Errorf("failed to start Docker Desktop: %w", err)
	}

	// 🔥 smart wait loop (early exit when ready)
	for i := 1; i <= 30; i++ {
		time.Sleep(1 * time.Second)

		// check if daemon is ready
		cmd := exec.Command("docker", "info")
		if err := cmd.Run(); err == nil {
			// ✅ ready → let runner verify + print final result
			return nil
		}

		fmt.Printf("   waiting for daemon... (%ds)\n", i)

		if i == 10 {
			fmt.Println("   still starting... Docker can be slow on first launch")
		}
	}

	return fmt.Errorf("Docker did not start within 30 seconds. Try opening Docker Desktop manually.")
}

func init() {
	// 🔥 MUST match checks + CLI exactly
	Register("docker", FixDocker)
}
