package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckPoetry() model.Result {
	res := model.Result{Name: "poetry"}
	out, err := exec.Command("poetry", "--version").Output()
	if err != nil {
		res.Status = model.Missing
		res.Message = "Poetry not installed or not in PATH"
		res.Fix = "Install Poetry via: brew install poetry or pipx install poetry"
		return res
	}
	res.Status = model.OK
	res.Message = strings.TrimSpace(string(out))
	return res
}

func init() {
	Register("poetry", CheckPoetry)
}

