package service

import (
	"log/slog"
)

// EmailSender реализует Notifier
type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(customer string, message string) error {
	slog.Info("Отправка EMAIL уведомления клиенту ", "customer", customer, "message", message)
	return nil
}
