package fix

import (
	"fmt"
	"os/exec"
	"time"
)

func init() {
	Register("docker", FixDocker)
}

func FixDocker(args []string) error {
	fmt.Println("→ starting Docker Desktop (this may take ~30–60s)...")

	// Start Docker Desktop (macOS)
	cmd := exec.Command("open", "-a", "Docker")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Docker: %w", err)
	}

	maxWait := 60 // seconds

	for i := 0; i < maxWait; i++ {
		time.Sleep(1 * time.Second)

		// show hint after 10 seconds
		if i == 10 {
			fmt.Print("\n   still starting... Docker can be slow on first launch\n")
		}

		// overwrite same line (clean progress)
		fmt.Printf("\r   waiting for daemon... (%ds)", i+1)

		check := exec.Command("docker", "info")
		if err := check.Run(); err == nil {
			fmt.Print("\r\033[K") // clear line completely
			fmt.Println("✓ Docker daemon is running")
			return nil
		}
	}

	fmt.Println() // ensure clean newline

	return fmt.Errorf(
		"Docker daemon not ready after %ds.\nTry:\n  open -a Docker\n  wait until Docker Desktop fully starts\n  then run: doctor fix docker",
		maxWait,
	)
}
