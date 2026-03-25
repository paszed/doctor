package fix

import (
	"fmt"
	"os/exec"
)

func FixKind(args []string) error {
	fmt.Println("→ Kind not found. Installing via Homebrew...")
	cmd := exec.Command("brew", "install", "kind")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func init() {
	Register("kind", FixKind)
}
