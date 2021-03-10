package main

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/redis"
)

func main() {
	mutex := redis.AddOrGetRedisLock("test_1", map[string]string{
		"host": "127.0.0.1",
		"port": "16379",
	})
	res := mutex.Lock()
	fmt.Println(res)
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
