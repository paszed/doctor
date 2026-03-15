package checks

import (
	"fmt"
	"os/exec"
	"strings"
)

func DockerCheck() {

	// check docker CLI
	cmd := exec.Command("docker", "--version")
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("%-12s ✗ not installed\n", "docker")
		return
	}

	// parse version
	version := strings.SplitN(string(out), "\n", 2)[0]
	version = strings.TrimPrefix(version, "Docker version ")
	version = strings.Split(version, ",")[0]

	fmt.Printf("%-12s ✓ %s\n", "docker", version)

	// check docker daemon
	cmd = exec.Command("docker", "info")
	err = cmd.Run()

	if err != nil {
		fmt.Printf("%-12s ✗ not running\n", "daemon")
		return
	}

	fmt.Printf("%-12s ✓ running\n", "daemon")
}
