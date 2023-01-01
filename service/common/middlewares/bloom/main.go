package bloom

import (
	"errors"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"service/common/constant"
	"strings"
)

type bloomf struct {
	bloom   *bloom.Filter //过滤器
	store   *redis.Redis  //redis
	seconds int           //失效时间
	bits    uint          //设置默认
	key     string        //redis中保存的key
}

func NewBloom(key string, addr string, rFunc func(r *redis.Redis), bits uint, seconds int) (*bloomf, error) {
	b := &bloomf{
		bloom:   nil,
		store:   nil,
		seconds: 3600,
		bits:    1024,
		key:     constant.SysSpace,
	}

	if len(key) <= 0 {
		key = utils.NewUuid()
	}
	//创建redis 类型为BitSet
	store := newBitSetRedis(addr, rFunc)

	if store == nil {
		return b, errors.New(constant.ErrBloom00_00)
	}

	// 声明一个bitSet, key="test_key"名且bits是1024位
	if bits > b.bits {
		b.bits = bits
	}
	bitSet := bloom.New(store, key, b.bits)

	if seconds > 0 {
		b.seconds = seconds
	}

	// 默认3600秒后过期
	e := store.Expire(key, b.seconds)
	if e != nil {
		return b, errors.New(constant.ErrBloom00_01)
	}

	b.bloom = bitSet
	b.store = store
	b.key = key

	return b, nil
}

func (b bloomf) Add(data []byte) error {
	return b.bloom.Add(data)
}

// Remove 清除已经添加到redis的key
func (b bloomf) Remove() error {
	return nil
}

func (b bloomf) Exists(data []byte) (bool, error) {
	return b.bloom.Exists(data)
}

// UpdExists 更新过期时间
func (b bloomf) UpdExists(seconds int) {
	redis := b.store
	if redis == nil {
		return
	}
	e := redis.Expire(b.key, seconds)
	if e != nil {
		logx.Error("布隆过滤器中更新key：%s,的时间失败", b.key)
	}
}

func newBitSetRedis(addr string, rFunc func(r *redis.Redis)) *redis.Redis {
	if len(strings.TrimSpace(addr)) <= 0 {
		return nil
	}
	// 初始化 redisBitSet
	var store *redis.Redis = nil
	if rFunc != nil {
		store = redis.New(addr, rFunc)
	} else {
		store = redis.New(addr)
	}

	return store
}
