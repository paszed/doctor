package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckGH checks if GitHub CLI is installed and available in PATH
func CheckGH() model.Result {
	cmd := exec.Command("gh", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "gh",
			Status:  model.Missing,
			Message: "GitHub CLI not installed or not in PATH",
			Fix:     "Install GitHub CLI: https://cli.github.com/",
		}
	}

	return model.Result{
		Name:    "gh",
		Status:  model.OK,
		Message: string(output),
	}
}

func init() {
	Register("gh", CheckGH)
}
