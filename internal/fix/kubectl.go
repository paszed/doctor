package fix

import (
	"fmt"
)

func FixKubectl(args []string) error {
	fmt.Println("[FIX] kubectl")
	fmt.Println("→ kubectl not found. Please install it manually:")
	fmt.Println("  https://kubernetes.io/docs/tasks/tools/")
	return nil
}

func init() {
	Register("kubectl", FixKubectl)
}
