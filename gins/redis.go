package gins

import (
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// ConfigureRedis - load redis configure
func ConfigureRedis(path string) *redis.Options {
	// default configure
	option := redis.Options{
		Addr: "6379",
	}

	// TODO: load from the path

	return &option
}

// SaveRedisEntity - HSET
func SaveRedisEntity(tx redis.Cmdable, hashkey string, values map[string]string) error {
	rs := tx.HSET(hashkey, values)
	return rs.Err()
}

// LoadRedisEntity - HMGET
func LoadRedisEntity(tx redis.Cmdable, hashkey string) (map[string]string, error) {
	rs := tx.HMGET()
	return rs.Val(), rs.Err()
}

// SaveRedisEntityIndex - ZADD
func SaveRedisEntityIndex(tx redis.Cmdable, hashkey string, primaryKey ...string) error {
	timestamp := float64(time.Now().Unix())
	indices := make([]*redis.Z, len(primaryKey))
}
