package main

import "fmt"

type Payment interface {
	Process()
}

type PaymentProcessor struct {
}

func (pp PaymentProcessor) ProcessPayment(payment Payment) {
	payment.Process()
}

type CreditCardPayment struct {
	Account string
	Amount  float64
}

func (ccp CreditCardPayment) Process() {
	fmt.Printf("Processing credit card payment of $%.2f from %s\n", ccp.Amount, ccp.Account)
	// Perform credit card payment specific logic
}

type DebitCardPayment struct {
	Account string
	Amount  float64
}

func (dbt DebitCardPayment) Process() {
	fmt.Printf("Processing debit card payment of $%.2f from %s\n", dbt.Amount, dbt.Account)
	// Perform debit card payment specific logic
}

type BankTransferPayment struct {
	Account string
	Amount  float64
}

func (btp BankTransferPayment) Process() {
	fmt.Printf("Processing bank transfer payment of $%.2f from %s\n", btp.Amount, btp.Account)
	// Perform bank transfer specific logic
}
