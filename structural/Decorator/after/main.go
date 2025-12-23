package main

import "fmt"

func main() {
	// シンプルなコーヒー
	coffee := NewCoffee("Espresso")
	fmt.Printf("%s: %d円\n", coffee.Description(), coffee.Cost())

	// ミルク入り
	coffeeWithMilk := NewMilk(NewCoffee("Espresso"))
	fmt.Printf("%s: %d円\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	// ミルクと砂糖入り
	coffeeWithMilkAndSugar := NewSugar(NewMilk(NewCoffee("Espresso")))
	fmt.Printf("%s: %d円\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())

	// ダブルミルク、砂糖、ホイップ入り
	fancyCoffee := NewWhip(
		NewSugar(
			NewMilk(
				NewMilk(
					NewCoffee("Espresso"),
				),
			),
		),
	)
	fmt.Printf("%s: %d円\n", fancyCoffee.Description(), fancyCoffee.Cost())
	// 様々な組み合わせに対応できる
}
