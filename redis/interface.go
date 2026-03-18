package redis

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -destination=mock_client.go -package=redis -mock_names "ClientInterface=MockClient" -source=interface.go

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// ClientInterface is the set of Redis operations provided by the library.
// *Client implements it; use MockClient in unit tests.
type ClientInterface interface {
	Close() error
	Ping(ctx context.Context) error

	Get(ctx context.Context, key string, obj interface{}) error
	Set(ctx context.Context, key string, ttl time.Duration, obj interface{}) error
	SetNoExp(ctx context.Context, key string, obj interface{}) error
	SetNX(ctx context.Context, key string, ttl time.Duration, obj interface{}) (bool, error)
	Del(ctx context.Context, key string) error
	DelMany(ctx context.Context, keyPattern string) error
	Incr(ctx context.Context, key string) (int64, error)
	Inc(ctx context.Context, key string, num int64) (int64, error)
	IsExist(ctx context.Context, key string) error
	GetKeys(ctx context.Context, keyPattern string, obj interface{}) error
	GetKeyValueMap(ctx context.Context, keyPattern string, data interface{}) error
	GetTTL(ctx context.Context, key string) (time.Duration, error)
	SetTTL(ctx context.Context, key string, ttl time.Duration) error
	Expire(ctx context.Context, key string, ttl time.Duration) error
	TTL(ctx context.Context, key string) (time.Duration, error)

	FlushDB(ctx context.Context) error
	FlushAll(ctx context.Context) error

	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)

	SetList(ctx context.Context, key string, obj interface{}) error
	LPush(ctx context.Context, key string, obj interface{}) error
	PopList(ctx context.Context, key string, obj interface{}) error
	RPop(ctx context.Context, key string, obj interface{}) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	LenList(ctx context.Context, key string) (int64, error)
	RPush(ctx context.Context, key string, values ...interface{}) error
	LPop(ctx context.Context, key string, dest interface{}) error
	LLen(ctx context.Context, key string) (int64, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) error

	SetSet(ctx context.Context, key string, ttl time.Duration, value interface{}) error
	IsMemberSet(ctx context.Context, key string, value string) (bool, error)
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SRem(ctx context.Context, key string, members ...interface{}) error
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SCard(ctx context.Context, key string) (int64, error)

	XAdd(ctx context.Context, args *redis.XAddArgs) (string, error)
	XRead(ctx context.Context, args *redis.XReadArgs) ([]redis.XStream, error)
	XReadGroup(ctx context.Context, args *redis.XReadGroupArgs) ([]redis.XStream, error)
	XAck(ctx context.Context, stream, group string, ids ...string) (int64, error)
	XGroupCreate(ctx context.Context, stream, group, start string) error
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) error
	XLen(ctx context.Context, stream string) (int64, error)
	XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error)
	XRangeN(ctx context.Context, stream, start, stop string, count int64) ([]redis.XMessage, error)

	Publish(ctx context.Context, channel string, message interface{}) error
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	PSubscribe(ctx context.Context, patterns ...string) *redis.PubSub
}
