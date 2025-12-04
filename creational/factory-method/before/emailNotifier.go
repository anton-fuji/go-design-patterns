package main

import "fmt"

type EmailNotifer struct {
	senderAddress string
}

// afterのINotifierインターフェースのメソッドをそのまま実装
func (en EmailNotifer) Send(recipient string, message string) error {
	fmt.Printf("To: %s, From: %s, Message: %s\n", recipient, en.senderAddress, message)
	return nil
}

func (en EmailNotifer) GetChannelName() string {
	return "Email"
}
