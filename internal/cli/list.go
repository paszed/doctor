package cli

import (
	"fmt"

	"github.com/paszed/doctor/internal/checks"
)

func RunList() {

	fmt.Println("AVAILABLE CHECKS")
	fmt.Println("----------------")

	for _, name := range checks.Names() {
		fmt.Println(name)
	}
}
