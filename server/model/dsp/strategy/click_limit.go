package strategy

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

const ClickExpireTime = 7 * 24 * time.Hour

type ClickLimit struct {
	ClickExpireTime time.Duration
}

func NewClickLimit(t time.Duration) *ClickLimit {
	return &ClickLimit{ClickExpireTime: t}
}

// SetLimitWithTime 设置访问次数
func (c *ClickLimit) Incr(key string) (int, error) {
	count, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	if count == 0 {
		pipe := global.GVA_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, c.ClickExpireTime)
		_, err = pipe.Exec(context.Background())
		return 1, err
	} else {
		// 次数
		if times, err := global.GVA_REDIS.Get(context.Background(), key).Int(); err != nil {
			return 0, err
		} else {
			err := global.GVA_REDIS.Incr(context.Background(), key).Err()
			return times + 1, err
		}
	}
}
