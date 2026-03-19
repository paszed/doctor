package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register("python3", CheckPython)
}

func CheckPython() model.Result {

	cmd := exec.Command("python3", "--version")

	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "python3",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "python3",
		Status:  model.OK,
		Message: version,
	}
}
