package checks

import (
	"github.com/paszed/doctor/internal/model"
	"os/exec"
)

func CheckGCloud() model.Result {
	_, err := exec.LookPath("gcloud")
	if err != nil {
		return model.Result{
			Name:    "gcloud",
			Status:  model.Missing,
			Message: "Google Cloud CLI not installed or not in PATH",
			Fix:     "Install gcloud CLI: https://cloud.google.com/sdk/docs/install",
		}
	}
	return model.Result{
		Name:    "gcloud",
		Status:  model.OK,
		Message: "Google Cloud CLI installed",
	}
}

func init() {
	Register("gcloud", CheckGCloud)
}
