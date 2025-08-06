package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	log    *slog.Logger
}

func NewProducer(brokers []string, topic string, clientID string, log *slog.Logger) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
			Async:    true,
		},
		log: log,
	}
}

func (p *Producer) Publish(ctx context.Context, key string, event interface{}) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.log.Error("failed to marshal event", slog.Any("error", err))
		return err
	}

	err = p.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: value,
		},
	)
	if err != nil {
		p.log.Error("failed to write message to kafka", slog.Any("error", err))
		return err
	}

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
