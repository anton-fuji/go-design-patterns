package main

import (
	"fmt"
	"log"
)

/*
インターフェースのみに依存するようにする
渡されるのがadapter なのか最新のAPIなのか知らなくても大丈夫
*/
func DecoupledPaymentProcess(p PaymentProcessor, amounts []float64) error {
	errs := make([]error, 0, len(amounts))

	for _, amt := range amounts {
		// ここで「intへキャスト」などのロジックに依存させない実装ができている
		if err := p.Pay(amt); err != nil {
			errs = append(errs, fmt.Errorf("payment %.2f failed: %w", amt, err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("process completed with errors: %w", errs[0])
	}
	return nil
}

func main() {
	legacy := &LegacyBank{}

	adapter := NewBankAdapter(legacy)

	amounts := []float64{100.50, 0, 5000.0}

	if err := DecoupledPaymentProcess(adapter, amounts); err != nil {
		log.Printf("Error: %v", err)
	}
}
