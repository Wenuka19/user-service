package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Wenuka19/user-service/internal/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewUserRepo(db *gorm.DB, cache *redis.Client) *UserRepo {
	return &UserRepo{db: db, cache: cache}
}

func (r *UserRepo) GetByID(id string) (*domain.User, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("user:%s", id)
	val, err := r.cache.Get(ctx, cacheKey).Result()

	if err == nil && val != "" {
		var cachedUser domain.User
		if err := json.Unmarshal([]byte(val), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}

	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data, _ := json.Marshal(user)
	r.cache.Set(ctx, cacheKey, data, time.Minute*10)

	return &user, nil
}

func (r *UserRepo) Save(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("user:%s", user.ID)
	data, _ := json.Marshal(user)
	r.cache.Set(ctx, cacheKey, data, time.Minute*10)

	return nil
}
