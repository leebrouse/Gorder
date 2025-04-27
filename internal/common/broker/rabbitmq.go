package broker

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// Connect with the rabbitmq
func Connect(user, password, host, port string) (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port)
	//connect with the rabbitmq
	conn, err := amqp.Dial(address)
	if err != nil {
		logrus.Fatal(err)
	}

	//create channel
	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatal(err)
	}

	// Event:  OrderCreated  kind:direct
	err = orderCreatedExchangeDeclare(ch)
	if err != nil {
		logrus.Fatal(err)
	}

	// Event: OrderPaid  kind:fanout
	err = orderPaidExchangeDeclare(ch)
	if err != nil {
		logrus.Fatal(err)
	}

	return ch, conn.Close
}

/** Exchange Declare,direct:点对点  fanout:广播消息 **/
// Event:  OrderCreated
func orderCreatedExchangeDeclare(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		EventOrderCreated,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// Event: OrderPaid
func orderPaidExchangeDeclare(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		EventOrderPaid,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

type RabbitMQHeaderCarrier map[string]any

// implement the TextMapCarrier interface
func (r RabbitMQHeaderCarrier) Get(key string) string {
	value, ok := r[key]
	if !ok {
		return ""
	}
	return value.(string)
}

func (r RabbitMQHeaderCarrier) Set(key string, value string) {
	r[key] = value
}

func (r RabbitMQHeaderCarrier) Keys() []string {
	keys := make([]string, len(r))

	for key := range r {
		keys = append(keys, key)
	}
	return keys
}

func InjectRabbitMQHeaders(ctx context.Context) map[string]any {
	carrier := make(RabbitMQHeaderCarrier)
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return carrier
}

func ExtractRabbitMQHeaders(ctx context.Context, headers map[string]any) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, RabbitMQHeaderCarrier(headers))
}
