package fix

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

func CheckOracle() model.Result {
	cmd := exec.Command("sqlplus", "-V")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return model.Result{
			Name:    "oracle",
			Status:  model.Missing,
			Message: "Oracle SQL CLI not installed or not in PATH",
			Fix:     "Install Oracle Instant Client and ensure `sqlplus` is in PATH",
		}
	}
	return model.Result{
		Name:    "oracle",
		Status:  model.OK,
		Message: strings.TrimSpace(string(output)),
	}
}

func init() {
	Register("oracle", CheckOracle)
}
