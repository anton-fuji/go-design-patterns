package main

import "fmt"

type SMSNotifier struct {
	provide string
}

func newSMSNotifier() INotifier {
	return &SMSNotifier{
		provide: "Hoge",
	}
}

func (sn SMSNotifier) Send(recipient string, message string) error {
	fmt.Printf("To: %s, From: %s, Message: %s\n", recipient, sn.provide, message)
	return nil
}

func (sn SMSNotifier) GetChannelName() string {
	return "SMS"
}
