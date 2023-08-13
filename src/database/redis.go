package database

import (
	"context"
	"go-notif/src/config"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewRedisConn(url string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.WithField("Addr", config.RedisHost()).Fatal("Failed to connect:", err)
	}

	log.Info("Success connect redis")
	return rdb
}
