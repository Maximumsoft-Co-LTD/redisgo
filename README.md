# redisgo

Redis client helpers for Go built on [go-redis/v8](https://github.com/go-redis/redis). JSON keys, lists, sets, hashes, streams, and pub/sub.

- **Go:** 1.21+
- **Redis:** any compatible server

## Install

```bash
go get github.com/Maximumsoft-Co-LTD/redisgo
```

Replace `github.com/Maximumsoft-Co-LTD/redisgo` with your module path from `go.mod`.

### Local development (use from another project on the same machine)

In the project that needs this library, add to `go.mod`:

```go
module myapp

require redisgo v0.0.0

replace redisgo => /path/to/redisgo   // path to the redisgo directory
```

Then in code use `import "redisgo/redis"` (or the module path from your replace).

## Quick start

```go
package main

import (
	"context"
	"time"

	"redisgo/redis"
)

func main() {
	// Direct connection (e.g. tests or one-off use)
	client := redis.New("localhost:6379")
	defer client.Close()

	// Or cached by name (same addr returns same client)
	client = redis.Connect(redis.DefaultConn, "localhost:6379")
	defer redis.Close(redis.DefaultConn)

	// JSON get/set (pass context for cancellation/timeout)
	ctx := context.Background()
	var out string
	client.Set(ctx, "key", 10*time.Minute, "hello")
	client.Get(ctx, "key", &out)
}
```

## API

### Keys (JSON)

| Method | Description |
|--------|-------------|
| `Get(ctx, key, &obj)` | Get and unmarshal into obj |
| `Set(ctx, key, ttl, obj)` | Set key with TTL (JSON marshalled) |
| `SetNoExp(ctx, key, obj)` | Set key without expiry (JSON marshalled) |
| `SetNX(ctx, key, ttl, obj)` | Set key only if not exists (JSON), returns (bool, error) |
| `Del(ctx, key)` | Delete key |
| `DelMany(ctx, pattern)` | Delete all keys matching pattern |
| `IsExist(ctx, key)` | Returns nil if key exists, error otherwise |
| `Incr(ctx, key)` | Increment by 1, returns new value |
| `Inc(ctx, key, n)` | Increment by n, returns new value |
| `GetKeys(ctx, pattern, &[]string)` | Scan keys matching pattern |
| `GetKeyValueMap(ctx, pattern, &map)` | Load keys and JSON values into map |
| `GetTTL(ctx, key)`, `SetTTL(ctx, key, ttl)` | Get/set TTL |
| `Ping(ctx)` | Check connection to Redis |

### Hashes

| Method | Description |
|--------|-------------|
| `HSet(ctx, key, field, value)` | Set hash field |
| `HGet(ctx, key, field)` | Get hash field value |
| `HGetAll(ctx, key)` | Get all hash fields and values |

### Lists

| Method | Description |
|--------|-------------|
| `SetList(ctx, key, obj)` | RPush JSON to list |
| `LPush(ctx, key, obj)` | LPush JSON to list |
| `PopList(ctx, key, &obj)` | LPop and unmarshal into obj |
| `RPop(ctx, key, &obj)` | RPop and unmarshal into obj |
| `LRange(ctx, key, start, stop)` | Get list range (returns []string) |
| `LTrim(ctx, key, start, stop)` | Trim list to range |
| `LenList(ctx, key)` | List length |

### Sets

| Method | Description |
|--------|-------------|
| `SetSet(ctx, key, ttl, value)` | SAdd value and set TTL on key |
| `IsMemberSet(ctx, key, value)` | Returns whether value is in set (value is string) |

### Streams

| Method | Description |
|--------|-------------|
| `XAdd(ctx, args)` | Add message; returns message ID |
| `XRead(ctx, args)` | Read from streams (Block < 0 = no block) |
| `XReadGroup(ctx, args)` | Read in consumer group |
| `XAck(ctx, stream, group, ids...)` | Acknowledge messages |
| `XGroupCreate`, `XGroupCreateMkStream` | Create consumer group |
| `XLen(ctx, stream)` | Number of entries |
| `XRange`, `XRangeN` | Range of messages (use "-", "+" for full) |

### Pub/Sub

| Method | Description |
|--------|-------------|
| `Publish(ctx, channel, message)` | Publish message to channel |
| `Subscribe(ctx, channels...)` | Subscribe to channels (returns *redis.PubSub; call Close when done) |
| `PSubscribe(ctx, patterns...)` | Subscribe to channels matching patterns (returns *redis.PubSub) |

## Connection helpers

- `DefaultConn` — constant for the default connection name ("default").
- `New(addr)` — create a client (not cached).
- `Connect(name, addr)` — get or create a cached client for that name.
- `Close(name)` — close and remove cached client for name.
- `CloseAll()` — close all cached clients and clear the cache.

## Development

```bash
go mod tidy
go build ./...
go test ./...
```

Example tests use Redis at `localhost:6379`.
