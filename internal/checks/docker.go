package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func DockerCheck() model.Result {

	cmd := exec.Command("docker", "--version")
	out, err := cmd.Output()

	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.Split(string(out), "\n")[0]
	version = strings.TrimPrefix(version, "Docker version ")
	version = strings.Split(version, ",")[0]

	return model.Result{
		Name:    "docker",
		Status:  model.OK,
		Message: version,
	}
}
