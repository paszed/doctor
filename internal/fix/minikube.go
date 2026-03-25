package fix

import (
	"fmt"
	"os/exec"
)

func FixMinikube(args []string) error {
	fmt.Println("→ Minikube not found. Installing via Homebrew...")
	cmd := exec.Command("brew", "install", "minikube")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func init() {
	Register("minikube", FixMinikube)
}
