package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

//author: gurbanli

func GeneratePasswordHash(password string) string{
	hashInstance := sha512.New()
	hashInstance.Write([]byte(password))
	return hex.EncodeToString(hashInstance.Sum(nil))
}