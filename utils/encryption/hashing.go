package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/openpgp/errors"
)

func GenerateHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	bs := hex.EncodeToString(h.Sum(nil))
	return bs
}

func CheckPassword(hash string, pass string) error {
	h := sha256.New()
	h.Write([]byte(pass))
	bs := hex.EncodeToString(h.Sum(nil))
	if bs != hash {
		return errors.InvalidArgumentError("Hashes do not match")
	}
	return nil
}
