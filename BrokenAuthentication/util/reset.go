package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateResetToken() string{
	time := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(time)
	return strconv.Itoa(RangeRandom(100000, 999999, time))
}

func RangeRandom(min int, max int, seed int64)int{
	rand.Seed(seed)
	return min + rand.Intn(max-min)
}
