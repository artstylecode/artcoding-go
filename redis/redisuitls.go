package redis

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"time"
)

type RedisLockOption struct {
	redsync.Option
}

func (t RedisLockOption) Apply(Mutex *redsync.Mutex) {
}

var ctx = context.Background()

//AddOrGetRedisLock 新增/获取redis锁
func AddOrGetRedisLock(lockName string, config map[string]string) *redsync.Mutex {

	connectUrl := fmt.Sprintf("%s:%s", config["host"], config["port"])
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     connectUrl,
		Password: "", // no password set
		DB:       0,
	})

	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutex := rs.NewMutex(lockName, redsync.WithExpiry(3*time.Second))
	return mutex
}
