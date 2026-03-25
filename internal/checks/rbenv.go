package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckRbenv checks if rbenv is installed
func CheckRbenv() model.Result {
	_, err := exec.LookPath("rbenv")
	if err != nil {
		return model.Result{
			Name:    "rbenv",
			Status:  model.Missing,
			Message: "rbenv not installed or not in PATH",
			Fix:     "Install rbenv: brew install rbenv",
		}
	}

	return model.Result{
		Name:    "rbenv",
		Status:  model.OK,
		Message: "rbenv installed",
	}
}

func init() {
	Register("rbenv", CheckRbenv)
}
