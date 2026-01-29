package redis

import (
	"context"
	"errors"
	"time"
)

// SetSet adds value to the set at key and sets TTL.
func (c *Client) SetSet(key string, ttl time.Duration, value interface{}) error {
	ctx := context.Background()
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
func (c *Client) IsMemberSet(key string, value string) (bool, error) {
	ctx := context.Background()
	return c.rdb.SIsMember(ctx, key, value).Result()
}
