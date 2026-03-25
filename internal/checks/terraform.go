package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckTerraform() model.Result {
	cmd := exec.Command("terraform", "version")
	if err := cmd.Run(); err != nil {
		return model.Result{
			Name:    "terraform",
			Status:  model.Missing,
			Message: "Terraform not installed or not in PATH",
			Fix:     "Install Terraform from https://www.terraform.io/downloads.html",
		}
	}
	return model.Result{
		Name:    "terraform",
		Status:  model.OK,
		Message: "✓ Terraform installed",
	}
}

func init() {
	Register("terraform", CheckTerraform)
}
