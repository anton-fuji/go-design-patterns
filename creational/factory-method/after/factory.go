package main

import "fmt"

const (
	TypeEmail = "email"
	TypeSMS   = "sms"
)

func GetNotifier(notifierType string) (INotifier, error) {
	if notifierType == TypeEmail {
		return newEmailNotifier(), nil
	}
	if notifierType == TypeSMS {
		return newSMSNotifier(), nil
	}
	return nil, fmt.Errorf("unsupported notifier type: %s", notifierType)
}
