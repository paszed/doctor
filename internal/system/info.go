package system

import (
	"fmt"
	"runtime"
)

func Info() {

	fmt.Println("SYSTEM")
	fmt.Println("--------------------")

	fmt.Printf("%-12s %s\n", "OS", runtime.GOOS)
	fmt.Printf("%-12s %s\n", "ARCH", runtime.GOARCH)
	fmt.Println()

}
