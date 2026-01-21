package internal

import (
	"fmt"
	"os"
)

// CheckPolicy exists to verify if the Terraform-managed policy file is present.
func CheckPolicy(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("SECURITY ALERT: Policy file %s is missing", filePath)
	}
	fmt.Printf("✅ Infrastructure Verified: %s found. \n", filePath)
	return nil
}
