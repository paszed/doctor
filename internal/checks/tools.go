package checks

import (
	"os/exec"
	"strings"
	"github.com/paszed/doctor/internal/ui"
)

func ToolVersion(name string) {
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
		ui.Fail(name, "not installed")
		MissingTools++
		return
	}

	version := strings.Split(string(out), "\n")[0]
	version = strings.Replace(version, name+" version ", "", 1)

	InstalledTools++
	ui.OK(name, version)
}
