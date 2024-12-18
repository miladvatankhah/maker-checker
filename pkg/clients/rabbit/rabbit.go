package rabbit

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Config struct {
	Url            string
	Heartbeat      int `mapstructure:"heartbeat"`
	ReconnectDelay int `mapstructure:"reconnectDelay"`
	MaxRetries     int `mapstructure:"maxRetries"`
}

type ConsumeOptions struct {
	QueueOptions
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      map[string]interface{}
}

type PublishOptions struct {
	QueueOptions
	ContentType  string
	DeliveryMode uint8
	Exchange     string
	Mandatory    bool
	Immediate    bool
}

type QueueOptions struct {
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       map[string]interface{}
}

const JsonContentType = "application/json"

type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	cfg        Config
}

var rabbit *Rabbit

func Dial(cfg Config) (*Rabbit, error) {
	if rabbit != nil {
		return rabbit, nil
	}

	rabbit = &Rabbit{cfg: cfg}
	if err := rabbit.connect(); err != nil {
		return nil, err
	}

	return rabbit, nil
}

func DialWithDefaults(cfg Config) (*Rabbit, error) {
	if rabbit != nil {
		return rabbit, nil
	}

	rabbit = &Rabbit{cfg: cfg}
	if err := rabbit.connect(); err != nil {
		return nil, err
	}

	return rabbit, nil
}

func (r *Rabbit) EnsureConnection() error {
	if r.connection != nil && !r.connection.IsClosed() {
		return nil
	}

	for attempt := 1; attempt <= r.cfg.MaxRetries; attempt++ {
		if err := r.connect(); err != nil {
			delay := time.Duration(r.cfg.ReconnectDelay) * time.Duration(attempt)
			time.Sleep(delay)
			continue
		}

		return nil
	}

	return errors.New("failed to reconnect after max retries")
}

func (r *Rabbit) Consume(queueName string, opts ConsumeOptions) (<-chan amqp.Delivery, error) {
	if err := r.EnsureConnection(); err != nil {
		return nil, err
	}

	queue, err := r.createQueue(queueName, opts.QueueOptions)
	if err != nil {
		return nil, err
	}

	return r.channel.Consume(
		queue.Name,
		opts.Consumer,
		opts.AutoAck,
		opts.Exclusive,
		opts.NoLocal,
		opts.NoWait,
		opts.Args,
	)
}

func (r *Rabbit) ConsumeWithDefaults(queueName string) (<-chan amqp.Delivery, error) {
	if err := r.EnsureConnection(); err != nil {
		return nil, err
	}

	consumeOpts := ConsumeOptions{QueueOptions: QueueOptions{durable: true}}
	queue, err := r.createQueue(queueName, consumeOpts.QueueOptions)
	if err != nil {
		return nil, err
	}

	return r.channel.Consume(
		queue.Name,
		consumeOpts.Consumer,
		consumeOpts.AutoAck,
		consumeOpts.Exclusive,
		consumeOpts.NoLocal,
		consumeOpts.NoWait,
		consumeOpts.Args,
	)
}

func (r *Rabbit) Publish(queueName string, body []byte, opts PublishOptions) error {
	if err := r.EnsureConnection(); err != nil {
		return err
	}

	queue, err := r.createQueue(queueName, opts.QueueOptions)
	if err != nil {
		return err
	}

	message := amqp.Publishing{
		ContentType:  opts.ContentType,
		Body:         body,
		DeliveryMode: opts.DeliveryMode,
	}

	return r.channel.Publish(
		opts.Exchange,
		queue.Name,
		opts.Mandatory,
		opts.Immediate,
		message,
	)
}

func (r *Rabbit) PublishWithDefaults(queueName string, body []byte) error {
	if err := r.EnsureConnection(); err != nil {
		return err
	}

	publishOpts := PublishOptions{QueueOptions: QueueOptions{durable: true}, ContentType: JsonContentType, DeliveryMode: 2}
	queue, err := r.createQueue(queueName, publishOpts.QueueOptions)
	if err != nil {
		return err
	}

	message := amqp.Publishing{
		ContentType:  publishOpts.ContentType,
		Body:         body,
		DeliveryMode: publishOpts.DeliveryMode,
	}

	return r.channel.Publish(
		publishOpts.Exchange,
		queue.Name,
		publishOpts.Mandatory,
		publishOpts.Immediate,
		message,
	)
}

func (r *Rabbit) connect() error {
	conn, err := amqp.DialConfig(r.cfg.Url, amqp.Config{Heartbeat: time.Duration(r.cfg.Heartbeat)})
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	r.connection = conn
	r.channel = ch

	return nil
}

func (r *Rabbit) createQueue(queueName string, opts QueueOptions) (*amqp.Queue, error) {
	queue, err := r.channel.QueueDeclare(queueName, opts.durable, opts.autoDelete, opts.exclusive, opts.noWait, opts.args)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func (r *Rabbit) Chan() *amqp.Channel {
	return r.channel
}

func (r *Rabbit) Shutdown() error {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			return err
		}
	}

	if r.connection != nil {
		if err := r.connection.Close(); err != nil {
			return err
		}
	}

	return nil
}
