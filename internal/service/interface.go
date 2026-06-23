package service

import "netologia-refactor/internal/model"

// RepositoryWriter определяет контракт записи заказа
type RepositoryWriter interface {
	CreateOrder(order model.Order) error
}

// Notifier определяет контракт отправки уведомлений
type Notifier interface {
	Send(customer string, message string) error
}
