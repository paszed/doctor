package checks

import (
	"github.com/paszed/doctor/internal/model"
	"os/exec"
)

func CheckSQLite() model.Result {
	_, err := exec.LookPath("sqlite3")
	if err != nil {
		return model.Result{
			Name:    "sqlite3",
			Status:  model.Missing,
			Message: "SQLite CLI not installed or not in PATH",
			Fix:     "Install SQLite: brew install sqlite (macOS) or sudo apt install sqlite3 (Linux)",
		}
	}
	return model.Result{
		Name:    "sqlite3",
		Status:  model.OK,
		Message: "SQLite CLI installed",
	}
}

func init() {
	Register("sqlite3", CheckSQLite)
}
