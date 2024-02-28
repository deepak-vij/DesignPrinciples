package main

func main() {
	payment1 := Payment{Type: "CreditCard", Account: "123456789", Amount: 100.0}
	ProcessPayment(payment1)

	payment2 := Payment{Type: "BankTransfer", Account: "987654321", Amount: 200.0}
	ProcessPayment(payment2)

	payment3 := Payment{Type: "PayPal", Account: "BTC123456", Amount: 50.0}
	ProcessPayment(payment3)
}
