package telnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Сообщения о закрытии соединения
const (
	ConnCloseByClient = "Connection closed."
	ConnCloseByHost   = "Connection closed by foreign host."
)

// Run - функция для запуска telnet клиента
func Run(in io.Reader, out io.Writer, args []string) error {
	// Парсинг опций
	opt, err := ParsOptions(args)
	if err != nil {
		return err
	}

	// Устанавливаем соединение с сервером
	conn, err := net.DialTimeout("tcp", opt.Host+":"+opt.Port, opt.Timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Обработка сигналов завершения программы (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Канал для сигнала завершения
	done := make(chan struct{})

	go func() {
		defer close(done)
		rIn := bufio.NewReader(in)
		rConn := bufio.NewReader(conn)

		// Читаем данные из входного потока
		for {
			data, err := rIn.ReadBytes('\n')
			if err != nil {
				// Обработка сигнала завершения ввода (Ctrl+D)
				if err == io.EOF {
					fmt.Fprintln(out, ConnCloseByClient)
					return
				}
				fmt.Fprintln(out, err)
			}

			// Отправка данных на сервер
			_, err = conn.Write(data)
			if err != nil {
				fmt.Fprintln(out, err)
			}

			// Получаем ответ от сервера
			text, err := rConn.ReadString('\n')
			if err != nil {
				fmt.Fprintln(out, ConnCloseByHost)
				return
			}

			// Выводит ответ в stdout
			fmt.Fprint(out, text)
		}
	}()

	// Ожидаем завершения программы
	select {
	case <-signalChan:
		close(signalChan)
		fmt.Fprintln(out, ConnCloseByClient)
	case <-done:
	}
	return nil
}
