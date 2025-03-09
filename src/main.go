package main

import (
	"warehouse-service/helpers"
	"warehouse-service/httphandler"
	"warehouse-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
	migrator "github.com/IvanSkripnikov/go-migrator"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatalf("Config error: %v", err)
	}

	// настройка коннекта к БД
	helpers.InitDatabase(config.Database)

	// выполнение миграций
	migrator.CreateTables(helpers.DB)

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
