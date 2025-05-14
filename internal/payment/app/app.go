package app

import "github.com/leebrouse/Gorder/payment/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreatePayment command.CreatePaymentHandler
}
