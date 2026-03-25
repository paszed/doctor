package fix

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func FixKubectl(args []string) error {
	fmt.Println("[FIX] kubectl")

	// macOS automatic installation
	if runtime.GOOS == "darwin" {
		fmt.Println("→ kubectl not found. Installing via Homebrew...")
		cmd := exec.Command("brew", "install", "kubectl")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// Linux / Windows fallback: print instructions
	fmt.Println("→ kubectl not found. Please install manually:")
	fmt.Println("  https://kubernetes.io/docs/tasks/tools/")
	return nil
}

func init() {
	Register("kubectl", FixKubectl)
}

