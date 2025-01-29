package lib

import (
	"strings"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type MyKafka struct {
	Producer sarama.SyncProducer
}

var ProducerPool []*MyKafka

func InitProducer() {
	brokerList := []string{"localhost:29092", "localhost:39092", "localhost:49092"}
	for _, p := range ProducerPool {
		(*p), _ = NewKafkaProducer(brokerList)
	}
}

func RegisterProducer(p *MyKafka) {
	ProducerPool = append(ProducerPool, p)
}

func NewKafkaProducer(brokers []string) (myKafka MyKafka, err error) {
	brokerList := strings.Join(brokers, ",")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		singleton.Logger.Error("Failed to create kafka producer", zap.String("error", err.Error()))
		return
	}
	myKafka = MyKafka{Producer: producer}
	return
}

func (m *MyKafka) Push(topic, key, message string) (err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}
	_, _, err = m.Producer.SendMessage(msg)
	if err != nil {
		singleton.Logger.Error("Failed to send message", zap.String("error", err.Error()))
		return
	}
	singleton.Logger.Info("Message sent", zap.String("topic", topic), zap.String("key", key), zap.String("message", message))
	return
}

func NewKafkaConsumer(brokers []string, group string, topics []string) (consumer sarama.ConsumerGroup, err error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err = sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		singleton.Logger.Error("Failed to create kafka consumer", zap.String("error", err.Error()))
		return
	}

	go func() {
		for err := range consumer.Errors() {
			singleton.Logger.Error("Consumer error", zap.String("error", err.Error()))
		}
	}()

	return
}
