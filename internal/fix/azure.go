package fix

import (
	"fmt"
	"os/exec"
)

func FixAzure(args []string) error {
	fmt.Println("→ az CLI not found. Installing via Homebrew...")
	cmd := exec.Command("brew", "install", "azure-cli")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func init() {
	Register("az", FixAzure)
}
