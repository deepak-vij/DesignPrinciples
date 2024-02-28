package main

import "fmt"

type Payment struct {
    Type    string
    Account string
    Amount  float64
}

func ProcessPayment(payment Payment) {
    switch payment.Type {
    case "CreditCard":
        // Process credit card payment
        fmt.Printf("Processing credit card payment of $%.2f from %s\n", payment.Amount, payment.Account)
        // Perform credit card payment specific logic
    case "BankTransfer":
        // Process bank transfer payment
        fmt.Printf("Processing bank transfer payment of $%.2f from %s\n", payment.Amount, payment.Account)
        // Perform bank transfer specific logic
    default:
        fmt.Println("Invalid payment type")
    }
}