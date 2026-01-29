package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// XAdd adds a message to a stream. args must not be nil. Returns the message ID.
func (c *Client) XAdd(args *redis.XAddArgs) (string, error) {
	ctx := context.Background()
	return c.rdb.XAdd(ctx, args).Result()
}

// XRead reads messages from streams. Use Streams as alternating stream keys and start IDs (e.g. []string{"s1", "0", "s2", "0"}).
// Block < 0 means no blocking; Block >= 0 blocks up to that duration.
func (c *Client) XRead(args *redis.XReadArgs) ([]redis.XStream, error) {
	ctx := context.Background()
	return c.rdb.XRead(ctx, args).Result()
}

// XReadGroup reads messages in a consumer group. Use Streams as alternating stream keys and IDs (e.g. []string{"s1", ">"}).
func (c *Client) XReadGroup(args *redis.XReadGroupArgs) ([]redis.XStream, error) {
	ctx := context.Background()
	return c.rdb.XReadGroup(ctx, args).Result()
}

// XAck acknowledges messages for a consumer group. Returns the number of messages acknowledged.
func (c *Client) XAck(stream, group string, ids ...string) (int64, error) {
	ctx := context.Background()
	return c.rdb.XAck(ctx, stream, group, ids...).Result()
}

// XGroupCreate creates a consumer group for stream starting at start (e.g. "0" for beginning).
func (c *Client) XGroupCreate(stream, group, start string) error {
	ctx := context.Background()
	return c.rdb.XGroupCreate(ctx, stream, group, start).Err()
}

// XGroupCreateMkStream creates a consumer group and the stream if it does not exist.
func (c *Client) XGroupCreateMkStream(stream, group, start string) error {
	ctx := context.Background()
	return c.rdb.XGroupCreateMkStream(ctx, stream, group, start).Err()
}

// XLen returns the number of entries in the stream.
func (c *Client) XLen(stream string) (int64, error) {
	ctx := context.Background()
	return c.rdb.XLen(ctx, stream).Result()
}

// XRange returns messages in the stream from start to stop (use "-" and "+" for full range).
func (c *Client) XRange(stream, start, stop string) ([]redis.XMessage, error) {
	ctx := context.Background()
	return c.rdb.XRange(ctx, stream, start, stop).Result()
}

// XRangeN returns up to count messages in the stream from start to stop.
func (c *Client) XRangeN(stream, start, stop string, count int64) ([]redis.XMessage, error) {
	ctx := context.Background()
	return c.rdb.XRangeN(ctx, stream, start, stop, count).Result()
}
