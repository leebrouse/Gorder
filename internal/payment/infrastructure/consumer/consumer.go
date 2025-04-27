package consumer

import (
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"

	"github.com/leebrouse/Gorder/common/broker"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/payment/app"
	"github.com/leebrouse/Gorder/payment/app/command"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Consumer struct {
	app app.Application
}

func NewConsumer(application app.Application) *Consumer {
	return &Consumer{
		app: application,
	}
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
		for meg := range megs {
			c.handleMessage(q, meg)
		}
	}()
	<-forever
}

// receive message function
func (c *Consumer) handleMessage(q amqp.Queue, meg amqp.Delivery) {
	logrus.Infof("Payment receive the message from %s,msg=%v ", q.Name, string(meg.Body))
	ctx := broker.ExtractRabbitMQHeaders(context.Background(), meg.Headers)
	tr := otel.Tracer("payment rabbitmq")
	_, span := tr.Start(ctx, fmt.Sprintf("rabbitmq.%s.consume", q.Name))
	defer span.End()

	o := &orderpb.Order{}
	if err := json.Unmarshal(meg.Body, o); err != nil {
		logrus.Infof("failed to unmarshall msg to order,err=%v", err)
		return
	}
	if _, err := c.app.Commend.CreatePayment.Handle(ctx, command.CreatePayment{Order: o}); err != nil {
		//TODO: retry
		logrus.Infof("failed to create order,err=%v", err)
		return
	}
	span.AddEvent("payment.created")
	logrus.Info("Consume success")
}
