package redis

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
)

// Get reads key and unmarshals JSON into obj.
func (c *Client) Get(key string, obj interface{}) error {
	ctx := context.Background()
	byteObj, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(byteObj), obj)
}

// Set marshals obj to JSON and sets key with TTL.
func (c *Client) Set(key string, ttl time.Duration, obj interface{}) error {
	ctx := context.Background()
	byteObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return c.rdb.SetEX(ctx, key, byteObj, ttl).Err()
}

// Del deletes key.
func (c *Client) Del(key string) error {
	ctx := context.Background()
	return c.rdb.Del(ctx, key).Err()
}

// DelMany finds keys matching keyPattern and deletes them.
func (c *Client) DelMany(keyPattern string) error {
	ctx := context.Background()
	keys, err := c.rdb.Keys(ctx, keyPattern).Result()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	return c.rdb.Del(ctx, keys...).Err()
}

// IsExist returns nil if key exists, error if it does not.
func (c *Client) IsExist(key string) error {
	ctx := context.Background()
	exists, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return errors.New("key does not exist")
	}
	return nil
}

// Inc increments key by num and returns the new value.
func (c *Client) Inc(key string, num int64) (int64, error) {
	ctx := context.Background()
	return c.rdb.IncrBy(ctx, key, num).Result()
}

// GetKeys writes keys matching keyPattern into obj. obj must be *[]string.
func (c *Client) GetKeys(keyPattern string, obj interface{}) error {
	ctx := context.Background()
	keys, err := c.rdb.Keys(ctx, keyPattern).Result()
	if err != nil {
		return err
	}
	if obj == nil {
		return nil
	}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	slice := v.Elem()
	if slice.Kind() != reflect.Slice {
		return nil
	}
	// set *[]string
	slice.Set(reflect.ValueOf(keys))
	return nil
}

// GetKeyValueMap loads all keys matching keyPattern and their JSON values into data.
// data must be a pointer to map[string]T where T matches the stored value types.
func (c *Client) GetKeyValueMap(keyPattern string, data interface{}) error {
	ctx := context.Background()
	keys, err := c.rdb.Keys(ctx, keyPattern).Result()
	if err != nil {
		return err
	}
	result := make(map[string]interface{})
	for _, key := range keys {
		val, err := c.rdb.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return err
		}
		var obj interface{}
		if err := json.Unmarshal([]byte(val), &obj); err != nil {
			return err
		}
		result[key] = obj
	}
	byteData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, data)
}

// GetTTL returns the TTL of key.
func (c *Client) GetTTL(key string) (time.Duration, error) {
	ctx := context.Background()
	return c.rdb.TTL(ctx, key).Result()
}

// SetTTL sets the TTL for key.
func (c *Client) SetTTL(key string, ttl time.Duration) error {
	ctx := context.Background()
	return c.rdb.Expire(ctx, key, ttl).Err()
}
