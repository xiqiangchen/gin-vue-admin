package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/robfig/cron/v3"
	"time"
)

func BidTimer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务， 1-59/5 * * * *
		//_, err := global.GVA_Timer.AddTaskByFunc("LoadPlans", "@every 5m", func() {
		_, err := global.GVA_Timer.AddTaskByFunc("LoadPlans", "30 */5 * * * *", func() {
			fmt.Println("执行时间", time.Now().Minute())
			if e := bid.Load(); e != nil {
				fmt.Println("bid load campaigns error:", e)
			}
		}, "定时更新计划和活动", option...)
		if err != nil {
			fmt.Println("add bid timer error:", err)
		}

	}()
}
