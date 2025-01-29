package utils

import "crypto/sha256"

func Sha256(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return s
}
