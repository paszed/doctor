package cli

import (
	"fmt"
	"os"
)

func RunInit(args []string) {
	// Check if file already exists
	if _, err := os.Stat("doctor.yaml"); err == nil {
		fmt.Println("doctor.yaml already exists")
		return
	}

	content := `ports:
  - 3000
  - 5432
`

	err := os.WriteFile("doctor.yaml", []byte(content), 0644)
	if err != nil {
		fmt.Println("failed to create doctor.yaml:", err)
		os.Exit(1)
	}

	fmt.Println("✓ created doctor.yaml")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  doctor diagnose")
	fmt.Println("  doctor fix")
}
