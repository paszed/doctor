package model

type CheckResult struct {
	Name    string
	Status  bool
	Version string
	Path    string
	Message string
}
