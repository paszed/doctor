package fix

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckPipenv() model.Result {
	res := model.Result{Name: "pipenv"}

	out, err := exec.Command("pipenv", "--version").Output()
	if err != nil {
		res.Status = model.Missing
		res.Message = "pipenv not installed or not in PATH"
		res.Fix = "Install pipenv via: pipx install pipenv or brew install pipenv"
		return res
	}
	res.Status = model.OK
	res.Message = strings.TrimSpace(string(out))
	return res
}

func init() {
	Register("pipenv", FixPipenv)
}
