package main

import "fmt"

type Coffee struct {
	name string
}

func (c *Coffee) Cost() int {
	return 300
}

func (c *Coffee) Description() string {
	return c.name
}

// ミルク入り
type CoffeeWithMilk struct {
	name string
}

func (c *CoffeeWithMilk) Cost() int {
	return 300 + 50
}

func (c *CoffeeWithMilk) Description() string {
	return c.name + " + Milk"
}

// 砂糖入りコーヒー
type CoffeeWithSugar struct {
	name string
}

func (c *CoffeeWithSugar) Cost() int {
	return 300 + 30
}

func (c *CoffeeWithSugar) Description() string {
	return c.name + " + Sugar"
}

// ミルクと砂糖入りコーヒー
type CoffeeWithMilkAndSugar struct {
	name string
}

func (c *CoffeeWithMilkAndSugar) Cost() int {
	return 300 + 50 + 30
}

func (c *CoffeeWithMilkAndSugar) Description() string {
	return c.name + " + Milk + Sugar"
}

// ホイップクリーム入りコーヒー
type CoffeeWithWhip struct {
	name string
}

func (c *CoffeeWithWhip) Cost() int {
	return 300 + 70
}

func (c *CoffeeWithWhip) Description() string {
	return c.name + " + Whip"
}

// 問題点: トッピングの組み合わせごとに構造体を作る必要がある
// ミルク+砂糖+ホイップの場合は？
// チョコ、キャラメルなど新しいトッピングを追加したら？
// 組み合わせの処理が煩雑になる

func main() {
	coffee := &Coffee{name: "Espresso"}
	fmt.Printf("%s: %d円\n", coffee.Description(), coffee.Cost())

	coffeeWithMilk := &CoffeeWithMilk{name: "Espresso"}
	fmt.Printf("%s: %d円\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	coffeeWithMilkAndSugar := &CoffeeWithMilkAndSugar{name: "Espresso"}
	fmt.Printf("%s: %d円\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())
}
