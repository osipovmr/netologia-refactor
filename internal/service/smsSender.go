package service

import (
	"log/slog"
)

// SMSSender реализует Notifier
type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(customer string, message string) error {
	slog.Info("Отправка SMS уведомления клиенту ", "customer", customer, "message", message)
	return nil
}
