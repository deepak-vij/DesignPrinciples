package main

import "fmt"

type PayPal struct {
	Account string
	Amount  float64
}

func (ppl PayPal) Process() {
	fmt.Printf("Processing PayPal payment of $%.2f from %s\n", ppl.Amount, ppl.Account)
	// Perform PayPal payment specific logic
}

func main() {
	paymentProcessor := PaymentProcessor{}
	creditCardPayment := CreditCardPayment{Account: "c123456789", Amount: 100.0}
	paymentProcessor.ProcessPayment(creditCardPayment)

	debitCardPayment := DebitCardPayment{Account: "d123456789", Amount: 500.0}
	paymentProcessor.ProcessPayment(debitCardPayment)

	bankTransferPayment := BankTransferPayment{Account: "987654321", Amount: 200.0}
	paymentProcessor.ProcessPayment(bankTransferPayment)

	payPal := PayPal{Account: "d-2123456789", Amount: 700.0}
	paymentProcessor.ProcessPayment(payPal)

}
