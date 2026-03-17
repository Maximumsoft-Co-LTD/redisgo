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

// SAdd adds one or more members to a set
func (c *Client) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return c.rdb.SAdd(ctx, key, members...).Err()
}

// SRem removes one or more members from a set
func (c *Client) SRem(ctx context.Context, key string, members ...interface{}) error {
	return c.rdb.SRem(ctx, key, members...).Err()
}

// SIsMember checks if a member exists in a set
func (c *Client) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return c.rdb.SIsMember(ctx, key, member).Result()
}

// SMembers returns all members of a set
func (c *Client) SMembers(ctx context.Context, key string) ([]string, error) {
	return c.rdb.SMembers(ctx, key).Result()
}

// SCard returns the number of elements in a set (Set Cardinality)
func (c *Client) SCard(ctx context.Context, key string) (int64, error) {
	return c.rdb.SCard(ctx, key).Result()
}
