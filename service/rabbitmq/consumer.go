package rabbitmq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"user-service/payload"
	"user-service/usecase"
	"user-service/util"
)

var ListQueue = []string{
	util.UPDATE_WALLET_AMOUNT,
}

type Consumer struct {
	Conn    *amqp.Connection
	UseCase *usecase.UseCase
}

func NewConsumer(c *amqp.Connection, uc *usecase.UseCase) *Consumer {
	return &Consumer{
		Conn:    c,
		UseCase: uc,
	}
}

func (c *Consumer) StartConsumer() {
	for _, queue := range ListQueue {
		go c.Consume(queue)
	}
}

func (c *Consumer) Consume(queueName string) {
	ch, err := c.Conn.Channel()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	var forever chan struct{}

	go func() {
		for m := range msgs {
			switch queueName {
			case util.UPDATE_WALLET_AMOUNT:
				log.Println("Consume queue" + util.UPDATE_WALLET_AMOUNT)
				var req payload.UpdateBalanceReq
				_ = json.Unmarshal(m.Body, &req)
				err := c.UseCase.WalletUseCase.Payment(req.Id, req.Amount)
				if err != nil {
					logrus.Error(err)
				}
			}
		}
	}()
	<-forever
	return
}

func (c *Consumer) Close() {
	c.Conn.Close()
}
