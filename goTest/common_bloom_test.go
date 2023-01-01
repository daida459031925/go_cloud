package goTest

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"os"
	"testing"
)

// 布隆过滤器测试
func TestBloom(t *testing.T) {
	// 初始化 redisBitSet
	store := redis.New("192.168.0.10:6379", func(r *redis.Redis) {
		r.Type = redis.NodeType
	})

	s, e := store.Get("")
	if e != nil {
		logx.Info("关闭程序")
		os.Exit(3)
	}
	logx.Info(s)
	// 声明一个bitSet, key="test_key"名且bits是1024位
	bitSet := bloom.New(store, "", 1024)
	bitSet.Add([]byte("asd"))
	logx.Info(bitSet)
}
