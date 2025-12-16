package services

import "fmt"

type InventorySystem struct{}

func (s *InventorySystem) CheckStock(productId string) bool {
	fmt.Printf("[在庫] 商品ID: %s の在庫確認中・・・\n", productId)
	// 実際はDBなどのロジックが入る
	return true
}
