package shell

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Константа для выхода из программы
const quit = "quit"

var (
	// ErrEmptyArg - ошибка при пустом аргументе
	ErrEmptyArg = errors.New("empty arg")
	// ErrCdTooManyArg - ошибка при слишком большом числе аргументов для команды "cd"
	ErrCdTooManyArg = errors.New("shell: cd: too many arguments")
	// ErrCdNoSuchFileOrDir - ошибка при попытке изменить директорию на несуществующую
	ErrCdNoSuchFileOrDir = errors.New("shell: cd: no such file or directory")
	// ErrKill - ошибка при неправильном использовании команды "kill"
	ErrKill = errors.New("shell: kill: usage: kill <PID>")
	// ErrKillNoSuchProcess - ошибка при попытке завершения несуществующего процесса
	ErrKillNoSuchProcess = errors.New("shell: kill: no such process")
)

// CommandParams - структура для хранения параметров команды
type CommandParams struct {
	command string   // Команда
	arg     []string // Аргументы команды
}

// NewCommandParams - конструктор для CommandParams
func NewCommandParams(command string, arg []string) *CommandParams {
	return &CommandParams{
		command: command,
		arg:     arg,
	}
}

// ExecutesCommand - выполняет команду в зависимости от ее типа
func (cp *CommandParams) ExecutesCommand(in io.Reader) (result string, err error) {
	switch cp.command {
	default:
		// Выполняем команду "exec"
		result, err = cp.exec(in)
	case "cd":
		// Выполняем команду "cd"
		err = cp.cd()
	case "pwd":
		// Выполняем команду "pwd"
		result, err = cp.pwd()
	case "kill":
		// Выполняем команду "kill"
		err = cp.kill()
	case "echo":
		// Выполняем команду "echo"
		result = cp.echo()
	case "ps":
		// Выполняем команду "ps"
		result, err = cp.ps()
	case "/quit":
		// Завершение работы программы
		result = quit
	}
	return
}

// PipeCommand - выполняет команду через пайплайн
func (cp *CommandParams) PipeCommand() (result string, err error) {
	// Проверка на пустые аргументы
	if len(cp.arg) == 0 {
		return "", ErrEmptyArg
	}

	var resBuff bytes.Buffer

	// Проход циклом по каждой команде в пайплайне
	for _, command := range cp.arg {
		// Разделение команды на отдельные слова
		splitCommand := strings.Fields(command)

		// Проверка на пустые аргументы в команде
		if len(splitCommand) == 0 {
			return "", ErrEmptyArg
		}

		// Обновление команды и аргументов в структуре CommandParams
		cp.command = splitCommand[0]
		cp.arg = splitCommand[1:]

		// Выполняем команду и записываем результат в буфер
		res, err := cp.exec(&resBuff)
		if err != nil {
			return "", err
		}
		_, err = resBuff.WriteString(res)
		if err != nil {
			return "", err
		}
	}
	return resBuff.String(), nil
}

// cd - выполнение команды изменения директории
func (cp *CommandParams) cd() error {
	// Измененяем директорию на домашнюю если не был передан путь
	if len(cp.arg) == 0 {
		return os.Chdir(os.Getenv("HOME"))
	}
	// Возвращаем ошику если слишком много аргументов для команды "cd"
	if len(cp.arg) > 1 {
		return ErrCdTooManyArg
	}
	err := os.Chdir(cp.arg[0])
	// Возвращаем ошибку если директория не найдена
	if err != nil {
		return ErrCdNoSuchFileOrDir
	}
	return nil
}

// pwd - выполнение команды вывода текущей директории
func (cp *CommandParams) pwd() (string, error) {
	// Получаем текущую директорию
	return os.Getwd()
}

// echo - выполнение команды вывода аргументов
func (cp *CommandParams) echo() string {
	// Объединение аргументов в одну строку с пробелами между ними
	return strings.Join(cp.arg, " ")
}

// kill - выполнение команды завершения процесса
func (cp *CommandParams) kill() error {
	// Возвращаем ошибку если PID не был передан
	if len(cp.arg) == 0 {
		return ErrKill
	}
	for _, v := range cp.arg {
		pid, err := strconv.Atoi(v)
		// Возвращаем ошибку если не число
		if err != nil {
			return ErrKill
		}
		// Отправляем сигнала завершения процессу
		err = syscall.Kill(pid, syscall.SIGTERM)
		// Процесс не найден
		if err != nil {
			return ErrKillNoSuchProcess
		}
	}
	return nil
}

// ps - выполнение команды вывода списка процессов
func (cp *CommandParams) ps() (string, error) {
	var resBuff bytes.Buffer
	// Создание команды ps axu
	cp.command = "ps"
	cp.arg = []string{"axu"}
	return cp.exec(&resBuff)
}

// exec - выполнение команды
func (cp *CommandParams) exec(in io.Reader) (string, error) {
	var resBuff, errBuff bytes.Buffer
	// Создаем команды с заданными аргументами
	cmd := exec.Command(cp.command, cp.arg...)
	// Перенаправляем stdin
	cmd.Stdin = in
	// Записываем stdout в буфер
	cmd.Stdout = &resBuff
	// Записываем stderr в буфер
	cmd.Stderr = &errBuff
	err := cmd.Run()
	if err != nil {
		// Возвращаем ошибку, если команда завершилась с ошибкой
		return "", err
	}
	// Возвращаем ошибку из stderr, если она есть
	if errBuff.Len() != 0 {
		return "", errors.New(errBuff.String())
	}
	// Убираем \n в конце вывода команды
	res := strings.TrimSuffix(resBuff.String(), "\n")
	return res, nil
}
