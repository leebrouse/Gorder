package app

import "github.com/leebrouse/Gorder/payment/app/command"

// Application
type Application struct {
	//	commend
	Commend Commend
}

// Commend
type Commend struct {
	CreatePayment command.CreatePaymentHandler
}
