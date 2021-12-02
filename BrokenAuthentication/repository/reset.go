package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func(r *Repository) InitializeResetTokenCacheRepository(){
	r.ResetTokenCache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       3,  // OTP cache db
	})
	r.Ctx = context.Background()
}

func(r *Repository) GetResetToken(username string) string{
	value, err := r.ResetTokenCache.Get(r.Ctx, username).Result()
	if err != nil{
		return ""
	}
	return value
}

func(r *Repository) SetResetToken(username string, resetToken string){
	err := r.ResetTokenCache.Set(r.Ctx, username, resetToken,0).Err()
	if err != nil{
		log.Fatal(err)
	}
}