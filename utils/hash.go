package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(s string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(s)); err != nil {
		return "", err
	}

	hashed := hex.EncodeToString(hash.Sum(nil))
	return hashed, nil
}
