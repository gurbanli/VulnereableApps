package util

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string{
	min := 1000
	max := 9999
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min))
}
