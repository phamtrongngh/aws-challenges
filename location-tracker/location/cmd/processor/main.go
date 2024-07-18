package main

import (
	"context"
	"encoding/json"
	mongorepo "location/internal/dataaccess/repo/mongo"
	"location/internal/model"
	"location/pkg/logger"

	kafkahandler "location/internal/api/kafka/handler"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

func init() {
	// Load .env file
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func main() {
	var (
		Brokers  = viper.GetString("KAFKA_BROKERS")
		Topic    = viper.GetString("KAFKA_TOPIC")
		GroupID  = viper.GetString("KAFKA_GROUP_ID")
		MongoURI = viper.GetString("MONGO_URI")
	)

	repo := mongorepo.NewMongoLocationRepo(MongoURI)

	handler := kafkahandler.NewHandler(repo)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": Brokers,
		"group.id":          GroupID,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{Topic}, nil); err != nil {
		panic(err)
	}

	log := logger.GetLogger()

	log.Infof("Starting the consumer group [%s] on topic [%s]", GroupID, Topic)
	for {
		ev := c.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Infof("Received a message: %s", e.Value)

			locUpdate := model.LocationUpdate{}

			if err := json.Unmarshal(e.Value, &locUpdate); err != nil {
				log.Errorln(err)
				continue
			}

			if err := handler.ProcessLocationUpdate(context.Background(), &locUpdate); err != nil {
				log.Errorln(err)
			}
		case kafka.Error:
			log.Errorf("Received an error: %s", e)
		}
	}
}
