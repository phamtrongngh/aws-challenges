package kafkabroker

import (
	"context"
	"encoding/json"
	"location/internal/model"
	"location/pkg/logger"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type kafkaLocationProducer struct {
	p     *kafka.Producer
	topic string
	log   *logrus.Logger
}

func NewKafkaLocationProducer(brokers string, topic string) *kafkaLocationProducer {
	hostname, _ := os.Hostname()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"client.id":         hostname,
		"acks":              "all",
	})

	if err != nil {
		panic(err)
	}

	log := logger.GetLogger()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("Delivery failed: %v", ev.TopicPartition)
				} else {
					log.Infof("Delivered message to %v", ev.TopicPartition)
				}
			case kafka.Error:
				log.Errorf("Kafka error: %v", ev)
			default:
				log.Infof("Kafka event: %v", ev)
			}
		}
	}()

	return &kafkaLocationProducer{p, topic, logger.GetLogger()}
}

func (k *kafkaLocationProducer) Produce(ctx context.Context, location *model.LocationUpdate) error {
	value, err := json.Marshal(location)
	if err != nil {
		return err
	}

	if err := k.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil); err != nil {
		return err
	}

	return nil
}
