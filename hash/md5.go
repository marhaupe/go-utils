package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 hashes input via crypto/md5 hashing implementation
func Md5(in string) string {
	hasher := md5.New()
	hasher.Write([]byte(in))
	return hex.EncodeToString(hasher.Sum(nil))
}
