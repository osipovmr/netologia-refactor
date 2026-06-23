package service

import "fmt"

// SMSSender реализует Notifier
type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(customer string, message string) error {
	fmt.Printf("Отправка SMS уведомления клиенту %s: %s\n", customer, message)
	return nil
}
