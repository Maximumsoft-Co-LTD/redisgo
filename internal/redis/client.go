package redis

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// Client wraps a redis.Client and provides key, list, set, and stream operations.
type Client struct {
	rdb *redis.Client
}

// ConnConfig holds env var keys for a named connection.
type ConnConfig struct {
	URLKey string // e.g. "REDIS_URL"
	DBKey  string // e.g. "REDIS_DB"
}

var (
	// DefaultConfig is the named config for "default" (REDIS_URL, REDIS_DB).
	DefaultConfig = map[string]ConnConfig{
		"default": {URLKey: "REDIS_URL", DBKey: "REDIS_DB"},
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
func Connect(name string) *Client {
	if c, ok := connCache.Load(name); ok {
		return c.(*Client)
	}
	cfg, ok := connConfigs[name]
	if !ok {
		return nil
	}
	addr := os.Getenv(cfg.URLKey)
	db, _ := strconv.Atoi(os.Getenv(cfg.DBKey))
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
