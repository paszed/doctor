package checks

import (
	"github.com/paszed/doctor/internal/model"
	"os/exec"
)

func CheckMSSQL() model.Result {
	_, err := exec.LookPath("sqlcmd")
	if err != nil {
		return model.Result{
			Name:    "mssql",
			Status:  model.Missing,
			Message: "MSSQL CLI (sqlcmd) not installed or not in PATH",
			Fix:     "Install sqlcmd: https://learn.microsoft.com/en-us/sql/tools/sqlcmd-utility",
		}
	}
	return model.Result{
		Name:    "mssql",
		Status:  model.OK,
		Message: "MSSQL CLI installed",
	}
}

func init() {
	Register("mssql", CheckMSSQL)
}
