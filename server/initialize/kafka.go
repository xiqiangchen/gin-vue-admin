package initialize

import (
	"github.com/IBM/sarama"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var KafkaProducer = new(kafkaProducer)

type (
	kafkaProducer struct{}
)

func (kafkaProducer *kafkaProducer) Initialization() error {

	// Kafka 配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	if global.GVA_CONFIG.Kafka.Producer != nil && len(global.GVA_CONFIG.Kafka.Brokers) > 0 {

		// 创建 Kafka 生产者实例
		if producer, err := sarama.NewAsyncProducer(global.GVA_CONFIG.Kafka.Brokers, config); err != nil {
			return errors.Wrap(err, "创建kafka生产者失败!")

		} else {
			global.GVA_KAFKA_PRODUCER = producer
		}

	}
	// 处理成功和错误的消息
	go func() {
		for {
			select {
			case err := <-global.GVA_KAFKA_PRODUCER.Errors():
				global.GVA_LOG.Error("Failed to send message:", zap.Error(err))
			}
		}
	}()
	return nil
}

func (kafkaProducer *kafkaProducer) Close() {
	if err := global.GVA_KAFKA_PRODUCER.Close(); err != nil {
		global.GVA_LOG.Error("Error closing Kafka producer:", zap.Error(err))
	}
}
