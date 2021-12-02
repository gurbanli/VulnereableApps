package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func(r *Repository) InitializeIPCacheRepository(){
	r.IPCache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // IP cache db
	})
	r.Ctx = context.Background()
}

func(r *Repository) GetIPBlockStatus(ip string, increaseAttempt bool) bool{
	value, err := r.IPCache.Get(r.Ctx, ip).Result()
	if err != nil{
		r.SetIPBlockStatus(ip, 1)
		return false
	}
	attempt, err := strconv.Atoi(value)
	if err != nil{
		log.Fatal(err)
	}
	if attempt >= 3 {
		return true
	}else{
		if increaseAttempt{
			attempt++
		}
		r.SetIPBlockStatus(ip, attempt)
		return false
	}

}

func (r *Repository) SetIPBlockStatus(ip string, attempt int){
	err := r.IPCache.Set(r.Ctx, ip, attempt,0).Err()
	if err != nil{
		log.Fatal(err)
	}
}
