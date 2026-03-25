package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckKubectl() model.Result {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		return model.Result{
			Name:    "kubectl",
			Status:  model.Missing,
			Message: "kubectl not installed or not in PATH",
			Fix:     "Install kubectl from https://kubernetes.io/docs/tasks/tools/",
		}
	}

	// Get version
	out, err := exec.Command(path, "version", "--client", "--short").Output()
	msg := strings.TrimSpace(string(out))
	if err != nil || msg == "" {
		msg = "kubectl installed but could not detect version"
	}

	return model.Result{
		Name:    "kubectl",
		Status:  model.OK,
		Message: msg,
	}
}

func init() {
	Register("kubectl", CheckKubectl)
}
