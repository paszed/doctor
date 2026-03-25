package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckPip() model.Result {
	res := model.Result{Name: "pip"}

	out, err := exec.Command("pip3", "--version").Output()
	if err != nil {
		res.Status = model.Missing
		res.Message = "pip not installed or not in PATH"
		res.Fix = "Install pip via python3 -m ensurepip or use Homebrew"
		return res
	}
	res.Status = model.OK
	res.Message = strings.TrimSpace(string(out))
	return res
}

func init() {
	Register("pip", CheckPip)
}
