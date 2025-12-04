package main

import "fmt"

type SMSNotifier struct {
	provider string
}

// afterのINotifierインターフェースのメソッドをそのまま実装
func (sn *SMSNotifier) Send(recipient string, message string) error {
	fmt.Printf("To: %s, Via: %s, Message: %s\n", recipient, sn.provider, message)
	return nil
}

func (sn *SMSNotifier) GetChannelName() string {
	return "SMS"
}
