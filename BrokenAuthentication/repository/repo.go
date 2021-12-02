package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)


type Repository struct {
	DB *gorm.DB
	IPCache *redis.Client
	UserCache *redis.Client
	OTPCache *redis.Client
	ResetTokenCache *redis.Client
	Ctx context.Context
}

var Repo *Repository

func (r *Repository) ClearCache(){
	r.IPCache.FlushAll(r.Ctx)
	r.UserCache.FlushAll(r.Ctx)
	r.OTPCache.FlushAll(r.Ctx)
	r.ResetTokenCache.FlushAll(r.Ctx)
}

