package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateResetToken() string{
	min := 100000
	max := 999999
	time := time.Now().UnixMilli()
	rand.Seed(time)
	fmt.Println(time)
	return strconv.Itoa(min + rand.Intn(max-min))
}