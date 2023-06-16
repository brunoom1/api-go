package helpers

import "os"

func GetSecurityKey() string {
	securityKey := os.Getenv("SECURITY_KEY")

	if securityKey == "" {
		securityKey = "myscriptteste"
	}

	return securityKey
}
