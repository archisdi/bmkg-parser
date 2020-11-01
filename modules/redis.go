package modules

import "github.com/go-redis/redis/v8"

// Redis ...
var Redis *redis.Client

// InitializeRedis ...
func InitializeRedis(host string, password string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
}
