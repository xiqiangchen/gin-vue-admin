package budget

import (
	"time"
)

type BudgetControl interface {
	// Exist 判断是否存在
	Exist(key string) bool
	// CheckOver if the action can be performed
	CheckOver(key string) bool
	// Update the control state after the action is performed
	Update(key string, impressionId string, amount float64, impressions, clicks int) error
	// SetLimits Set the budget limits and impression limits for a user
	SetLimits(key string, dailyBudgetLimit, totalBudgetLimit float64, dailyImpressionLimit, totalImpressionLimit, dailyClickLimit, totalClickLimit int)
	Get(key string) string
	GetBudgetRecord(key string) (*BudgetRecord, bool)
}

type BudgetRecord struct {
	Date                 time.Time `json:"date"`
	DailyLimit           float64   `json:"daily_limit"`
	TotalLimit           float64   `json:"total_limit"`
	DailyUsage           float64   `json:"daily_usage"`
	TotalUsage           float64   `json:"total_usage"`
	DailyImpressionLimit int       `json:"daily_impression_limit"`
	TotalImpressionLimit int       `json:"total_impression_limit"`
	DailyImpressions     int       `json:"daily_impressions"`
	TotalImpressions     int       `json:"total_impressions"`
	DailyClickLimit      int       `json:"daily_Click_limit"`
	TotalClickLimit      int       `json:"total_Click_limit"`
	DailyClicks          int       `json:"daily_Clicks"`
	TotalClicks          int       `json:"total_Clicks"`
}

func (record *BudgetRecord) CheckBudgetOver() bool {
	// 检查是否是当天的记录
	now := time.Now()
	if now.Year() == record.Date.Year() && now.YearDay() == record.Date.YearDay() {
		if record.DailyLimit > 0 && record.DailyUsage >= record.DailyLimit ||
			record.DailyImpressionLimit > 0 && record.DailyImpressions >= record.DailyImpressionLimit ||
			record.DailyClickLimit > 0 && record.DailyClicks >= record.DailyClickLimit {
			return true
		}
	} else {
		// 如果不是当天的记录，重置每日消耗和曝光数
		record.Date = now
		record.DailyUsage = 0
		record.DailyImpressions = 0
		record.DailyClicks = 0
	}

	// 检查总消耗和总曝光数是否超过限制
	return record.TotalLimit > 0 && record.TotalUsage >= record.TotalLimit ||
		record.TotalImpressionLimit > 0 && record.TotalImpressions >= record.TotalImpressionLimit ||
		record.TotalClickLimit > 0 && record.TotalClicks >= record.TotalClickLimit
}
