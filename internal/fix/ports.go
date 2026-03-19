package fix

import (
	"fmt"
	"os/exec"
	"strings"
)

func init() {
	Register("ports", FixPorts)
}

func FixPorts(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: doctor fix ports <port>")
	}

	port := args[0]

	fmt.Printf("→ checking port %s...\n", port)

	// Find process using port
	cmd := exec.Command("lsof", "-ti", ":"+port)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		fmt.Printf("✓ port %s is already free\n", port)
		return nil
	}

	pids := strings.Fields(string(out))

	for _, pid := range pids {
		fmt.Printf("→ killing process %s on port %s...\n", pid, port)

		killCmd := exec.Command("kill", "-9", pid)
		if err := killCmd.Run(); err != nil {
			fmt.Printf("✗ failed to kill %s\n", pid)
			continue
		}

		fmt.Printf("✓ killed %s\n", pid)
	}

	return nil
}
