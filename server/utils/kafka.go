package utils

import (
	"github.com/IBM/sarama"
)

func SendMsg(producer sarama.AsyncProducer, topic string, msg []byte) {

	// 进入统计
	// 构建并异步发送消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}
	producer.Input() <- message

}

func SendStrMsg(producer sarama.AsyncProducer, topic string, msg string) {
	// 进入统计
	// 构建并异步发送消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	producer.Input() <- message
}
