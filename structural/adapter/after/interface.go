package main

type PaymentProcessor interface {
	Pay(amount float64) error
}
