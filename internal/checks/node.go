package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func init() {
	Register("node", CheckNode)
}

func CheckNode() model.Result {

	cmd := exec.Command("node", "--version")

	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "node",
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))

	return model.Result{
		Name:    "node",
		Status:  model.OK,
		Message: version,
	}
}
