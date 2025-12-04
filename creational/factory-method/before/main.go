package main

import "fmt"

func main() {
	// Before: 具象クラスを直接インスタンス化し、具体的な型に依存
	emailService := EmailNotifer{
		senderAddress: "no-reply@company.com", // 初期化ロジックがクライアントに露出
	}

	smsService := &SMSNotifier{
		provider: "Hoge", // 初期化ロジックがクライアントに露出
	}

	// Before: 具象型ごとの関数を呼び出すか、呼び出し側のロジックが複雑化
	// (ここでは、具象型を引数に取る関数を定義したと仮定)
	sendEmailNotification(emailService, "user@example.com", "ご注文の商品が発送されました")
	println()
	sendSMSNotification(smsService, "+7-90xxxxxxxxxxxxx", "残高が少なくなっています")
}

// 具象型 EmailNotifier を引数に取る
func sendEmailNotification(n EmailNotifer, recipient string, message string) {
	fmt.Printf("--- Sending Via %s ----\n", n.GetChannelName()) // 具象メソッド呼び出し
	if err := n.Send(recipient, message); err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	}
}

// 具象型 *SMSNotifier を引数に取る
func sendSMSNotification(n *SMSNotifier, recipient string, message string) {
	fmt.Printf("--- Sending Via %s ----\n", n.GetChannelName()) // 具象メソッド呼び出し
	if err := n.Send(recipient, message); err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	}
}
