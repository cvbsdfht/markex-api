package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/markex-api/pkg/core/config"
	"github.com/markex-api/pkg/core/logger"
)

func NewConnect(cfg *config.Configuration, log logger.Logger) *redis.Client {
	// Connect to server
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.Database, // use default DB
	})

	// Test connection
	if _, err := rdb.Set("ping", "pong", 30*time.Second).Result(); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Info("Pinged your deployment. You successfully connected to Redis!")

	return rdb
}
