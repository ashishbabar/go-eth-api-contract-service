package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type IKafkaUtil interface {
	Produce(data []byte) bool
}
type kafkaUtil struct {
	producer *kafka.Producer
}

func NewKafkaUtil() (IKafkaUtil, error) {

	conf := readConfig("./kafka.config")
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return nil, err
	}
	return &kafkaUtil{producer: p}, nil
}

func readConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}

func (p kafkaUtil) Produce(data []byte) bool {
	topic := viper.GetString("DeploymentTopic")
	delivery_chan := make(chan kafka.Event, 10000)
	p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data},
		delivery_chan,
	)
	e := <-delivery_chan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(delivery_chan)
	return true
}
