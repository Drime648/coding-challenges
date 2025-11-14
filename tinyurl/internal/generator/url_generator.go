package generator

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"fmt"
)

//plan: sha256 the url, then get the base58 of it



func generateHash(input string) []byte {
	h := sha256.New()
	h.Write([]byte(input))
	return h.Sum(nil)
}

func base58Encode(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", fmt.Errorf("Unable to Encode the hash: %s | Error: %v",bytes, err)
	}
	return string(encoded), nil
}

func generateUrl(originalUrl string) (string, error) {
	hash := generateHash(originalUrl)
	encoded, err := base58Encode(hash)
	if err != nil {
		return "", err
	}
	return encoded[:8], nil //get first 8 chars only
}
