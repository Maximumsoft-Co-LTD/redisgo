package redis

import (
	"context"
	"errors"
	"time"
)

// SetSet adds value to the set at key and sets TTL.
func (c *Client) SetSet(ctx context.Context, key string, ttl time.Duration, value interface{}) error {
	count, err := c.rdb.SAdd(ctx, key, value).Result()
	if err != nil || count == 0 {
		return err
	}
	ok, err := c.rdb.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("set expire time failed")
	}
	return nil
}

// IsMemberSet returns whether value is a member of the set at key.
func (c *Client) IsMemberSet(ctx context.Context, key string, value string) (bool, error) {
	return c.rdb.SIsMember(ctx, key, value).Result()
}
