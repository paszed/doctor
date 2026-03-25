package fix

import (
	"fmt"
	"os/exec"
)

func FixAWS(args []string) error {
	fmt.Println("→ aws not found. Installing via Homebrew...")
	cmd := exec.Command("brew", "install", "awscli")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func init() {
	Register("aws", FixAWS)
}
