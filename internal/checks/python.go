package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

// CheckPython verifies that Python 3 is installed and accessible.
func CheckPython() model.Result {
	res := model.Result{Name: "python3"}

	out, err := exec.Command("python3", "--version").Output()
	if err != nil {
		res.Status = model.Missing
		res.Message = "Python 3 not installed or not in PATH"
		res.Fix = "Install Python 3 from https://www.python.org/downloads/"
		return res
	}

	res.Status = model.OK
	res.Message = strings.TrimSpace(string(out))
	return res
}

// init registers the Python check in the global registry
func init() {
	Register("python3", CheckPython)
}

