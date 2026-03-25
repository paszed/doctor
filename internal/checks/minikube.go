package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

func CheckMinikube() model.Result {
	_, err := exec.LookPath("minikube")
	if err != nil {
		return model.Result{
			Name:    "minikube",
			Status:  model.Missing,
			Message: "Minikube not installed or not in PATH",
			Fix:     "Install Minikube: https://minikube.sigs.k8s.io/docs/start/",
		}
	}
	return model.Result{
		Name:    "minikube",
		Status:  model.OK,
		Message: "Minikube installed",
	}
}

func init() {
	Register("minikube", CheckMinikube)
}
