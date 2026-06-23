package service

import (
	"fmt"
	"netologia-refactor/internal/model"
)

// OrderService зависит только от RepositoryWriter и Notifier
type OrderService struct {
	repo     RepositoryWriter
	notifier Notifier
}

func NewOrderService(repo RepositoryWriter, notifier Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := model.Order{
		Customer: customer,
		Products: fmt.Sprintf("%v", products),
		Total:    total,
		Status:   "pending",
	}

	if err := s.repo.CreateOrder(order); err != nil {
		return err
	}

	msg := fmt.Sprintf("Заказ успешно создан на сумму %.2f", total)
	if err := s.notifier.Send(customer, msg); err != nil {
		return err
	}

	return nil
}
