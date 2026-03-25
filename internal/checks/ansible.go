package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckAnsible() model.Result {
	cmd := exec.Command("ansible", "--version")
	out, err := cmd.Output()

	if err != nil {
		return model.Result{
			Name:    "ansible",
			Status:  model.Missing,
			Message: "Ansible not installed or not in PATH",
			Fix:     "Install Ansible: brew install ansible (macOS) or pip install ansible",
		}
	}

	versionLine := strings.SplitN(string(out), "\n", 2)[0]

	return model.Result{
		Name:    "ansible",
		Status:  model.OK,
		Message: versionLine,
	}
}

func init() {
	Register("ansible", CheckAnsible)
}
