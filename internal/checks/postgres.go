package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckPostgres() model.Result {
	cmd := exec.Command("psql", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "postgres",
			Status:  model.Missing,
			Message: "PostgreSQL not installed or not in PATH",
			Fix:     "Install PostgreSQL and ensure `psql` is in PATH",
		}
	}

	version := strings.TrimSpace(string(out))
	return model.Result{
		Name:    "postgres",
		Status:  model.OK,
		Message: version,
	}
}

func init() {
	Register("postgres", CheckPostgres)
}
