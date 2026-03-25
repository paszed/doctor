package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckNVM checks if nvm (Node Version Manager) is installed
func CheckNVM() model.Result {
	_, err := exec.LookPath("nvm")
	if err != nil {
		return model.Result{
			Name:    "nvm",
			Status:  model.Missing,
			Message: "nvm not installed or not in PATH",
			Fix:     "Install nvm: https://github.com/nvm-sh/nvm#installing-and-updating",
		}
	}

	return model.Result{
		Name:    "nvm",
		Status:  model.OK,
		Message: "nvm installed",
	}
}

func init() {
	Register("nvm", CheckNVM)
}
