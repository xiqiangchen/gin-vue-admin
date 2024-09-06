package budget

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// global.GVA_REDIS

var _ BudgetControl = (*RedisBudgetControl)(nil)
var CostExpireTime = 14 * 24 * time.Hour

type RedisBudgetControl struct {
	pvExpireTime time.Duration
}

func NewRedisBudgetControl(expire time.Duration) *RedisBudgetControl {
	return &RedisBudgetControl{
		pvExpireTime: expire,
	}
}

func (r RedisBudgetControl) Exist(key string) bool {
	if count, err := global.GVA_REDIS.Exists(context.Background(), key).Result(); err != nil || count == 0 {
		return false
	}
	return true
}

func (r RedisBudgetControl) CheckOver(key string) bool {

	val, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	if err != nil || len(val) == 0 {
		return false
	}

	record, ok := r.GetBudgetRecord(key)
	if !ok {
		return false
	}
	return record.CheckBudgetOver()
}

func (r RedisBudgetControl) Update(key string, impressionId string, amount float64, impressions int) error {
	uid := key + "_" + impressionId

	// 检查缓存中是否存在该操作的唯一标识
	if count, err := global.GVA_REDIS.Exists(context.Background(), uid).Result(); err == nil && count > 0 {
		return fmt.Errorf("曝光组合(%s)已存在，不再计费", uid)
	} else {
		global.GVA_REDIS.Set(context.Background(), uid, "", r.pvExpireTime)
	}

	record, ok := r.GetBudgetRecord(key)
	if !ok {
		return fmt.Errorf("no budget record found for user: %s", key)
	}

	// 检查是否是当天的记录，如果不是则重置
	now := time.Now()
	if now.Year() != record.Date.Year() || now.YearDay() != record.Date.YearDay() {
		record.Date = now
		record.DailyUsage = 0
		record.DailyImpressions = 0
	}

	// 更新每日消耗和曝光数
	record.DailyUsage += amount
	record.DailyImpressions += impressions

	// 更新总消耗和曝光数
	record.TotalUsage += amount
	record.TotalImpressions += impressions

	// 存储更新后的记录
	if byt, err := json.Marshal(record); err == nil {
		global.GVA_REDIS.Set(context.Background(), key, string(byt), CostExpireTime)
	}

	return nil
}

func (r RedisBudgetControl) SetLimits(key string, dailyLimit, totalLimit float64, dailyImpressionLimit, totalImpressionLimit int) {
	record, ok := r.GetBudgetRecord(key)
	if ok {
		now := time.Now()
		record.Date = now
		record.DailyLimit = dailyLimit
		record.TotalLimit = totalLimit
		record.DailyImpressionLimit = dailyImpressionLimit
		record.TotalImpressionLimit = totalImpressionLimit
	} else {
		record = &BudgetRecord{
			Date:                 time.Now(),
			DailyLimit:           dailyLimit,
			TotalLimit:           totalLimit,
			DailyImpressionLimit: dailyImpressionLimit,
			TotalImpressionLimit: totalImpressionLimit,
		}
	}
	if byt, err := json.Marshal(record); err == nil {
		global.GVA_REDIS.Set(context.Background(), key, string(byt), CostExpireTime)
	}
}

func (r RedisBudgetControl) Get(key string) string {
	val, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	if err != nil || len(val) == 0 {
		return ""
	}
	return val
}

func (r RedisBudgetControl) GetBudgetRecord(key string) (*BudgetRecord, bool) {
	val, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	if err != nil || len(val) == 0 {
		return nil, false
	}
	record := &BudgetRecord{}

	if err = json.Unmarshal([]byte(val), record); err != nil {
		return nil, false
	}
	return record, true
}
