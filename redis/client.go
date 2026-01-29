package redis

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const DB0 = 0

// DefaultConn is the default connection name used by Connect/Close.
const DefaultConn = "default"

// Client wraps a redis.Client and provides key, hash, list, set, stream, and pub/sub operations.
type Client struct {
	rdb *redis.Client
}

// ConnConfig holds env var keys for a named connection.
type ConnConfig struct {
	URLKey string // e.g. "REDIS_URL"
	DBKey  string // e.g. "REDIS_DB"
}

var (
	// DefaultConfig is the named config for DefaultConn (REDIS_URL, REDIS_DB).
	DefaultConfig = map[string]ConnConfig{
		DefaultConn: {URLKey: "REDIS_URL", DBKey: "REDIS_DB"},
	}
	connCache   sync.Map
	connConfigs = DefaultConfig
)

// New creates a new Redis client (uncached). Use for tests or one-off connections.
func New(addr string, db int) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "",
		DB:          db,
		IdleTimeout: 1 * time.Minute,
	})
	return &Client{rdb: rdb}
}

// Connect returns a cached Client for the given name. Config is read from
// connConfigs (env vars). First call for a name creates and caches the client.
func Connect(name string, addr string, db int) *Client {
	if c, ok := connCache.Load(name); ok {
		return c.(*Client)
	}

	client := New(addr, db)
	connCache.Store(name, client)
	return client
}

// Close closes this client's connection. For cached clients from Connect(name),
// prefer Close(name) or CloseAll() so the cache is updated.
func (c *Client) Close() error {
	return c.rdb.Close()
}

// Close removes the cached client for name and closes it. Next Connect(name)
// will create a new client.
func Close(name string) error {
	v, ok := connCache.LoadAndDelete(name)
	if !ok {
		return nil
	}
	return v.(*Client).Close()
}

// CloseAll closes all cached clients and clears the cache.
func CloseAll() {
	connCache.Range(func(key, value interface{}) bool {
		value.(*Client).Close()
		connCache.Delete(key)
		return true
	})
}
