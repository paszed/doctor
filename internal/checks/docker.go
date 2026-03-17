package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register(CheckDocker)
}

func CheckDocker() model.Result {

	cmd := exec.Command("docker", "--version")

	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "docker",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "docker",
		Status:  model.OK,
		Message: version,
	}
}
