package bc

import (
	"crypto/sha256"
	"fmt"
)

func GetCarHash(car CarData) []byte {
	h := sha256.New()
	h.Write([]byte(car.ToString()))
	return h.Sum(nil)
}

func GetHash(text string) []byte {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	return h.Sum(nil)
}

func PrintBytes(bt []byte) string {
	return fmt.Sprintf("%x", bt)
}
