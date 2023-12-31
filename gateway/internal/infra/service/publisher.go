package service

import (
	"context"
	"fmt"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewPublisher(cfg *config.Config) (service.Publisher, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.MQUser, cfg.MQPassword, cfg.MQHost, cfg.MQPort))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.MQQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &publisherImpl{
		cfg: cfg,
		q:   q,
	}, nil
}

type publisherImpl struct {
	cfg *config.Config
	q   amqp.Queue
}

func (p *publisherImpl) Publish(ctx context.Context, contentType string, body []byte) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", p.cfg.MQUser, p.cfg.MQPassword, p.cfg.MQHost, p.cfg.MQPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err := ch.PublishWithContext(
		ctx,
		"",
		p.q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  contentType,
			Body:         body,
		},
	); err != nil {
		return err
	}

	return nil
}
