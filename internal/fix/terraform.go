package fix

import (
	"fmt"
)

func FixTerraform(args []string) error {
	// Auto-installing Terraform is platform-dependent, so we just give guidance
	fmt.Println("→ terraform not found. Please install it manually:")
	fmt.Println("   https://www.terraform.io/downloads.html")
	return nil
}

func init() {
	Register("terraform", FixTerraform)
}
