package consumer

import (
	"github.com/leebrouse/Gorder/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Listen(ch *amqp.Channel) {
	//declare the queue
	q, err := ch.QueueDeclare(broker.EventOrderCreated, true, false, false, false, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	//consume the message
	megs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logrus.Warnf("fail to consume,queue=%s,err=%v ", q.Name, err)
	}

	// loop for always reading the message
	var forever chan struct{}
	go func() {
		for msg := range megs {
			c.handleMessage(ch, q, msg)
		}
	}()
	<-forever
}

// receive message function
func (c *Consumer) handleMessage(ch *amqp.Channel, q amqp.Queue, meg amqp.Delivery) {
	logrus.Infof("Payment receive the message from %s,msg=%v ", q.Name, string(meg.Body))
}
