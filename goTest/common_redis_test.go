package goTest

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

var (
	rdb *redis.Client
)

func TestRedis(t *testing.T) {
	V9Example()
}

// 初始化普通连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.100:6379",
		Password: "",  // no password set
		DB:       1,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()

	result, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		return err
	}
	logx.Info(result)
	return err
}

func V9Example() {
	ctx := context.Background()
	if err := initClient(); redis.Nil != err {
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
