package fix

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// FixJava attempts to ensure Java is installed and available in PATH
func FixJava(args []string) error {
	fmt.Println("[FIX] java")

	// Try to check Java version
	cmd := exec.Command("java", "-version")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println("✓ Java is already installed")
		return nil
	}

	// If not found, print installation instructions based on OS
	fmt.Println("→ java not found.")
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("  On macOS: brew install openjdk")
	case "linux":
		fmt.Println("  On Linux (Debian/Ubuntu): sudo apt install openjdk-17-jdk")
	case "windows":
		fmt.Println("  On Windows: Download and install JDK from https://adoptopenjdk.net/")
	default:
		fmt.Println("  Please install JDK manually.")
	}

	return fmt.Errorf("java not installed or not in PATH: %s", strings.TrimSpace(string(out)))
}

// Register Java fix
func init() {
	Register("java", FixJava)
}

