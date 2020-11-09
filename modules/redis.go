package modules

import (
	"encoding/json"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

// RedisClient ...
type RedisClient struct {
	*redis.Client
}

// Redis ...
var Redis RedisClient

// InitializeRedis ...
func InitializeRedis(host string, username string, password string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Username: username,
		Password: password,
		DB:       0,
	})

	Redis = RedisClient{
		client,
	}

	return nil
}

// SetCache ...
func (c *RedisClient) SetCache(key string, value interface{}) error {
	p, err := json.Marshal(value)

	if err != nil {
		return err
	}

	c.Set(context.TODO(), key, p, time.Minute * 10)

	return nil
}

// GetCache ...
func (c *RedisClient) GetCache(key string, dest interface{}) error {
	p := c.Get(context.TODO(), key)
	return json.Unmarshal([]byte(p.Val()), dest)
}
