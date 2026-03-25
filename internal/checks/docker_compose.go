package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckDockerCompose() model.Result {
	cmd := exec.Command("docker-compose", "--version")
	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "docker-compose",
			Status:  model.Missing,
			Message: "Docker Compose not installed or not in PATH",
			Fix:     "Install Docker Compose (`brew install docker-compose` on macOS or see https://docs.docker.com/compose/install/)",
		}
	}
	return model.Result{
		Name:    "docker-compose",
		Status:  model.OK,
		Message: string(out),
	}
}

func init() {
	Register("docker-compose", CheckDockerCompose)
}
