package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func(r *Repository) InitializeOTPCacheRepository(){
	r.OTPCache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // OTP cache db
	})
	r.Ctx = context.Background()
}

func(r *Repository) GetOTP(username string) string{
	value, err := r.OTPCache.Get(r.Ctx, username).Result()
	if err != nil{
		return ""
	}
	return value
}

func(r *Repository) SetOTP(username string, otpCode string){
	err := r.OTPCache.Set(r.Ctx, username, otpCode,0).Err()
	if err != nil{
		log.Fatal(err)
	}
}