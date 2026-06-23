package main

import (
	"fmt"
	"log"
	"netologia-refactor/internal/db"
	"netologia-refactor/internal/repository"
	"netologia-refactor/internal/service"
)

func main() {
	database, err := db.Open("orders.db")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer database.Close()

	if err := db.Migrate(database); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	repo := repository.NewSQLiteOrderRepository(database)

	emailSender := service.NewEmailSender()
	emailOrderService := service.NewOrderService(repo, emailSender)

	fmt.Println("=== Создание заказа с EmailSender ===")
	if err := emailOrderService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5); err != nil {
		log.Fatal(err)
	}

	smsSender := service.NewSMSSender()
	smsOrderService := service.NewOrderService(repo, smsSender)

	fmt.Println("=== Создание заказа с SMSSender ===")
	if err := smsOrderService.CreateOrder("Пётр", []string{"orange", "mango"}, 20.0); err != nil {
		log.Fatal(err)
	}
}
