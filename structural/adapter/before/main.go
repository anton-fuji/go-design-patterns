package main

import (
	"fmt"
	"log"
)

func DirectPayment(bank *LegacyBank, amounts []float64) error {
	errs := make([]error, 0, len(amounts))

	for _, amt := range amounts {
		yen := int(amt)
		if err := bank.Transfer(yen); err != nil {
			wrappedErr := fmt.Errorf("transfer failed for %.2f: %w", amt, err)
			errs = append(errs, wrappedErr)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("process finished with errors: %w", errs[0])
	}
	return nil
}

func main() {
	bank := &LegacyBank{}
	amounts := []float64{100.50, 0, 500.0}

	if err := DirectPayment(bank, amounts); err != nil {
		log.Printf("Error: %v", err)
	}

}
