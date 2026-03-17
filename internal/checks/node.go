package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register(CheckNode)
}

func CheckNode() model.Result {

	path, err := exec.LookPath("node")
	if err != nil {
		return model.Result{
			Name:    "node",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	out, err := exec.Command("node", "--version").Output()
	if err != nil {
		return model.Result{
			Name:    "node",
			Status:  model.Missing,
			Path:    path,
			Message: "node command failed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "node",
		Status:  model.OK,
		Version: version,
		Path:    path,
	}
}
