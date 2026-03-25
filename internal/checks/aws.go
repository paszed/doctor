package checks

import (
	"github.com/paszed/doctor/internal/model"
	"os/exec"
)

func CheckAWS() model.Result {
	_, err := exec.LookPath("aws")
	if err != nil {
		return model.Result{
			Name:    "aws",
			Status:  model.Missing,
			Message: "AWS CLI not installed or not in PATH",
			Fix:     "Install AWS CLI: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html",
		}
	}
	return model.Result{
		Name:    "aws",
		Status:  model.OK,
		Message: "AWS CLI installed",
	}
}

func init() {
	Register("aws", CheckAWS)
}
