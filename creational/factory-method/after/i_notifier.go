package main

type INotifier interface {
	Send(recipient string, message string) error
	GetChannelName() string
}
