package cli

import (
	"fmt"
	"os"

	"github.com/paszed/doctor/internal/checks"
)

func RunPort() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: doctor port <port>")
		return
	}

	port := os.Args[2]
	checks.PortProcess(port)
}
