package helpers

import (
	"crypto/sha256"
)

func Encrypt(pass string) []byte {
	sh := sha256.New()
	sh.Write([]byte(pass))
	return sh.Sum(nil)
}

func EncryptVerify(pass string, verify string) bool {
	sh := sha256.New()
	sh.Write([]byte(pass))
	hash := sh.Sum(nil)

	return string(hash) == verify
}
