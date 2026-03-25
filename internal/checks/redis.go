package checks

import (
	"os/exec"

	"github.com/paszed/doctor/internal/model"
)

// CheckRedis checks if the `redis-cli` tool is installed.
func CheckRedis() model.Result {
	_, err := exec.LookPath("redis-cli")
	if err != nil {
		return model.Result{
			Name:    "redis",
			Status:  model.Missing,
			Message: "Redis CLI not installed or not in PATH",
			Fix:     "Install Redis CLI via Homebrew: brew install redis",
		}
	}

	return model.Result{
		Name:    "redis",
		Status:  model.OK,
		Message: "Redis CLI installed",
	}
}

func init() {
	Register("redis", CheckRedis) // ✅ Only the CheckFunc
}
