package checks

import (
	"time"

	"github.com/paszed/doctor/internal/model"
)

func CheckSleep() model.Result {
	time.Sleep(5 * time.Second) // intentionally longer than timeout

	return model.Result{
		Name:    "sleep",
		Status:  model.OK,
		Message: "should not reach here",
	}
}

func init() {
	Register("sleep", CheckSleep)
}
