package checks

import (
	"os/exec"
	"strings"

	"github.com/paszed/doctor/internal/model"
)

// CheckMongo verifies if the MongoDB CLI (mongo) is installed.
func CheckMongo() model.Result {
	res := model.Result{Name: "mongo"}

	cmd := exec.Command("mongo", "--version")
	out, err := cmd.Output()
	if err != nil {
		res.Status = model.Missing
		res.Message = "MongoDB CLI not installed or not in PATH"
		res.Fix = "Install MongoDB CLI: brew install mongodb/brew/mongodb-community-shell (macOS) or sudo apt install mongodb-clients (Linux)"
		return res
	}

	version := strings.TrimSpace(string(out))
	res.Status = model.OK
	res.Message = version
	return res
}

func init() {
	Register("mongo", CheckMongo)
}
