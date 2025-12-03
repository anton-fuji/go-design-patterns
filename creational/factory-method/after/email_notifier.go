package main

import "fmt"

type EmailNotifier struct {
	senderAddress string
}

func newEmailNotifier() INotifier {
	return &EmailNotifier{
		senderAddress: "no-reply@company.com",
	}
}

func (en *EmailNotifier) Send(recipient string, message string) error {
	fmt.Printf("To: %s, From: %s, Message: %s\n", recipient, en.senderAddress, message)
	return nil
}

func (en EmailNotifier) GetChannelName() string {
	return "Email"
}
