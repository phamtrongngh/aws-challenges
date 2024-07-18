package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Load .env file
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

var log *logrus.Logger

func init() {
	Level := viper.GetString("LOG_LEVEL")

	log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
}

func GetLogger() *logrus.Logger {
	return log
}
