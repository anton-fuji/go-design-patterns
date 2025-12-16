package services

import "fmt"

type PaymentSystem struct{}

func (s *PaymentSystem) MakePayment(amount int) error {
	fmt.Printf("[決済] %d円 の支払いを処理中...\n", amount)
	return nil
}
