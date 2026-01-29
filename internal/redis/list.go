package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

// SetList marshals obj to JSON and RPush to key.
func (c *Client) SetList(key string, obj interface{}) error {
	ctx := context.Background()
	byteObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return c.rdb.RPush(ctx, key, byteObj).Err()
}

// PopList LPop from key and unmarshals JSON into obj.
func (c *Client) PopList(key string, obj interface{}) error {
	ctx := context.Background()
	byteObj, err := c.rdb.LPop(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return err
		}
		return err
	}
	return json.Unmarshal([]byte(byteObj), obj)
}

// LenList returns the length of the list at key.
func (c *Client) LenList(key string) (int64, error) {
	ctx := context.Background()
	return c.rdb.LLen(ctx, key).Result()
}
