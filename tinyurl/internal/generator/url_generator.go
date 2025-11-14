package generator

import (
	"crypto/sha256"
)

//plan: sha256 the url, then get the base58 of it



func generateHash(input string) []byte {
	return sha256.Sum256([]byte(input))
}
