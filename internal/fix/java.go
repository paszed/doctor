package fix

import (
	"fmt"
	"os/exec"
	"runtime"
)

// FixJava ensures Java is installed and available in PATH
func FixJava(args []string) error {
	fmt.Println("[FIX] java")

	// Check if Java is already installed
	cmd := exec.Command("java", "-version")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ Java is already installed")
		return nil
	}

	// Install Java depending on the OS
	if runtime.GOOS == "darwin" { // macOS
		fmt.Println("→ Java not found. Installing via Homebrew...")
		err := exec.Command("brew", "install", "openjdk").Run()
		if err != nil {
			return fmt.Errorf("failed to install Java via Homebrew: %w", err)
		}
	} else if runtime.GOOS == "linux" { // Linux
		fmt.Println("→ Java not found. Installing via apt...")
		err := exec.Command("sudo", "apt", "install", "-y", "default-jdk").Run()
		if err != nil {
			return fmt.Errorf("failed to install Java via apt: %w", err)
		}
	} else {
		fmt.Println("→ Please install Java manually: https://www.java.com")
	}

	fmt.Println("✓ fix completed")
	return nil
}

// Register the Java fix on init
func init() {
	Register("java", FixJava)
}
