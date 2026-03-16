package model

type Status int

const (
	OK Status = iota
	Missing
	Warning
)

type Result struct {
	Name    string
	Status  Status
	Message string
}
