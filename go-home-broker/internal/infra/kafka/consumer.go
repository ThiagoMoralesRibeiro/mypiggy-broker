package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}

}

func (consume *Consumer) Consume(msgChannel chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(consume.ConfigMap)
	if err != nil {
		panic(err)

	}

	err = consumer.SubscribeTopics(consume.Topics, nil)

	if err != nil {
		panic(err)

	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			msgChannel <- msg

		}

	}

}
