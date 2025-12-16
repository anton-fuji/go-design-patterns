package main

import (
	"fmt"

	"github.com/anton-fuji/go-design-patterns/creational/Facade/after/facade"
)

func main() {
	facade := facade.NewOrderFacade()

	productID := "item-123"
	amount := 5000

	err := facade.PlaceOrder(productID, amount)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
