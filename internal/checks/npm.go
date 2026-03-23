package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckNPM() model.Result {
	cmd := exec.Command("npm", "--version")
	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "npm",
			Status:  model.Missing,
			Message: "npm not installed or not in PATH",
			Fix:     "Install Node.js which includes npm",
		}
	}
	version := strings.TrimSpace(string(out))
	return model.Result{
		Name:    "npm",
		Status:  model.OK,
		Message: "v" + version,
	}
}

func CheckPNPM() model.Result {
	cmd := exec.Command("pnpm", "--version")
	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "pnpm",
			Status:  model.Missing,
			Message: "pnpm not installed or not in PATH",
			Fix:     "Install pnpm via npm: npm install -g pnpm",
		}
	}
	version := strings.TrimSpace(string(out))
	return model.Result{
		Name:    "pnpm",
		Status:  model.OK,
		Message: "v" + version,
	}
}

func CheckYarn() model.Result {
	cmd := exec.Command("yarn", "--version")
	out, err := cmd.Output()
	if err != nil {
		return model.Result{
			Name:    "yarn",
			Status:  model.Missing,
			Message: "yarn not installed or not in PATH",
			Fix:     "Install yarn via npm: npm install -g yarn",
		}
	}
	version := strings.TrimSpace(string(out))
	return model.Result{
		Name:    "yarn",
		Status:  model.OK,
		Message: "v" + version,
	}
}

func init() {
	Register("npm", CheckNPM)
	Register("pnpm", CheckPNPM)
	Register("yarn", CheckYarn)
}
