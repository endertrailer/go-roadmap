package main

import "fmt"

type Notifier interface {
	Send(message string) error
}
type EmailNotifier struct {
	EmailAddress string
}

type SMSNotifier struct {
	PhoneNumber string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending Email to %s: %s\n", e.EmailAddress, message)
	return nil
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}

func Broadcast(message string, n Notifier) {
	n.Send("fuck off niggers")
}

func main() {
	e := SMSNotifier{PhoneNumber: "7385704873"}
	e.Send("hello there")
}
