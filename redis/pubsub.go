package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Publish publishes message to channel.
func (c *Client) Publish(ctx context.Context, channel string, message interface{}) error {
	return c.rdb.Publish(ctx, channel, message).Err()
}

// Subscribe subscribes to the given channels and returns a PubSub. Caller must call pubsub.Close() when done.
func (c *Client) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return c.rdb.Subscribe(ctx, channels...)
}

// PSubscribe subscribes to channels matching the given patterns and returns a PubSub. Caller must call pubsub.Close() when done.
func (c *Client) PSubscribe(ctx context.Context, patterns ...string) *redis.PubSub {
	return c.rdb.PSubscribe(ctx, patterns...)
}
