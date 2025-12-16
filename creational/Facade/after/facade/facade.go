package facade

import (
	"fmt"

	"github.com/anton-fuji/go-design-patterns/creational/Facade/after/services"
)

type OrderFacade struct {
	inventory *services.InventorySystem
	payment   *services.PaymentSystem
	logistics *services.LogisticsSystem
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory: &services.InventorySystem{},
		payment:   &services.PaymentSystem{},
		logistics: &services.LogisticsSystem{},
	}
}

func (f *OrderFacade) PlaceOrder(productId string, amount int) error {
	fmt.Println("--- 注文プロセス開始 ---")

	// 在庫確認
	if !f.inventory.CheckStock(productId) {
		return fmt.Errorf("在庫がありません")
	}

	// 決済
	if err := f.payment.MakePayment(amount); err != nil {
		return fmt.Errorf("決済に失敗しました: %v", err)
	}

	f.logistics.ShipItem(productId)

	fmt.Println("--- 注文プロセス終了 ---")
	return nil
}
