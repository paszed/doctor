package checks

import (
	"github.com/paszed/doctor/internal/model"
	"os/exec"
)

func CheckAzure() model.Result {
	_, err := exec.LookPath("az")
	if err != nil {
		return model.Result{
			Name:    "az",
			Status:  model.Missing,
			Message: "Azure CLI not installed or not in PATH",
			Fix:     "Install Azure CLI: https://learn.microsoft.com/en-us/cli/azure/install-azure-cli",
		}
	}
	return model.Result{
		Name:    "az",
		Status:  model.OK,
		Message: "Azure CLI installed",
	}
}

func init() {
	Register("az", CheckAzure)
}
