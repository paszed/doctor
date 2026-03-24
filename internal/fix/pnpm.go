package fix

import (
	"fmt"
	"os/exec"
)

// FixPNPM tries to install pnpm globally using npm.
func FixPNPM(args []string) error {
	cmd := exec.Command("npm", "install", "-g", "pnpm")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("→ pnpm not found.")
		fmt.Println("  Install via npm manually: npm install -g pnpm")
		return err
	}
	fmt.Printf("✓ pnpm installed\nOutput:\n%s\n", out)
	return nil
}

func init() {
	Register("pnpm", FixPNPM)
}

