package util

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

func Hash(password string) string{
	time.Sleep(1 * time.Second)
	hashInstance := sha512.New()
	hashInstance.Write([]byte(password))
	return hex.EncodeToString(hashInstance.Sum(nil))
}