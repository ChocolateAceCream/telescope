package utils

import "crypto/sha256"

func Sha256(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return s
}

// generic function to convert any slice to []interface{}
func SliceToInterfaceSlice[T any](slice []T) []interface{} {
	i := make([]interface{}, len(slice))
	for k, v := range slice {
		i[k] = v
	}
	return i
}
