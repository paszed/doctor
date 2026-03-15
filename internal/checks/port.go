package checks

import (
	"os/exec"
	"strings"
)

func PortProcess(port string) (bool, string) {

	cmd := exec.Command("lsof", "-i:"+port)
	out, err := cmd.Output()

	if err != nil {
		return false, ""
	}

	lines := strings.Split(string(out), "\n")

	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		process := fields[0]

		return true, process
	}

	return false, ""
}
