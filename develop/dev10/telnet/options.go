package telnet

import (
	"errors"
	"flag"
	"time"
)

// Ошибка при недостаточном количестве аргументов
var (
	ErrMissArgs = errors.New("telnet: insufficient arguments; telnet [host] [port]")
)

// Options - структура для хранения опций утилиты telnet
type Options struct {
	Host    string        // Хост
	Port    string        // Порт
	Timeout time.Duration // Таймаут
}

// NewOptions - конструктор Options
func NewOptions(host, port string, timeout time.Duration) Options {
	return Options{
		Host:    host,
		Port:    port,
		Timeout: timeout,
	}
}

// ParsOptions - парсит переданные аргументы и возвращает структуру Options
func ParsOptions(args []string) (Options, error) {
	tcpSet := flag.NewFlagSet("telnet", flag.ExitOnError)
	timeout := tcpSet.Duration("timeout", time.Second*10, "--timeout 10s")
	tcpSet.DurationVar(timeout, "t", time.Second*10, "-t 10s")

	if err := tcpSet.Parse(args); err != nil {
		return Options{}, err
	}
	args = tcpSet.Args()
	// Проверяем наличие хоста и порта 
	if len(args) != 2 {
		return Options{}, ErrMissArgs
	}
	
	return NewOptions(args[0], args[1], *timeout), nil
}
