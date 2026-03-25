package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckMySQL verifies if the `mysql` client is installed and accessible.
func CheckMySQL() model.Result {
	_, err := exec.LookPath("mysql")
	if err != nil {
		return model.Result{
			Name:    "mysql",
			Status:  model.Missing,
			Message: "MySQL client not installed or not in PATH",
			Fix:     "Install MySQL client via Homebrew: brew install mysql",
		}
	}

	return model.Result{
		Name:    "mysql",
		Status:  model.OK,
		Message: "MySQL client installed",
	}
}

func init() {
	Register("mysql", CheckMySQL)
}
