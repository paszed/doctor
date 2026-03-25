package fix

import (
	"fmt"
	"os/exec"
	"runtime"
)

// FixMaven attempts to install or verify Maven
func FixMaven(args []string) error {
	fmt.Println("[FIX] maven")

	cmd := exec.Command("mvn", "-v")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ Maven is already installed")
		return nil
	}

	fmt.Println("→ Maven not found. Installing...")

	if runtime.GOOS == "darwin" {
		err := exec.Command("brew", "install", "maven").Run()
		if err != nil {
			return fmt.Errorf("failed to install Maven via brew: %w", err)
		}
	} else if runtime.GOOS == "linux" {
		err := exec.Command("sudo", "apt", "install", "-y", "maven").Run()
		if err != nil {
			return fmt.Errorf("failed to install Maven via apt: %w", err)
		}
	} else {
		fmt.Println("Please install Maven manually: https://maven.apache.org/download.cgi")
	}

	fmt.Println("✓ fix completed")
	return nil
}

// init registers this fix in the registry
func init() {
	Register("maven", FixMaven)
}
