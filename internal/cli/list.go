package cli

import (
	"fmt"

	"github.com/paszed/doctor/internal/checks"
)

func RunList() {

	names := checks.Names()

	fmt.Printf("AVAILABLE CHECKS (%d) \n", len(names))
	fmt.Println("----------------")

	for _, name := range names {
		fmt.Println(name)
	}
}
