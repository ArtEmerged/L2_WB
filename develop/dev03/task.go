package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrArgs       = errors.New("sort: the number of argems has been exceeded")
	ErrColumn     = errors.New("sort: invalid number in column option")
	ErrNoSuchFile = errors.New("sort: no such file or directory")
)

type Parameters struct {
	columnFlag int
	numberFlag bool
	reversFlag bool
	uniqueFlag bool
	args       []string
}

func main() {
	// str, _ := strconv.Atoi("200hello")
	// fmt.Println(str)
	parameters := Parameters{}
	flag.IntVar(&parameters.columnFlag, "k", 1, "")      // -k — указание колонки для сортировки
	flag.BoolVar(&parameters.numberFlag, "n", false, "") // -n — сортировать по числовому значению
	flag.BoolVar(&parameters.reversFlag, "r", false, "") // -r — сортировать в обратном порядке
	flag.BoolVar(&parameters.uniqueFlag, "u", false, "") // -u — не выводить повторяющиеся строки
	flag.Parse()

	parameters.args = flag.Args()
	err := Run(parameters)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(parameters Parameters) error {
	// Проверяем правильность переданных аргументов
	err := validate(parameters.columnFlag, parameters.args)
	if err != nil {
		return err
	}

	splitLines, err := readFile(flag.Arg(0))
	if err != nil {
		return err
	}

	// Если флаг -u == true, то формируем множество из строк
	if parameters.uniqueFlag {
		splitLines = leaveUnique(splitLines)
	}

	sort(splitLines, parameters.columnFlag, parameters.numberFlag, parameters.reversFlag)
	outputResult(splitLines)
	return nil
}

// validate - Проверяет переданные аргументы на соответсвие требованям для корректонй работы программы.
// Если ошибки небыли выявленны, то функция возвращает nil
func validate(columnFlag int, args []string) error {
	// Если было переданно больше чем один аргумент, то возвращаем ошибку ErrArgs
	if len(args) != 1 {
		return ErrArgs
	}
	// Если указанная колонка для сортировки меньше одного, то возвращаем ошибку ErrColumn
	if columnFlag < 1 {
		return ErrColumn
	}
	return nil
}

// leaveUnique - возвращает множество из строк
func leaveUnique(lines []string) []string {
	// Обявляем map set для нахождения уникальных строк
	set := map[string]struct{}{}
	// Проходимся циклом по всем линиям
	for i := 0; i < len(lines); i++ {
		// Если строка уже есть в нашем множестве, то мы меняем её с последним елементом
		// и уменьшаем длину среза на один
		if _, ok := set[lines[i]]; ok {
			// Меняем местами повторяющуюся строчку с последней
			lines[i], lines[len(lines)-1] = lines[len(lines)-1], lines[i]
			// Уменшьнам длину среза на 1
			lines = lines[:len(lines)-1]
		}
		set[lines[i]] = struct{}{}
	}
	return lines
}

// readFile -  выполняет чтение из переданого файла
// возвращает []string - содержимое файла разделенное на \n, err == nil
// если произошла ошибка при открытие файла, то возращает err
func readFile(nameFile string) ([]string, error) {
	// Открываем файл который клиент передал через аргумент
	file, err := os.Open(nameFile)
	// Если путь к файлу не найден, то возвращаем ошибку об отсутствии файла
	if err != nil {
		return nil, ErrNoSuchFile
	}
	// Выполняем отложенный вызов закрытия файла
	defer file.Close()

	// Объявляем переменную res для записи файла построчно
	res := make([]string, 0)

	// Обявляем новый сканер buffS
	buffS := bufio.NewScanner(file)

	// Проходимя циклом по всему файлу
	for buffS.Scan() {
		// Добавляем в переменую res построчные данные типа string c помощью метода Text()
		res = append(res, buffS.Text())
	}

	return res, buffS.Err()
}

func outputResult(output []string) {
	for _, line := range output {
		fmt.Println(line)
	}
}

func sort(lines []string, columnFlag int, numberFlag, reversFlag bool) {
	slices.SortFunc(lines, func(a, b string) int {
		//Если флаг -r == true, то сортируем в обратной последовательности
		if reversFlag {
			return compare(b, a, columnFlag, numberFlag)
		}
		return compare(a, b, columnFlag, numberFlag)
	})

}

func compare(a, b string, column int, numberFlag bool) int {
	if column > 1 {
		splitA := strings.Fields(a)
		splitB := strings.Fields(b)
		if column <= len(splitA) && column <= len(splitB) {
			a = splitA[column-1]
			b = splitB[column-1]
		}
	}
	if numberFlag {
		return compareNumbers(a, b)
	}

	return strings.Compare(a, b)
}

func compareNumbers(a, b string) int {
	var numA int
	var numB int
	for i, symbol := range a {
		if !unicode.IsDigit(symbol) {
			numA, _ = strconv.Atoi(a[:i])
		}
	}
	for i, symbol := range b {
		if !unicode.IsDigit(symbol) {
			numB, _ = strconv.Atoi(b[:i])
		}
	}
	if numA < numB {
		return -1
	}
	if numA > numB {
		return +1
	}
	return 0
}
