package budget

import (
	"encoding/json"
	"fmt"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"sync"
	"time"
)

var _ BudgetControl = (*LocalBudgetControl)(nil)

type BudgetControl interface {
	// Exist 判断是否存在
	Exist(key string) bool
	// CheckOver if the action can be performed
	CheckOver(key string) bool
	// Update the control state after the action is performed
	Update(key string, impressionId string, amount float64, impressions int) error
	// SetLimits Set the budget limits and impression limits for a user
	SetLimits(key string, dailyBudgetLimit, totalBudgetLimit float64, dailyImpressionLimit, totalImpressionLimit int)
	Print(key string) string
}
type LocalBudgetControl struct {
	data  sync.Map
	cache local_cache.Cache
}

type budgetRecord struct {
	Date                 time.Time `json:"date"`
	DailyLimit           float64   `json:"daily_limit"`
	TotalLimit           float64   `json:"total_limit"`
	DailyUsage           float64   `json:"daily_usage"`
	TotalUsage           float64   `json:"total_usage"`
	DailyImpressionLimit int       `json:"daily_impression_limit"`
	TotalImpressionLimit int       `json:"total_impression_limit"`
	DailyImpressions     int       `json:"daily_impressions"`
	TotalImpressions     int       `json:"total_impressions"`
}

func NewLocalBudgetControl(d time.Duration) *LocalBudgetControl {
	// 创建一个新的缓存，默认过期时间为2小时，清理间隔为10分钟
	c := local_cache.NewCache(
		local_cache.SetDefaultExpire(d),
	)
	return &LocalBudgetControl{cache: c}
}

func (bc *LocalBudgetControl) Exist(key string) bool {
	_, ok := bc.data.Load(key)
	return ok
}

func (bc *LocalBudgetControl) SetLimits(key string, dailyLimit, totalLimit float64, dailyImpressionLimit, totalImpressionLimit int) {
	if value, ok := bc.data.Load(key); ok {
		now := time.Now()
		record := value.(*budgetRecord)
		record.Date = now
		record.DailyLimit = dailyLimit
		record.TotalLimit = totalLimit
		record.DailyImpressionLimit = dailyImpressionLimit
		record.TotalImpressionLimit = totalImpressionLimit
	} else {
		bc.data.Store(key, &budgetRecord{
			Date:                 time.Now(),
			DailyLimit:           dailyLimit,
			TotalLimit:           totalLimit,
			DailyImpressionLimit: dailyImpressionLimit,
			TotalImpressionLimit: totalImpressionLimit,
		})
	}
}

func (bc *LocalBudgetControl) CheckOver(key string) bool {
	now := time.Now()

	value, ok := bc.data.Load(key)
	if !ok {
		return false // No record found, cannot proceed
	}

	record := value.(*budgetRecord)

	// 检查是否是当天的记录
	if now.Year() == record.Date.Year() && now.YearDay() == record.Date.YearDay() {
		if record.DailyLimit > 0 && record.DailyUsage >= record.DailyLimit || record.DailyImpressionLimit > 0 && record.DailyImpressions >= record.DailyImpressionLimit {
			return true
		}
	} else {
		// 如果不是当天的记录，重置每日消耗和曝光数
		record.Date = now
		record.DailyUsage = 0
		record.DailyImpressions = 0
	}

	// 检查总消耗和总曝光数是否超过限制
	return record.TotalLimit > 0 && record.TotalUsage >= record.TotalLimit || record.TotalImpressionLimit > 0 && record.TotalImpressions >= record.TotalImpressionLimit
}

func (bc *LocalBudgetControl) Print(key string) string {

	value, ok := bc.data.Load(key)
	if !ok {
		return ""
	}

	record := value.(*budgetRecord)

	byt, _ := json.Marshal(record)
	return string(byt)
}

func (bc *LocalBudgetControl) Update(key string, impressionId string, amount float64, impressions int) error {
	now := time.Now()
	uid := key + "_" + impressionId

	// 检查缓存中是否存在该操作的唯一标识
	if _, found := bc.cache.Get(uid); found {
		return fmt.Errorf("曝光组合(%s)已存在，不再计费", uid)
	} else {
		bc.cache.SetDefault(uid, struct{}{})
	}

	value, ok := bc.data.Load(key)
	if !ok {
		return fmt.Errorf("no budget record found for user: %s", key)
	}

	record := value.(*budgetRecord)

	// 检查是否是当天的记录，如果不是则重置
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
	bc.data.Store(key, record)

	return nil
}
