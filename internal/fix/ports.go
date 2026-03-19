package fix

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const timeout = 2 * time.Second

func FixPorts(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("ports fix requires at least one port")
	}

	for _, port := range args {
		fmt.Printf("→ checking port %s...\n", port)

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, "lsof", "-ti", ":"+port)
		out, _ := cmd.Output()

		// timeout case
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("! timeout checking port %s\n", port)
			continue
		}

		// no process using port
		if len(out) == 0 {
			fmt.Printf("✓ port %s is already free\n", port)
			continue
		}

		pids := strings.Fields(string(out))

		for _, pid := range pids {
			// get process name
			nameCmd := exec.Command("ps", "-p", pid, "-o", "comm=")
			nameOut, _ := nameCmd.Output()
			name := strings.TrimSpace(string(nameOut))

			if name == "" {
				name = "unknown"
			} else {
				parts := strings.Split(name, "/")
				name = parts[len(parts)-1]
			}

			fmt.Printf("→ killing %s (pid %s) on port %s...\n", name, pid, port)

			// try graceful kill
			killCmd := exec.Command("kill", pid)
			if err := killCmd.Run(); err == nil {
				fmt.Printf("✓ killed %s (pid %s) on port %s\n", name, pid, port)
				continue
			}

			// fallback to force kill
			killCmd = exec.Command("kill", "-9", pid)
			if err := killCmd.Run(); err != nil {
				fmt.Printf("✗ failed to kill %s (pid %s)\n", name, pid)
				continue
			}

			fmt.Printf("✓ force killed %s (pid %s) on port %s\n", name, pid, port)
		}

		fmt.Println()
	}

	return nil
}

func init() {
	Register("ports", func() error {
		return fmt.Errorf("ports fix requires arguments (e.g. doctor fix ports 3000)")
	})
}
