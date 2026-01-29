package redis_test

import (
	"context"
	"fmt"

	"redisgo/redis"

	goredis "github.com/go-redis/redis/v8"
)

// ExampleClient_XAdd demonstrates adding a message to a stream.
func ExampleClient_XAdd() {
	client := redis.New("localhost:6379", redis.DB0)
	defer client.Close()
	id, err := client.XAdd(context.Background(), &goredis.XAddArgs{
		Stream: "example:mystream",
		Values: map[string]interface{}{"field": "value"},
	})
	if err != nil {
		fmt.Println("xadd error:", err)
		return
	}
	fmt.Println("id length:", len(id) > 0)
}

// ExampleClient_XLen demonstrates stream length.
func ExampleClient_XLen() {
	client := redis.New("localhost:6379", 0)
	defer client.Close()
	n, err := client.XLen(context.Background(), "example:mystream")
	if err != nil {
		fmt.Println("xlen error:", err)
		return
	}
	fmt.Println("len:", n >= 0)
}

// ExampleClient_XRange demonstrates reading a range of stream messages.
func ExampleClient_XRange() {
	client := redis.New("localhost:6379", redis.DB0)
	defer client.Close()
	msgs, err := client.XRange(context.Background(), "example:mystream", "-", "+")
	if err != nil {
		fmt.Println("xrange error:", err)
		return
	}
	fmt.Println("messages:", len(msgs) >= 0)
}
