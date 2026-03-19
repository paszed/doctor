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
	fmt.Println("→ starting Docker...")

	// macOS: start Docker Desktop
	cmd := exec.Command("open", "-a", "Docker")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start Docker: %w", err)
	}

	// Wait until Docker is actually running
	for i := 0; i < 10; i++ {
		check := exec.Command("docker", "info")
		if err := check.Run(); err == nil {
			fmt.Println("✓ Docker started")
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("Docker is still not running (try opening Docker Desktop manually)")
}
