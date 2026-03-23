package checks

// RegisterAll registers all available CLI checks that take no arguments
func RegisterAll() {
	Register("docker", CheckDocker)
	Register("git", CheckGit)
	Register("go", CheckGo)
	Register("node", CheckNode)
	Register("python3", CheckPython)
	Register("sleep", CheckSleep)
}
