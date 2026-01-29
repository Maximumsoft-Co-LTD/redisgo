package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

// SetList marshals obj to JSON and RPush to key.
func (c *Client) SetList(ctx context.Context, key string, obj interface{}) error {
	byteObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return c.rdb.RPush(ctx, key, byteObj).Err()
}

// LPush marshals obj to JSON and LPush to key.
func (c *Client) LPush(ctx context.Context, key string, obj interface{}) error {
	byteObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return c.rdb.LPush(ctx, key, byteObj).Err()
}

// PopList LPop from key and unmarshals JSON into obj.
func (c *Client) PopList(ctx context.Context, key string, obj interface{}) error {
	byteObj, err := c.rdb.LPop(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return err
		}
		return err
	}
	return json.Unmarshal([]byte(byteObj), obj)
}

// RPop pops from the right of the list at key and unmarshals JSON into obj.
func (c *Client) RPop(ctx context.Context, key string, obj interface{}) error {
	byteObj, err := c.rdb.RPop(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return err
		}
		return err
	}
	return json.Unmarshal([]byte(byteObj), obj)
}

// LRange returns elements in the list at key in the range [start, stop].
func (c *Client) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.rdb.LRange(ctx, key, start, stop).Result()
}

// LTrim trims the list at key so it contains only elements from start to stop.
func (c *Client) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.rdb.LTrim(ctx, key, start, stop).Err()
}

// LenList returns the length of the list at key.
func (c *Client) LenList(ctx context.Context, key string) (int64, error) {
	return c.rdb.LLen(ctx, key).Result()
}
