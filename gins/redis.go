package gins

import (
	"context"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/pengtaikorea-adtech/go-utils/slices"
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

const redisContextClientKey = ".redis.ctx"

// ConnectRedis basic connecting function with redis
func ConnectRedis(config *redis.Options) *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(config)
	}
	return redisClient
}

// ConfigureRedisConnection - Configure Redis build configuration for redis
func ConfigureRedisConnection(conf map[string]interface{}) *redis.Options {

	configure := redis.Options{
		Addr: "127.0.0.1:6379",
	}
	if conf == nil {
	}

	if v, has := conf["network"]; has {
		configure.Network = v.(string)
	}
	if v, has := conf["addr"]; has {
		configure.Addr = v.(string)
	}
	if v, has := conf["username"]; has {
		configure.Username = v.(string)
	}
	if v, has := conf["password"]; has {
		configure.Password = v.(string)
	}

	return &configure
}

// BuildRedisConnectorMiddleware - build & return middleware
func BuildRedisConnectorMiddleware(config *redis.Options) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, exists := ctx.Get(redisContextClientKey); !exists {
			cnx := ConnectRedis(config)
			ctx.Set(redisContextClientKey, cnx)
			defer cnx.Close()

			// Run next function
			ctx.Next()
		}
	}

}

// RedisContextClient - retrieve redis client connection from gin context
func RedisContextClient(ctx *gin.Context) *redis.Client {
	return ctx.MustGet(redisContextClientKey).(*redis.Client)
}

// SaveRedisEntity - HSET
func SaveRedisEntity(tx redis.Cmdable, hashkey string, values map[string]string) error {
	rs := tx.HSet(context.Background(), hashkey, values)
	return rs.Err()
}

// LoadRedisEntity - HMGET
func LoadRedisEntity(tx redis.Cmdable, hashkey string) (map[string]string, error) {
	rs := tx.HGetAll(context.Background(), hashkey)

	return rs.Val(), rs.Err()
}

// SaveRedisEntityIndex - ZADD
func SaveRedisEntityIndex(tx redis.Cmdable, hashkey string, primaryKey ...string) error {
	timestamp := float64(time.Now().Unix())

	indices, _ := slices.Map(func(e interface{}, i int, s interface{}) (interface{}, error) {
		pk := e.(string)
		return &redis.Z{
			Score:  timestamp,
			Member: pk,
		}, nil
	}, primaryKey, reflect.TypeOf(&redis.Z{}))
	puts := indices.([]*redis.Z)
	tx.ZAdd(context.Background(), hashkey, puts...)
	return nil
}
