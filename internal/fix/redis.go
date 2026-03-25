package fix

import (
	"fmt"
	"os/exec"
)

// FixRedis attempts to install Redis CLI automatically on macOS.
func FixRedis(args []string) error {
	cmd := exec.Command("brew", "install", "redis")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install Redis CLI: %v", err)
	}
	return nil
}

func init() {
	Register("redis", FixRedis) // ✅ Only register as FixFunc
}
