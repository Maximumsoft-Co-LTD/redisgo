package redis_test

import (
	"context"
	"fmt"

	"github.com/Maximumsoft-Co-LTD/redisgo/redis"
)

// ExampleClient_list demonstrates SetList (RPush), PopList (LPop), and LenList (LLEN).
func ExampleClient_list() {
	client := redis.New("localhost:6379", redis.DB0)
	defer client.Close()

	key := "example:mylist"
	ctx := context.Background()

	// RPush with JSON (order: first pushed = first popped by LPop)
	err := client.SetList(ctx, key, map[string]string{"name": "first"})
	if err != nil {
		fmt.Println("rpush error:", err)
		return
	}
	client.SetList(ctx, key, map[string]string{"a": "1"})
	client.SetList(ctx, key, "item1")
	client.SetList(ctx, key, "item2")

	// LLEN
	n, err := client.LenList(ctx, key)
	if err != nil {
		fmt.Println("llen error:", err)
		return
	}
	fmt.Println("len:", n)

	// LPop and JSON unmarshalling
	var out map[string]string
	err = client.PopList(ctx, key, &out)
	if err != nil {
		fmt.Println("lpop error:", err)
		return
	}
	fmt.Println("popped:", out["name"])
	// Output:
	// len: 4
	// popped: first
}
