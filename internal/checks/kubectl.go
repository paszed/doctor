package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckKubectl() model.Result {
	cmd := exec.Command("kubectl", "version", "--client", "--short")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "kubectl",
			Status:  model.Missing,
			Message: "kubectl not installed or not in PATH",
			Fix:     "Install kubectl from https://kubernetes.io/docs/tasks/tools/",
		}
	}
	return model.Result{
		Name:    "kubectl",
		Status:  model.OK,
		Message: string(output),
	}
}

func init() {
	Register("kubectl", CheckKubectl)
}
