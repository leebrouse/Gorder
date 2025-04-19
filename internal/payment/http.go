package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/Gorder/common/broker"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/payment/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

type PaymentHandler struct {
	channel *amqp.Channel
}

func NewPaymentHandler(ch *amqp.Channel) *PaymentHandler {
	return &PaymentHandler{channel: ch}
}

// router
func (h *PaymentHandler) RegisterRoutes(c *gin.Engine) {
	c.POST("/api/webhook", h.handleWebhook)
}

// handler
func (h *PaymentHandler) handleWebhook(c *gin.Context) {
	logrus.Info("Got webhook from stripe")

	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Infof("Error reading request body: %v\n", err)
		c.JSON(http.StatusServiceUnavailable, err)
		return
	}

	event, err := webhook.ConstructEvent(payload, c.Request.Header.Get("Stripe-Signature"), viper.GetString("endpoint-stripe-secret"))

	if err != nil {
		logrus.Infof("Error verifying webhook signature: %v\n", err)
		c.JSON(http.StatusBadRequest, err) // Return a 400 error on a bad signature
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case stripe.EventTypeCheckoutSessionCompleted:
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			logrus.Infof("error unmarshal event.data.raw into session,err=%v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		// pay success and send message to the rabbitmq
		if session.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
			logrus.Infof("payment for checkout session %v success!", session.ID)

			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			var items []*orderpb.Item
			_ = json.Unmarshal([]byte(session.Metadata["items"]), &items)

			marshalledOrder, err := json.Marshal(&domain.Order{
				ID:          session.Metadata["orderID"],
				CustomerID:  session.Metadata["customerID"],
				Status:      string(stripe.CheckoutSessionPaymentStatusPaid),
				PaymentLink: session.Metadata["paymentLink"],
				Items:       items,
			})
			if err != nil {
				logrus.Infof("error marshal domain.order ,err=%v", err)
				c.JSON(http.StatusBadRequest, err)
				return
			}

			_ = h.channel.PublishWithContext(
				ctx,
				broker.EventOrderPaid,
				"",
				false,
				false,
				amqp.Publishing{
					ContentType:  "application/json",
					DeliveryMode: amqp.Persistent,
					Body:         marshalledOrder,
				})
			logrus.Infof("message published to %s,body:%s", broker.EventOrderPaid, string(marshalledOrder))
		}
	}
	c.JSON(http.StatusOK, nil)
}
