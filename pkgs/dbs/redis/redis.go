package redis

import (
	"context"
	"encoding/json"
	goredis "github.com/go-redis/redis/v8"
	"time"
	"trekkstay/pkgs/log"
)

// Timeout is the default timeout for redis operations
const (
	Timeout = 1
)

// redis is a struct that implements the Redis interface.
type redis struct {
	cmd goredis.Cmdable
}

// NewRedis creates a new Redis instance and returns a Redis interface.
func NewRedis(connection Connection) Redis {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	rdb := goredis.NewClient(&goredis.Options{
		Addr:     connection.Address,
		Password: connection.Password,
		DB:       connection.Database,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.JsonLogger.Error(pong, err)
		return nil
	}

	return &redis{
		cmd: rdb,
	}
}

// IsConnected checks if the redis connection is alive.
func (r *redis) IsConnected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	if r.cmd == nil {
		return false
	}

	_, err := r.cmd.Ping(ctx).Result()
	if err != nil {
		return false
	}

	return true
}

// Get gets the value of a key and stores it in the value pointer.
func (r *redis) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	strValue, err := r.cmd.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(strValue), value)
	if err != nil {
		return err
	}

	return nil
}

// SetWithExpiration sets the value of a key with an expiration time.
func (r *redis) SetWithExpiration(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	bData, _ := json.Marshal(value)
	err := r.cmd.Set(ctx, key, bData, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Set sets the value of a key.
func (r *redis) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	bData, _ := json.Marshal(value)
	err := r.cmd.Set(ctx, key, bData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Remove removes a key from redis.
func (r *redis) Remove(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	err := r.cmd.Del(ctx, keys...).Err()
	if err != nil {
		return err
	}

	return nil
}

// Keys returns all keys matching a pattern.
func (r *redis) Keys(pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	keys, err := r.cmd.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// RemovePattern removes all keys matching a pattern.
func (r *redis) RemovePattern(pattern string) error {
	keys, err := r.Keys(pattern)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	err = r.Remove(keys...)
	if err != nil {
		return err
	}

	return nil
}
