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
			cnx := RedisConnector{Config: config}
			ctx.Set(redisContextClientKey, &cnx)
			defer cnx.Client.Close()

			// Run next function
			ctx.Next()
		}
	}

}

// RedisContextClient - retrieve redis client connection from gin context
func RedisContextClient(ctx *gin.Context) *RedisConnector {
	return ctx.MustGet(redisContextClientKey).(*RedisConnector)
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

// RedisConnector connecting redis
type RedisConnector struct {
	Config *redis.Options
	Client *redis.Client
}

// InitRedis - initialize the connection
func InitRedis(conf map[string]interface{}) *RedisConnector {
	rc := RedisConnector{Config: ConfigureRedisConnection(conf)}
	return &rc
}

//
func (rc RedisConnector) context() context.Context {
	return context.Background()
}

func (rc RedisConnector) connect() *redis.Client {
	if rc.Client == nil {
		rc.Client = ConnectRedis(rc.Config)
	}
	return rc.Client
}

// Pipeline - pipelining
func (rc RedisConnector) Pipeline(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return rc.connect().Pipelined(rc.context(), fn)
}

// Transaction - transact redis
func (rc RedisConnector) Transaction(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return rc.connect().TxPipelined(rc.context(), fn)
}

// TTL - time left to expire
func (rc RedisConnector) TTL(key string) (time.Duration, error) {
	rs := rc.connect().TTL(rc.context(), key)
	return rs.Val(), rs.Err()
}

// PipedTTL - ttl in transaction
func PipedTTL(tx redis.Cmdable, key string) error {
	return tx.TTL(context.Background(), key).Err()
}

// EXISTS - redis exists
func (rc RedisConnector) EXISTS(key string) (bool, error) {
	rs := rc.connect().Exists(rc.context(), key)
	return 0 < rs.Val(), rs.Err()
}

// PipedExists - redis exists in pipeline
func PipedExists(tx redis.Cmdable, key ...string) error {
	return tx.Exists(context.Background(), key...).Err()
}

// EXPIRE - expires
func (rc RedisConnector) EXPIRE(key string, exp time.Duration) (bool, error) {
	rs := rc.connect().Expire(rc.context(), key, exp)
	return rs.Val(), rs.Err()
}

// PipedExpire - expires
func PipedExpire(tx redis.Cmdable, key string, exp time.Duration) error {
	return tx.Expire(context.Background(), key, exp).Err()
}

// GET - redis get, simplified
func (rc RedisConnector) GET(key string) (string, error) {
	rs := rc.connect().Get(rc.context(), key)
	return rs.Val(), rs.Err()
}

// PipedGet - redis get, in transaction
func PipedGet(tx redis.Cmdable, key string) error {
	return tx.Get(context.Background(), key).Err()
}

// SET - redis set, simplified
func (rc RedisConnector) SET(key string, val string, expires time.Duration) (string, error) {
	rs := rc.connect().Set(rc.context(), key, val, expires)
	return rs.Val(), rs.Err()
}

// PipedSet - redis set, in transaction
func PipedSet(tx redis.Cmdable, key string, val string) error {
	return tx.Set(context.Background(), key, val, 0).Err()
}

// HSET - redis hset
func (rc RedisConnector) HSET(key string, hash string, val string) (int64, error) {
	rs := rc.connect().HSet(rc.context(), key, hash, val)
	return rs.Val(), rs.Err()
}

// PipedHSet - redis hset
func PipedHSet(tx redis.Cmdable, key string, hash string, val string) error {
	return tx.HSet(context.Background(), key, hash, val).Err()
}

// HMSET - redis hmset
func (rc RedisConnector) HMSET(key string, vals map[string]string) (int64, error) {
	rs := rc.connect().HSet(rc.context(), key, vals)
	return rs.Val(), rs.Err()
}

// PipedHMSet - redis hset
func PipedHMSet(tx redis.Cmdable, key string, vals map[string]string) error {
	return tx.HSet(context.Background(), key, vals).Err()
}

// HGET - redis hget
func (rc RedisConnector) HGET(key string, hash string) (string, error) {
	rs := rc.connect().HGet(rc.context(), key, hash)
	return rs.Val(), rs.Err()
}

// PipedHGet - redis HGet
func PipedHGet(tx redis.Cmdable, key string, hash string) error {
	return tx.HGet(context.Background(), key, hash).Err()
}

// HMGET - redis hmset
func (rc RedisConnector) HMGET(key string, hashes ...string) ([]string, error) {
	rs := rc.connect().HMGet(rc.context(), key, hashes...)
	if rs.Err() == nil {
		rets := make([]string, len(rs.Val()))
		for i, v := range rs.Val() {
			rets[i] = v.(string)
		}
		return rets, nil
	}
	return nil, rs.Err()
}

// PipedHMGet - redis hmget
func PipedHMGet(tx redis.Cmdable, key string, hashes ...string) error {
	return tx.HMGet(context.Background(), key, hashes...).Err()
}

// HGETALL - redis hgetall
func (rc RedisConnector) HGETALL(key string) (map[string]string, error) {
	rs := rc.connect().HGetAll(rc.context(), key)
	return rs.Val(), rs.Err()
}

// PipedHGetAll - redis hgetall
func PipedHGetAll(tx redis.Cmdable, key string) error {
	return tx.HGetAll(context.Background(), key).Err()
}

// ZADD - redis zadd
func (rc RedisConnector) ZADD(key string, vals map[string]float64) (int64, error) {
	rs := rc.connect().ZAdd(rc.context(), key, zaddmaps(vals)...)
	return rs.Val(), rs.Err()
}

// PipedZAdd - redis zadd
func PipedZAdd(tx redis.Cmdable, key string, vals map[string]float64) error {
	return tx.ZAdd(context.Background(), key, zaddmaps(vals)...).Err()
}

// ZRANGE - redis zrange
func (rc RedisConnector) ZRANGE(key string, start int64, end int64) ([]string, error) {
	rs := rc.connect().ZRange(rc.context(), key, start, end)
	return rs.Val(), rs.Err()
}

// PipedZRange - redis zrange
func PipedZRange(tx redis.Cmdable, key string, start int64, end int64) error {
	return tx.ZRange(context.Background(), key, start, end).Err()
}

func zaddmaps(vals map[string]float64) []*redis.Z {
	members := make([]*redis.Z, len(vals))
	for h, s := range vals {
		members[len(members)] = &redis.Z{
			Member: h,
			Score:  s,
		}
	}
	return members
}
