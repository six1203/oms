package global

import (
	"github.com/Shopify/sarama"
	"time"
)

func NewKafkaProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = time.Duration(500) * time.Millisecond
	return config
}
