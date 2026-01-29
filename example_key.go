package redis_test

import (
	"context"
	"fmt"
	"time"

	"redisgo/redis"
)

// ExampleNew shows creating a direct client with New (uncached).
func ExampleNew() {
	client := redis.New("localhost:6379", redis.DB0)
	defer client.Close()
	fmt.Println("client created")
	// Output: client created
}

// ExampleConnect_reuse shows reusing a single cached connection for multiple operations.
// Requires Redis. Run with: go test -run ExampleConnect_reuse (with Redis up).
func ExampleConnect() {
	client := redis.Connect(redis.DefaultConn, "localhost:6379", redis.DB0)
	if client == nil {
		fmt.Println("no config")
		return
	}
	defer redis.Close(redis.DefaultConn)

	type Item struct{ Name string }
	ctx := context.Background()
	err := client.Set(ctx, "example:reuse:key", 10*time.Second, &Item{Name: "reused"})
	if err != nil {
		fmt.Println("set error:", err)
		return
	}

	var out Item
	_ = client.Get(ctx, "example:reuse:key", &out)

	_ = client.Del(ctx, "example:reuse:key")
	fmt.Println(out.Name)
	// Output: reused
}

// ExampleClient_Get demonstrates Get and JSON unmarshalling.
// Requires Redis and a key set first (e.g. ExampleClient_Set).
func ExampleClient_Get() {
	client := redis.New("localhost:6379", redis.DB0)
	defer client.Close()
	type Item struct{ Name string }
	var out Item
	err := client.Get(context.Background(), "example:key1", &out)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	fmt.Println(out.Name)
}
