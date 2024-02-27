package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config содержит настройки сервера
type Config struct {
	// Порт, на котором запускается сервер
	Port string
	// Время ожидания запроса
	Timeout time.Duration
	// Время простоя до закрытия соединения
	IdleTimeout time.Duration
}

// InitConfig загружает настройки из файла .env и возвращает Config и ошибку, если таковая возникла
func InitConfig(cfgPath string) (Config, error) {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load(cfgPath)
	if err != nil {
		return Config{}, err
	}

	// Парсим значения таймаутов из переменных окружения
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return Config{}, err
	}
	idleTimeout, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		return Config{}, err
	}


	return Config{
		Port:        os.Getenv("APP_PORT"),
		Timeout:     timeout,
		IdleTimeout: idleTimeout,
	}, nil
}
