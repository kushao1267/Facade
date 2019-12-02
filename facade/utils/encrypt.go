package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash get md5 hash from a string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
