package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register(CheckPython)
}

func CheckPython() model.Result {

	path, err := exec.LookPath("python3")
	if err != nil {
		return model.Result{
			Name:    "python3",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	out, err := exec.Command("python3", "--version").Output()
	if err != nil {
		return model.Result{
			Name:    "python3",
			Status:  model.Missing,
			Path:    path,
			Message: "python command failed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "python3",
		Status:  model.OK,
		Version: version,
		Path:    path,
	}
}
