package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Run - запуск UNIX-шелл-утилиту
func Run(in io.Reader, out io.Writer) error {
	var res string
	read := bufio.NewReader(in)

	for {
		// Определяем рабочую директорию
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		// Выводим строку для ввода
		fmt.Fprintf(os.Stdout, "%s >", pwd)

		// Читаем текст до \n
		input, err := read.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		// Убираем лишние пробелы
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		// Проверяем наличие пайплайна
		pipe := strings.Contains(input, " | ")
		// Запускаем обработчик если есть пайплайн
		if pipe {
			splitArgs := strings.Split(input, " | ")
			// Создаем структуру с командами разделёнными на пайплайн
			commandParams := NewCommandParams("", splitArgs)
			// Вызываем метод для обработки пайплайна
			res, err = commandParams.PipeCommand()
		}
		// Запускаем обработчик если нет пайплайна
		if !pipe {
			splitArgs := strings.Fields(input)
			// Создаем структуру с командой и аргументами
			commandParams := NewCommandParams(splitArgs[0], splitArgs[1:])
			// Вызываем метод для обработки команд
			res, err = commandParams.ExecutesCommand(in)
		}

		if err != nil {
			// Если ошибка не равна EmptyArg то печатаем её
			if err.Error() != ErrEmptyArg.Error() {
				fmt.Fprintln(out, err.Error())
			}
			continue
		}
		// Если результат == выходу, то завершаем выполнение
		if res == quit {
			return nil
		}
		// Выводим результат команды, если это не пустая строка
		if res != "" {
			fmt.Fprintln(out, res)
		}
	}
}
