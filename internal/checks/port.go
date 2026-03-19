package checks

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/paszed/doctor/internal/model"
)

const portTimeout = 2 * time.Second

func CheckPort(port string) model.Result {
	ctx, cancel := context.WithTimeout(context.Background(), portTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lsof", "-ti", ":"+port)
	out, _ := cmd.Output()

	// timeout case
	if ctx.Err() == context.DeadlineExceeded {
		return model.Result{
			Name:    port,
			Status:  1,
			Message: "timeout while checking port",
		}
	}

	// port is free
	if len(out) == 0 {
		return model.Result{
			Name:    port,
			Status:  0,
			Message: "free",
		}
	}

	// get first PID
	pids := strings.Fields(string(out))
	pid := pids[0]

	// get process name
	nameCmd := exec.Command("ps", "-p", pid, "-o", "comm=")
	nameOut, _ := nameCmd.Output()
	processName := strings.TrimSpace(string(nameOut))

	if processName == "" {
		processName = "unknown"
	} else {
		parts := strings.Split(processName, "/")
		processName = parts[len(parts)-1]
	}

	return model.Result{
		Name:    port,
		Status:  1,
		Message: fmt.Sprintf("in use (%s)", processName),
		Fix:     fmt.Sprintf("doctor fix ports %s", port),
	}
}

func init() {
	Register("port", func() model.Result {
		// default port check (optional, can extend later)
		return CheckPort("3000")
	})
}
