package generator

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"fmt"
	"math/big"
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

func GenerateUrl(originalUrl string) (string, error) {
	hash := generateHash(originalUrl)
	generatedNumber := new(big.Int).SetBytes(hash).Uint64()
	encoded, err := base58Encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}
	return encoded[:8], nil //get first 8 chars only
}
