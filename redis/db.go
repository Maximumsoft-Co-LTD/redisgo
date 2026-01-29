package redis

import "context"

// DefaultDB is the default Redis logical database index (0).

// FlushDB removes all keys in the current database.
func (c *Client) FlushDB(ctx context.Context) error {
	return c.rdb.FlushDB(ctx).Err()
}

// FlushAll removes all keys from all databases.
func (c *Client) FlushAll(ctx context.Context) error {
	return c.rdb.FlushAll(ctx).Err()
}
