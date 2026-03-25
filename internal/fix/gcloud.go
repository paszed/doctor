package fix

import (
	"fmt"
	"os/exec"
)

func FixGCloud(args []string) error {
	fmt.Println("→ gcloud CLI not found. Installing via Homebrew...")
	cmd := exec.Command("brew", "install", "google-cloud-sdk")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func init() {
	Register("gcloud", FixGCloud)
}
