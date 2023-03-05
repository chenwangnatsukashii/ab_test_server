package provider

import (
	"github.com/Shopify/sarama"
	"line_china/search_server/src/utils"
	"log"
	"time"
)

var producer sarama.AsyncProducer

// InitAsyncKafka 异步的初始化
func InitAsyncKafka() error {
	log.Println("开始初始化KafKa")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机向partition发送消息，默认设置8个分区
	config.Producer.Return.Successes = true                   // 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	//使用配置,新建一个异步生产者
	var err error
	if producer, err = sarama.NewAsyncProducer(utils.Config.Kafka.Ip, config); err != nil {
		return err
	}

	//循环判断哪个通道发送过来数据.
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				log.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				log.Println("err: ", fail.Err)
			}
		}
	}(producer)

	log.Println("成功初始化KafKa")
	return nil
}

// AsyncProducer 异步生产者
func AsyncProducer(publishId int, msgInfo []byte) error {
	// defer producer.AsyncClose()
	key, err := utils.IntToBytes(publishId)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     utils.Config.Kafka.Topic,
		Key:       sarama.ByteEncoder(key),
		Value:     sarama.ByteEncoder(msgInfo),
		Timestamp: time.Now(),
	}

	// 使用通道发送
	producer.Input() <- msg
	return nil
}
