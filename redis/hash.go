package redis

import (
	"context"
)

// HSet sets field in the hash at key to value.
func (c *Client) HSet(ctx context.Context, key, field string, value interface{}) error {
	return c.rdb.HSet(ctx, key, field, value).Err()
}

// HGet returns the value of field in the hash at key.
func (c *Client) HGet(ctx context.Context, key, field string) (string, error) {
	return c.rdb.HGet(ctx, key, field).Result()
}

// HGetAll returns all fields and values in the hash at key as a map.
func (c *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.rdb.HGetAll(ctx, key).Result()
}
