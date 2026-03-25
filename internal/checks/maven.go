package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckMaven() model.Result {
	cmd := exec.Command("mvn", "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "maven",
			Status:  model.Missing,
			Message: "Maven not installed or not in PATH",
			Fix:     "Install Maven: brew install maven (macOS) or sudo apt install maven (Linux)",
		}
	}

	version := strings.SplitN(string(output), "\n", 2)[0]
	return model.Result{
		Name:    "maven",
		Status:  model.OK,
		Message: version,
	}
}

func init() {
	Register("maven", CheckMaven)
}
