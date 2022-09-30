package bloom

import (
	"errors"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"strings"
)

func NewBloom(key string, addr string, rFunc func(r *redis.Redis), i uint) (*bloom.Filter, error) {
	if len(key) <= 0 {
		key = utils.NewUuid()
	}
	if len(strings.TrimSpace(addr)) <= 0 {
		return nil, errors.New("")
	}
	// 初始化 redisBitSet
	var store *redis.Redis = nil
	if rFunc != nil {
		store = redis.New(addr, rFunc)
	} else {
		store = redis.New(addr)
	}
	// 声明一个bitSet, key="test_key"名且bits是1024位
	if i < 1024 {
		i = 1024
	}
	bitSet := bloom.New(store, key, i)

	store.Expire()

	bitSet.
}
