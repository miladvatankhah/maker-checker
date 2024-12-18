package events

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQEventPublisher struct {
	channel *amqp.Channel
}

func NewRabbitMQEventPublisher(channel *amqp.Channel) *RabbitMQEventPublisher {
	return &RabbitMQEventPublisher{channel: channel}
}

func (p *RabbitMQEventPublisher) Publish(event interface{}) {
	body, err := json.Marshal(event)
	if err != nil {
		return
	}

	p.channel.Publish(
		"",
		"message_approved_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
