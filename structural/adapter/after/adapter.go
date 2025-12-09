package main

import "fmt"

// BankAdapter にLegacyBankを委譲させる
type BankAdapter struct {
	legacy *LegacyBank
}

func NewBankAdapter(l *LegacyBank) *BankAdapter {
	return &BankAdapter{legacy: l}
}

func (a *BankAdapter) Pay(amount float64) error {
	yen := int(amount)
	if err := a.legacy.Transfer(yen); err != nil {
		return fmt.Errorf("adapter transfer failed: %w", err)
	}
	return nil
}
