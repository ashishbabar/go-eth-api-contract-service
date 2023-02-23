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
	newContractID, err := ser.redisUtil.GetVal(ctx, viper.GetString("RedisContractIDKey"))
	if err != nil {
		ser.logger.Error(err.Error())
		return err
	}
	ser.logger.Info("Retrieved new contract ID from redis : " + newContractID)

	contract.ID = newContractID
	jsonContract, err := json.Marshal(contract)
	if err != nil {
		ser.logger.Error(err.Error())
		return err
	}
	ser.logger.Info("Marshelled contract in json string ", zap.Any("JsonContract", jsonContract))
	ser.kafkaUtil.Produce(jsonContract)
	ser.logger.Info("Produced message to kafka")
	return nil
}
