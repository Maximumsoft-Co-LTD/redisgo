# smskub

Redis client helpers built on [go-redis/v8](https://github.com/go-redis/redis). Provides a thin wrapper with JSON key/value, list, set, and stream helpers.

## Module

- **Module:** `smskub`
- **Internal package:** `smskub/internal/redis` (only importable from within this repository)

## Usage

```go
import "smskub/internal/redis"

// One-off or test connection
client := redis.New("localhost:6379", 0)
defer client.Close()

// Cached connection (config from env: REDIS_URL, REDIS_DB)
client = redis.Connect("default")
if client != nil {
    defer redis.Close("default")
}
```

### Keys (JSON)

- `Get(key, &obj)`, `Set(key, ttl, obj)` — JSON get/set with TTL
- `Del(key)`, `DelMany(pattern)` — delete keys
- `IsExist(key)`, `Inc(key, n)` — existence and increment
- `GetKeys(pattern, &[]string)`, `GetKeyValueMap(pattern, &map)` — key scan and bulk load
- `GetTTL(key)`, `SetTTL(key, ttl)` — TTL get/set

### Lists

- `SetList(key, obj)` — RPush JSON
- `PopList(key, &obj)` — LPop and unmarshal
- `LenList(key)` — list length

### Sets

- `SetSet(key, ttl, value)` — SAdd and set TTL
- `IsMemberSet(key, value)` — SIsMember

### Streams

- `XAdd(args)`, `XRead(args)`, `XReadGroup(args)` — produce/consume
- `XAck`, `XGroupCreate`, `XGroupCreateMkStream`
- `XLen(stream)`, `XRange(stream, start, stop)`, `XRangeN(stream, start, stop, count)`

## Build and test

```bash
go mod tidy
go build ./...
go test ./...
```

Example tests in `example_test.go` assume a Redis instance on `localhost:6379` for examples that hit the server.
