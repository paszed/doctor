package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

var DevTools = []string{
	"git",
	"node",
	"python3",
}

func ToolVersion(name string) model.Result {

	args := []string{"--version", "-v", "-V", "version"}

	var out []byte
	var err error

	for _, arg := range args {

		cmd := exec.Command(name, arg)
		out, err = cmd.Output()

		if err == nil {
			break
		}
	}

	if err != nil {
		return model.Result{
			Name:    name,
			Status:  model.Missing,
			Message: "not installed",
		}
	}

	version := strings.TrimSpace(string(out))
	version = strings.Split(version, "\n")[0]
	version = strings.ReplaceAll(version, name+" version ", "")

	return model.Result{
		Name:    name,
		Status:  model.OK,
		Message: version,
	}
}

func GitCheck() model.Result {
	return ToolVersion("git")
}

func NodeCheck() model.Result {
	return ToolVersion("node")
}

func PythonCheck() model.Result {
	return ToolVersion("python3")
}
