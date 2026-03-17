package cli

import (
	"fmt"
	"os"
)

func RunFix() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: doctor fix <port>")
		return
	}

	port := os.Args[2]
	fmt.Println("Attempting to free port", port)
}
