package budget

import (
	"encoding/json"
	"fmt"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"sync"
	"time"
)

var _ BudgetControl = (*LocalBudgetControl)(nil)

type LocalBudgetControl struct {
	data  sync.Map          // 用于统计数据
	cache local_cache.Cache // 用于曝光去重
}

func NewLocalBudgetControl(d time.Duration) *LocalBudgetControl {
	// 创建一个新的缓存，默认过期时间为2小时，清理间隔为10分钟
	c := local_cache.NewCache(
		local_cache.SetDefaultExpire(d),
	)
	return &LocalBudgetControl{cache: c}
}

func (bc *LocalBudgetControl) Exist(key string) bool {
	_, ok := bc.GetBudgetRecord(key)
	return ok
}

func (bc *LocalBudgetControl) GetBudgetRecord(key string) (*BudgetRecord, bool) {
	if v, ok := bc.data.Load(key); ok {
		return v.(*BudgetRecord), true
	}
	return nil, false
}

func (bc *LocalBudgetControl) SetLimits(key string, dailyLimit, totalLimit float64, dailyImpressionLimit, totalImpressionLimit, dailyClickLimit, totalClickLimit int) {
	if record, ok := bc.GetBudgetRecord(key); ok {
		now := time.Now()
		record.Date = now
		record.DailyLimit = dailyLimit
		record.TotalLimit = totalLimit
		record.DailyImpressionLimit = dailyImpressionLimit
		record.TotalImpressionLimit = totalImpressionLimit
		record.DailyClickLimit = dailyClickLimit
		record.TotalClickLimit = totalClickLimit
	} else {
		bc.data.Store(key, &BudgetRecord{
			Date:                 time.Now(),
			DailyLimit:           dailyLimit,
			TotalLimit:           totalLimit,
			DailyImpressionLimit: dailyImpressionLimit,
			TotalImpressionLimit: totalImpressionLimit,
			DailyClickLimit:      dailyClickLimit,
			TotalClickLimit:      totalClickLimit,
		})
	}
}

func (bc *LocalBudgetControl) CheckOver(key string) bool {
	if v, ok := bc.GetBudgetRecord(key); ok {
		return v.CheckBudgetOver()
	}
	return false
}

func (bc *LocalBudgetControl) Get(key string) string {
	if v, ok := bc.GetBudgetRecord(key); ok {
		byt, _ := json.Marshal(v)
		return string(byt)
	}
	return ""
}

func (bc *LocalBudgetControl) Update(key string, impressionId string, amount float64, impressions, clicks int) error {
	uid := key + "_" + impressionId

	// 检查缓存中是否存在该操作的唯一标识
	if _, found := bc.cache.Get(uid); found {
		return fmt.Errorf("曝光组合(%s)已存在，不再计费", uid)
	} else {
		bc.cache.SetDefault(uid, struct{}{})
	}

	record, ok := bc.GetBudgetRecord(key)
	if !ok {
		return fmt.Errorf("no budget record found for user: %s", key)
	}

	// 检查是否是当天的记录，如果不是则重置
	now := time.Now()
	if now.Year() != record.Date.Year() || now.YearDay() != record.Date.YearDay() {
		record.Date = now
		record.DailyUsage = 0
		record.DailyImpressions = 0
		record.DailyClicks = 0
	}

	// 更新每日消耗和曝光数
	record.DailyUsage += amount
	record.DailyImpressions += impressions

	// 更新总消耗和曝光数
	record.TotalUsage += amount
	record.TotalImpressions += impressions

	// 更新点击
	record.DailyClicks += clicks
	record.TotalClicks += clicks

	// 存储更新后的记录
	bc.data.Store(key, record)

	return nil
}
