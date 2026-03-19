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
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Fix     string `json:"fix,omitempty"`
}
