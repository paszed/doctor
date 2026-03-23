package fix

import (
	"fmt"
	"os/exec"
	"time"
)

// FixJava attempts to ensure Java is installed and in PATH.
func FixJava(args []string) error {
	fmt.Println("[FIX] java")
	// Check if java is already installed
	_, err := exec.LookPath("java")
	if err == nil {
		fmt.Println("✓ Java is already installed")
		return nil
	}

	fmt.Println("→ Java not found, attempting to install...")

	// Example for macOS: install openjdk via brew
	cmd := exec.Command("brew", "install", "openjdk")
	err = cmd.Run()
	if err != nil {
		fmt.Println("✗ failed to install Java automatically")
		fmt.Println("→ Please install JDK manually and ensure `java` is in PATH")
		return err
	}

	// Give system a few seconds to refresh PATH
	time.Sleep(2 * time.Second)

	fmt.Println("✓ Java installed")
	return nil
}

// Register in the fix registry
func init() {
	Register("java", FixJava)
}
