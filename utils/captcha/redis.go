package captcha

import (
	"context"
	"time"

	"eva/global"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) *RedisStore {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) {
	err := global.EVA_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.EVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.EVA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.EVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.EVA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.EVA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
