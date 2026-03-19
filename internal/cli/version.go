package cli

import "fmt"

const version = "0.1.0"

func RunVersion(args []string) {
	fmt.Println("Doctor version", version)
}
