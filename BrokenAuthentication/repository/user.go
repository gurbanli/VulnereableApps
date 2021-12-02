package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func(r *Repository) InitializeUserCacheRepository(){
	r.UserCache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // User cache db
	})
	r.Ctx = context.Background()
}

func(r *Repository) GetUserBlockStatus(username string, increaseAttempt bool) bool{
	value, err := r.UserCache.Get(r.Ctx, username).Result()
	if err != nil{
		r.SetIPBlockStatus(username, 1)
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
		r.SetIPBlockStatus(username, attempt)
		return false
	}

}


func (r *Repository) SetUserBlockStatus(username string, attempt int){
	err := r.UserCache.Set(r.Ctx, username, attempt,0).Err()
	if err != nil{
		log.Fatal(err)
	}
}

