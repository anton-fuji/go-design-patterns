package main

import (
	"errors"
	"fmt"
)

type LegacyBank struct{}

func (b *LegacyBank) Transfer(yen int) error {
	if yen <= 0 {
		return errors.New("amount must be positive")
	}
	fmt.Printf("[Legacy] Transferred: %d JPY\n", yen)
	return nil
}
