package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckJava() model.Result {
	cmd := exec.Command("java", "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "java",
			Status:  model.Missing,
			Message: "Java not installed or not in PATH",
			Fix:     "Install JDK and ensure `java` is in PATH",
		}
	}

	versionInfo := strings.Split(string(out), "\n")[0] // first line usually contains version
	versionInfo = strings.TrimSpace(versionInfo)
	return model.Result{
		Name:    "java",
		Status:  model.OK,
		Message: versionInfo,
	}
}

func init() {
	Register("java", CheckJava)
}
