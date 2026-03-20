package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckGit() model.Result {

	cmd := exec.Command("git", "--version")

	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "git",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "git",
		Status:  model.OK,
		Message: version,
	}
}
