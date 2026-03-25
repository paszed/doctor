package checks

import (
	"fmt"
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckHelm() model.Result {
	cmd := exec.Command("helm", "version", "--short")
	output, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "helm",
			Status:  model.Missing,
			Message: "Helm not installed or not in PATH",
			Fix:     "Install Helm from https://helm.sh/docs/intro/install/",
		}
	}

	return model.Result{
		Name:    "helm",
		Status:  model.OK,
		Message: fmt.Sprintf("%s", output),
	}
}

func init() {
	Register("helm", CheckHelm)
}
