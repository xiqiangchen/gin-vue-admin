package dsp

import (
	"github.com/IBM/sarama"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type ImpressionApi struct {
}

var impressionService = service.ServiceGroupApp.DspGroup.ImpressionService

func (impressionApi *ImpressionApi) ImpressionTrack(c *gin.Context) {
	var imp dsp.Track
	if err := c.ShouldBindQuery(&imp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imp.Check(); err != nil {
		response.IllegalWithMessage("非法请求", c)
		return
	}

	// 进入统计
	// 构建并异步发送消息
	message := &sarama.ProducerMessage{
		Topic: global.GVA_CONFIG.Kafka.Producer.Topic,
		Value: sarama.StringEncoder("Hello Kafka!"),
	}
	global.GVA_KAFKA_PRODUCER.Input() <- message

}
