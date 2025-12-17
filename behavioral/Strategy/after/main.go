package main

import (
	"fmt"

	"github.com/anton-fuji/go-design-patterns/behavioral/Strategy/after/cache"
)

func main() {
	fmt.Println("--- After Strategy Pattern ---")

	// LRUを使ってキャッシュを作成
	lru := &cache.Lru{}
	c := cache.NewCache(lru)

	c.Add("a", "1")
	c.Add("b", "2")

	// Capacity　Over
	c.Add("c", "3")

	fifo := &cache.Fifo{}
	c.SetEvictionAlgo(fifo)

	// FIFOに切り替わる
	c.Add("d", "4")

	lfu := &cache.Lfu{}
	c.SetEvictionAlgo(lfu)
	c.Add("e", "5")

}
