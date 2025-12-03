package main

import (
	"fmt"
)

func main() {
	emailService, err := GetNotifier(TypeEmail)
	if err != nil {
		fmt.Errorf("error: %w", err)
		return
	}
	smsService, err := GetNotifier(TypeSMS)
	if err != nil {
		fmt.Errorf("error: %w", err)
		return
	}

	sendNotification(emailService, "user@example.com", "ご注文の商品が発送されました")
	println()
	sendNotification(smsService, "+7-90xxxxxxxxxxxxx", "残高が少なくなっています")

}

func sendNotification(n INotifier, recipient string, message string) {
	fmt.Printf("--- Sending Via %s ----\n", recipient)
	if err := n.Send(recipient, message); err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	}
}
