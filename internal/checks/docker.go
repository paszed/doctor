package checks

import (
	"context"
	"os/exec"
	"time"

	"github.com/paszed/doctor/internal/model"
)

const dockerTimeout = 2 * time.Second

func CheckDocker() model.Result {
	ctx, cancel := context.WithTimeout(context.Background(), dockerTimeout)
	defer cancel()

	// check if docker exists
	path, err := exec.LookPath("docker")
	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  1,
			Message: "docker not installed",
			Fix:     "install docker",
		}
	}

	// check if docker daemon is running
	cmd := exec.CommandContext(ctx, "docker", "info")
	if err := cmd.Run(); err != nil {
		return model.Result{
			Name:    "docker",
			Status:  1,
			Path:    path,
			Message: "Docker installed but NOT running",
			Fix:     "doctor fix docker",
		}
	}

	// get version
	versionCmd := exec.CommandContext(ctx, "docker", "--version")
	out, err := versionCmd.Output()
	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  0,
			Path:    path,
			Message: "docker installed (version unknown)",
		}
	}

	return model.Result{
		Name:    "docker",
		Status:  0,
		Path:    path,
		Message: string(out),
	}
}

func init() {
	Register("docker", CheckDocker)
}
