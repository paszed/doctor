package model

type Status int

const (
	OK Status = iota
	Missing
	Warning
)

type Result struct {
	Name    string `json:"name"`
	Status  Status `json:"status"`
	Version string `json:"version,omitempty"`
	Path    string `json:"path,omitempty"`
	Message string `json:"message,omitempty"`
}
