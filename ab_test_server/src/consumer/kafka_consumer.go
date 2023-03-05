package consumer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"line_china/ab_test_server/src/service"
	"line_china/ab_test_server/src/utils"
	"line_china/common/model"
	"log"
	"sync"
)

var KafkaClient sarama.Consumer

// Setup is run at the beginning of a new session, before ConsumeClaim
func (k *Kafka) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("setup")
	session.ResetOffset(k.topics[0], 0, 13, "")
	log.Println(session.Claims())
	// Mark the consumer as ready
	close(k.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (k *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (k *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// <https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29>
	// 具体消费消息
	for message := range claim.Messages() {
		log.Printf("[topic:%s] [partiton:%d] [offset:%d] [value:%s] [time:%v]\n",
			message.Topic, message.Partition, message.Offset, string(message.Value), message.Timestamp)

		ConfigData := model.ResultModel{}
		err := json.Unmarshal(message.Value, &ConfigData)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		// 从连接池获取Redis连接
		conn := utils.RedisPool.Get()

		n, err := conn.Do("SETNX", ConfigData.UserId, 1)
		if err != nil {
			log.Println(err)
		}
		if n == int64(1) {
			service.AddOneRecord(ConfigData)
		}
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

type Kafka struct {
	brokers           []string
	topics            []string
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
	assignor          string
}

func NewKafka() *Kafka {
	return &Kafka{
		brokers:           utils.Config.Kafka.Ip,
		topics:            utils.Config.Kafka.Topic,
		group:             utils.Config.Kafka.Group,
		channelBufferSize: 1000,
		ready:             make(chan bool),
		version:           "2.8.0",
		assignor:          utils.Config.Kafka.Assignor,
	}
}

func (k *Kafka) KafkaInit() func() {
	log.Print("开始初始化Kafka")

	version, err := sarama.ParseKafkaVersion(k.version)
	if err != nil {
		log.Printf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	// 分区分配策略
	switch k.assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundRobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", k.assignor)
	}

	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ChannelBufferSize = k.channelBufferSize // channel长度

	// 创建client
	newClient, err := sarama.NewClient(k.brokers, config)
	if err != nil {
		log.Println(err)
	}

	// 获取所有的topic
	topics, err := newClient.Topics()
	if err != nil {
		log.Println(err)
	}
	log.Println("topics: ", topics)

	// 根据client创建consumerGroup
	client, err := sarama.NewConsumerGroupFromClient(k.group, newClient)
	if err != nil {
		log.Printf("Error creating consumer group client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, k.topics, k); err != nil {
				// 当setup失败的时候，error会返回到这里
				log.Printf("Error from consumer: %v", err)
				return
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx)
				return
			}
			k.ready = make(chan bool)
		}
	}()
	<-k.ready

	log.Print("成功初始化Kafka")

	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Println("kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Printf("Error closing client: %v", err)
		}
	}
}
