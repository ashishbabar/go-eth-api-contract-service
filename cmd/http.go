package main

import (
	"context"

	"github.com/ashishbabar/go-eth-api-contract-service/server"
	"github.com/ashishbabar/go-eth-api-contract-service/utils"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func main() {
	logger := utils.ZapLogger
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Error reading env file " + err.Error())
	}

	httpServer := server.NewServer(logger)
	logger.Info("Created httpServer instance")
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
