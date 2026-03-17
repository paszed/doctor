package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register(CheckGit)
}

func CheckGit() model.Result {

	path, err := exec.LookPath("git")
	if err != nil {
		return model.Result{
			Name:    "git",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	out, err := exec.Command("git", "--version").Output()
	if err != nil {
		return model.Result{
			Name:    "git",
			Status:  model.Missing,
			Path:    path,
			Message: "git command failed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "git",
		Status:  model.OK,
		Version: version,
		Path:    path,
	}
}
