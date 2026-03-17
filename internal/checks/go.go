package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register(CheckGo)
}

func CheckGo() model.Result {

	cmd := exec.Command("go", "version")

	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "go",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "go",
		Status:  model.OK,
		Message: version,
	}
}
