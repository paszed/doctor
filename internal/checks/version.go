package checks

import "fmt"

var Version = "0.1.0"

func DoctorVersion() {
	fmt.Printf("doctor v%s\n", Version)
}
