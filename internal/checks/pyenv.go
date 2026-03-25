package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckPyenv checks if pyenv is installed
func CheckPyenv() model.Result {
	_, err := exec.LookPath("pyenv")
	if err != nil {
		return model.Result{
			Name:    "pyenv",
			Status:  model.Missing,
			Message: "pyenv not installed or not in PATH",
			Fix:     "Install pyenv: brew install pyenv",
		}
	}

	return model.Result{
		Name:    "pyenv",
		Status:  model.OK,
		Message: "pyenv installed",
	}
}

func init() {
	Register("pyenv", CheckPyenv)
}
