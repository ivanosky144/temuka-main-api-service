package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	conn *amqp.Connection
	url  string
}

func NewRabbitMQConnection(rabbitmqUrl string) (*RabbitMQConnection, error) {
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		return nil, err
	}

	wrapper := &RabbitMQConnection{
		conn: conn,
		url:  rabbitmqUrl,
	}

	go wrapper.handleReconnect()

	log.Println("Connected to RabbitMQ")
	return wrapper, nil
}

func (r *RabbitMQConnection) handleReconnect() {
	for {
		err, ok := <-r.conn.NotifyClose(make(chan *amqp.Error))
		if !ok {
			log.Println("RabbitMQ connection closed normally")
			return
		}

		log.Printf("RabbitMQ connection lost: %v", err)

		for {
			time.Sleep(5 * time.Second)
			log.Println("Reconnecting to RabbitMQ...")

			conn, err := amqp.Dial(r.url)
			if err == nil {
				r.conn = conn
				log.Println("RabbitMQ reconnected successfully!")
				break
			}

			log.Printf("Reconnect failed: %v", err)
		}
	}
}

func (r *RabbitMQConnection) Channel() (*RabbitMQChannel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}

	wrapper := &RabbitMQChannel{
		ch: ch,
	}

	go wrapper.handleChannelClose(r)
	return wrapper, nil
}

type RabbitMQChannel struct {
	ch *amqp.Channel
}

func (c *RabbitMQChannel) handleChannelClose(r *RabbitMQConnection) {
	err, ok := <-c.ch.NotifyClose(make(chan *amqp.Error))
	if !ok {
		return
	}

	log.Printf("Channel closed: %v", err)

	for {
		time.Sleep(2 * time.Second)

		ch, err := r.conn.Channel()
		if err == nil {
			c.ch = ch
			log.Println("RabbitMQ channel recreated successfully")
			return
		}

		log.Printf("Failed to recreate channel: %v", err)
	}
}

func (c *RabbitMQChannel) Close() error {
	return c.ch.Close()
}

func (c *RabbitMQChannel) RegisterExchange(exchangeName, exchangeType string, durable, autoDelete bool) error {
	err := c.ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		durable,
		autoDelete,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare exchange: %w", err)
	}

	return nil
}

func (c *RabbitMQChannel) InitQueue(exchangeName, routingKey string, durable, autoDelete bool) (amqp.Queue, error) {
	queue, err := c.ch.QueueDeclare(
		routingKey, // queue name (using routingKey as queue name)
		durable,    // durable
		autoDelete, // auto-delete
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare queue: %w", err)
	}

	if exchangeName != "" {
		err = c.ch.QueueBind(
			queue.Name,   // queue name
			routingKey,   // routing key
			exchangeName, // exchange name
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			return amqp.Queue{}, fmt.Errorf("failed to bind queue: %w", err)
		}
	}

	return queue, nil
}

func (c *RabbitMQChannel) PublishMessage(exchangeName, routingKey string, body interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal message body: %w", err)
	}

	err = c.ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         payload,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
