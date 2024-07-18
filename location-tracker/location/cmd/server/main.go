package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "location/docs" // import the generated docs package

	httpapi "location/internal/api/http"
	"location/pkg/logger"

	kafkabroker "location/internal/dataaccess/broker/kafka"
	mongorepo "location/internal/dataaccess/repo/mongo"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Location API
// @version 1.0
func main() {
	var (
		ServerPort = viper.GetString("SERVER_PORT")
		Brokers    = viper.GetString("KAFKA_BROKERS")
		Topic      = viper.GetString("KAFKA_TOPIC")
		ServerAddr = ":" + ServerPort
		MongoURI   = viper.GetString("MONGO_URI")
	)

	log := logger.GetLogger()

	p := kafkabroker.NewKafkaLocationProducer(Brokers, Topic)
	r := mongorepo.NewMongoLocationRepo(MongoURI)

	handler := httpapi.NewHandler(p, r)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/location", handler.UpdateLocation)
	mux.HandleFunc("/v1/location/upload", handler.UploadLocation)
	mux.HandleFunc("/v1/location/{device_id}", handler.GetLatestLocation)

	// Serve swagger docs
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("docs/swagger.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var swaggerDoc map[string]interface{}
		if err := json.Unmarshal(file, &swaggerDoc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		swaggerDoc["host"] = fmt.Sprintf("0.0.0.0:%s", ServerPort)

		modifiedSwaggerDoc, err := json.Marshal(swaggerDoc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(modifiedSwaggerDoc)
	})

	log.Infoln("Server is running on", ServerAddr)
	if err := http.ListenAndServe(ServerAddr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
