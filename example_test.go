package redis_test

import (
	"fmt"
	"os"
	"time"

	"smskub/internal/redis"
)

// ExampleNew shows creating a direct client with New (uncached).
func ExampleNew() {
	client := redis.New("localhost:6379", 0)
	defer client.Close()
	fmt.Println("client created")
	// Output: client created
}

// ExampleConnect_reuse shows reusing a single cached connection for multiple operations.
// Requires Redis. Run with: go test -run ExampleConnect_reuse (with Redis up).
func ExampleConnect() {
	_ = os.Setenv("REDIS_URL", "localhost:6379")
	_ = os.Setenv("REDIS_DB", "0")
	client := redis.Connect("default")
	if client == nil {
		fmt.Println("no config")
		return
	}
	defer redis.Close("default")

	type Item struct{ Name string }
	err := client.Set("example:reuse:key", 10*time.Second, &Item{Name: "reused"})
	if err != nil {
		fmt.Println("set error:", err)
		return
	}

	var out Item
	_ = client.Get("example:reuse:key", &out)

	_ = client.Del("example:reuse:key")
	fmt.Println(out.Name)
	// Output: reused
}


// ExampleClient_Get demonstrates Get and JSON unmarshalling.
// Requires Redis and a key set first (e.g. ExampleClient_Set).
func ExampleClient_Get() {
	client := redis.New("localhost:6379", 0)
	defer client.Close()
	type Item struct{ Name string }
	var out Item
	err := client.Get("example:key1", &out)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	fmt.Println(out.Name)
}

