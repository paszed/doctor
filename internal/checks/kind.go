package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckKind() model.Result {
	_, err := exec.LookPath("kind")
	if err != nil {
		return model.Result{
			Name:    "kind",
			Status:  model.Missing,
			Message: "Kind not installed or not in PATH",
			Fix:     "Install Kind: https://kind.sigs.k8s.io/docs/user/quick-start/",
		}
	}
	return model.Result{
		Name:    "kind",
		Status:  model.OK,
		Message: "Kind installed",
	}
}

func init() {
	Register("kind", CheckKind)
}
