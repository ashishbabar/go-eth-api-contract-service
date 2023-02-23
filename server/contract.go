package server

import (
	"github.com/ashishbabar/go-eth-api-contract-service/contracts"
	"github.com/ashishbabar/go-eth-api-contract-service/utils"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func RegisterContract(router *mux.Router, logger *zap.Logger, redisClient utils.IRedisUtil, kafkaClient utils.IKafkaUtil) {
	service := contracts.NewService(logger, redisClient, kafkaClient)
	handler := contracts.NewHandler(service, logger)
	router.HandleFunc("/contract", handler.CreateContract).Methods("POST")
}
