package contracts

import (
	"context"
	"encoding/json"

	"github.com/ashishbabar/go-eth-api-contract-service/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type IService interface {
	Create(contract Contract) error
}

type service struct {
	logger    *zap.Logger
	redisUtil utils.IRedisUtil
	kafkaUtil utils.IKafkaUtil
}

func NewService(logger *zap.Logger, redisUtil utils.IRedisUtil, kafkaUtil utils.IKafkaUtil) IService {
	return &service{
		logger:    logger,
		redisUtil: redisUtil,
		kafkaUtil: kafkaUtil,
	}
}

func (ser *service) Create(contract Contract) error {
	ctx := context.Background()
	val := ser.redisUtil.GetVal(ctx, viper.GetString("ContractID"))
	ser.logger.Info("Retrieved new contract ID from redis : " + val)

	contract.ID = val
	jsonContract, err := json.Marshal(contract)
	if err != nil {
		ser.logger.Error(err.Error())
		return err
	}
	ser.logger.Info("Marshelled contract in json string")
	ser.kafkaUtil.Produce(jsonContract)
	ser.logger.Info("Produced message to kafka")
	return nil
}
