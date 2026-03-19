package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckDocker() model.Result {

	path, err := exec.LookPath("docker")

	if err != nil {
		return model.Result{
			Name:   "docker",
			Status: model.Missing,
		}
	}

	out, err := exec.Command("docker", "--version").Output()

	if err != nil {
		return model.Result{
			Name:   "docker",
			Status: model.Warning,
			Path:   path,
		}
	}

	return model.Result{
		Name:    "docker",
		Status:  model.OK,
		Path:    path,
		Message: strings.TrimSpace(string(out)),
	}
}

func init() {
	Register("docker", CheckDocker)
}
