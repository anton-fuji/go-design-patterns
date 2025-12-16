package services

import "fmt"

type LogisticsSystem struct{}

func (s *LogisticsSystem) ShipItem(productId string) {
	fmt.Printf("[配送] 商品ID: %s の発送準備完了\n", productId)
}
