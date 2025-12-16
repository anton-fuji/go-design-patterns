package main

import (
	"fmt"
)

// --- サブシステム ---
type InventorySystem struct{}

func (s *InventorySystem) CheckStock(productID string) bool {
	fmt.Printf("[在庫] 商品ID: %s の在庫確認中...\n", productID)
	return true
}

type PaymentSystem struct{}

func (s *PaymentSystem) MakePayment(amount int) error {
	fmt.Printf("[決済] %d円 の支払いを処理中...\n", amount)
	return nil
}

type LogisticsSystem struct{}

func (s *LogisticsSystem) ShipItem(productID string) {
	fmt.Printf("[配送] 商品ID: %s の発送準備完了。\n", productID)
}

func main() {
	// 【問題1】
	// クライアントが全てのサブシステムを初期化する必要がある
	inventory := &InventorySystem{}
	payment := &PaymentSystem{}
	logistics := &LogisticsSystem{}

	productID := "item-123"
	amount := 5000

	fmt.Println("--- 注文プロセス開始 (手動) ---")

	// 【問題2】
	// クライアントが「正しい呼び出し順序」と「エラー処理」を
	// 全て記述しなければならない（ビジネスロジックが露出している）

	// Step 1: 在庫確認
	if !inventory.CheckStock(productID) {
		fmt.Println("エラー: 在庫がありません")
		return
	}

	// Step 2: 決済
	if err := payment.MakePayment(amount); err != nil {
		fmt.Println("エラー: 決済に失敗しました", err)
		return
	}

	// Step 3: 配送
	logistics.ShipItem(productID)

	fmt.Println("--- 注文プロセス完了 ---")
}
