package command

import (
	"github.com/leebrouse/Gorder/common/decorator"
	"github.com/leebrouse/Gorder/common/genproto/orderpb"
	"github.com/leebrouse/Gorder/payment/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CreatePayment struct {
	Order *orderpb.Order
}

type CreatePaymentHandler decorator.QueryHandler[CreatePayment, string]

type createPaymentHandler struct {
	processor domain.Processor
	orderGRPC OrderService
}

func (c createPaymentHandler) Handle(ctx context.Context, cmd CreatePayment) (string, error) {
	//new span in payment service
	//tr := otel.Tracer("payment")
	//ctx, _ = tr.Start(ctx, "create_payment")

	//create payment link
	link, err := c.processor.CreatePaymentLink(ctx, cmd.Order)
	if err != nil {
		return "", err
	}

	//log
	logrus.Infof("Create payment link for order: %s success,payment link: %s", cmd.Order.ID, link)
	//update order status
	newOrder := &orderpb.Order{
		ID:          cmd.Order.ID,
		CustomerID:  cmd.Order.CustomerID,
		Status:      "waiting_for_payment",
		Items:       cmd.Order.Items,
		PaymentLink: link,
	}

	err = c.orderGRPC.UpdateOrder(ctx, newOrder)
	if err != nil {
		return "", err
	}
	return link, err
}

func NewCreatePaymentHandler(
	processor domain.Processor,
	orderGRPC OrderService,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CreatePaymentHandler {
	return decorator.ApplyCommandDecorators[CreatePayment, string](
		createPaymentHandler{processor: processor, orderGRPC: orderGRPC},
		logger,
		metricClient,
	)
}
