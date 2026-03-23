package checks

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckDocker() model.Result {
	// Check if docker exists
	_, err := exec.LookPath("docker")
	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  model.Missing,
			Message: "Docker is not installed",
			Fix:     "Install Docker Desktop",
		}
	}

	// Check if daemon is running + get version
	cmd := exec.Command("docker", "version", "--format", "{{.Server.Version}}")
	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  model.Missing,
			Message: "Docker installed but NOT running",
			Fix:     "doctor fix docker",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "docker",
		Status:  model.OK,
		Message: fmt.Sprintf("Docker version %s", version),
	}
}
