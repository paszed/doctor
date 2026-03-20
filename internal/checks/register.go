package checks

func RegisterAll() {
	Register("docker", CheckDocker)
	Register("git", CheckGit)
	Register("go", CheckGo)
	Register("node", CheckNode)
	Register("python3", CheckPython)
	Register("sleep", CheckSleep)
}
