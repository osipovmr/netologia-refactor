package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"netologia-refactor/internal/db"
	"netologia-refactor/internal/repository"
	"netologia-refactor/internal/service"
)

// initDB отвечает только за инициализацию БД и миграции
func initDB(dbClient db.DB, dsn string) (*sql.DB, error) {
	database, err := dbClient.Open(dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := dbClient.Migrate(database); err != nil {
		// если миграции не прошли, соединение закрываем
		database.Close()
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return database, nil
}

// runApp отвечает только за бизнес-логику
func runApp(database *sql.DB) error {
	repo := repository.NewSQLiteOrderRepository(database)

	emailSender := service.NewEmailSender()
	emailOrderService := service.NewOrderService(repo, emailSender)

	slog.Info("=== Создание заказа с EmailSender ===")
	if err := emailOrderService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5); err != nil {
		return err
	}

	smsSender := service.NewSMSSender()
	smsOrderService := service.NewOrderService(repo, smsSender)

	slog.Info("=== Создание заказа с SMSSender ===")
	if err := smsOrderService.CreateOrder("Пётр", []string{"orange", "mango"}, 20.0); err != nil {
		return err
	}

	return nil
}

func main() {
	dbClient := db.NewSQLiteDB()
	database, err := initDB(dbClient, "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := runApp(database); err != nil {
		log.Fatal(err)
	}
}
