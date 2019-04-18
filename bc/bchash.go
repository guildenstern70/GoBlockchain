package bc

import (
	"crypto/sha256"
	"fmt"
)

func GetHash(text string) []byte {
	h := sha256.New()
	h.Write([]byte(text))
	return h.Sum(nil)
}

func PrintBytes(bt []byte) string {
	return fmt.Sprintf("%x", bt)
}
