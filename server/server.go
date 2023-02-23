package server

import (
	"net/http"
	"time"

	"github.com/ashishbabar/go-eth-api-contract-service/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewServer(logger *zap.Logger) *http.Server {
	router := mux.NewRouter()
	redisClient := utils.NewRedisUtil()
	kafkaClient, _ := utils.NewKafkaUtil()
	logger.Info("Loaded all the utilities")

	RegisterContract(router, logger, redisClient, kafkaClient)

	logger.Info("Registered contract modules")

	logger.Info("Starting server at " + viper.GetString("HOST") + ":" + viper.GetString("PORT"))
	return &http.Server{
		Handler:      router,
		Addr:         viper.GetString("HOST") + ":" + viper.GetString("PORT"),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}
