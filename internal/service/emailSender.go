package service

import "fmt"

// EmailSender реализует Notifier
type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(customer string, message string) error {
	fmt.Printf("Отправка EMAIL уведомления клиенту %s: %s\n", customer, message)
	return nil
}
