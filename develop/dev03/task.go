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
	// ErrArgs - ошибка некорректной передачи аргументов
	ErrArgs = errors.New("sort: the number of arguments has been exceeded")
	// ErrColumn - ошибка некорректной передачи колонки для сортировки
	ErrColumn = errors.New("sort: invalid number in column option")
	// ErrNoSuchFile - не удалось прочитать файл
	ErrNoSuchFile = errors.New("sort: no such file or directory")
)

// Parameters - структура для передачи флагов и аргументов и исполняющую функцию
type Parameters struct {
	columnFlag  int
	numberFlag  bool
	reverseFlag bool
	uniqueFlag  bool
	args        []string
}

func main() {
	parameters := Parameters{}
	flag.IntVar(&parameters.columnFlag, "k", 1, "-k 1 specify the column for sorting")                  // -k — указание колонки для сортировки
	flag.BoolVar(&parameters.numberFlag, "n", false, "-n  compare according to string numerical value") // -n — сортировать по числовому значению
	flag.BoolVar(&parameters.reverseFlag, "r", false, "-r reverse the result of comparisons")           // -r — сортировать в обратном порядке
	flag.BoolVar(&parameters.uniqueFlag, "u", false, "-u unique")                                       // -u — не выводить повторяющиеся строки
	flag.Parse()

	parameters.args = flag.Args()
	res, err := Run(parameters)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	outputResult(res)
}

// Run - выполняем процесс чтения и обработки данных для сортировки
func Run(parameters Parameters) ([]string, error) {
	// Проверяем правильность переданных аргументов
	err := validate(parameters.columnFlag, parameters.args)
	if err != nil {
		return nil, err
	}

	splitLines, err := readFile(parameters.args[0])
	if err != nil {
		return nil, err
	}

	// Если флаг -u == true, то формируем множество из строк
	if parameters.uniqueFlag {
		splitLines = leaveUnique(splitLines)
	}

	sort(splitLines, parameters.columnFlag, parameters.numberFlag, parameters.reverseFlag)
	return splitLines, nil
}

// validate - Проверяет переданные аргументы на соответствие требованиям для корректной работы программы.
// Если ошибки не были выявлены, то функция возвращает nil
func validate(columnFlag int, args []string) error {
	// Если было передано больше чем один аргумент, то возвращаем ошибку ErrArgs
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
	// Объявляем map set для нахождения уникальных строк
	set := map[string]struct{}{}
	// Объявляем новый срез для хранения множеств
	setSlice := make([]string, 0, len(lines))
	// Проходимся циклом по всем линиям
	for i := 0; i < len(lines); i++ {
		// Если строки ещё нет в нашем множестве, то мы кладем её в срез и в map
		if _, ok := set[lines[i]]; !ok {
			setSlice = append(setSlice, lines[i])
			set[lines[i]] = struct{}{}
		}
	}
	return setSlice
}

// readFile -  выполняет чтение из переданного файла
// возвращает []string - содержимое файла разделенное на \n, err == nil
// если произошла ошибка при открытии файла, то возвращает err
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

	// Объявляем новый сканер buffS
	buffS := bufio.NewScanner(file)

	// Проходим циклом по всему файлу
	for buffS.Scan() {
		// Добавляем в переменную res построчные данные типа string с помощью метода Text()
		res = append(res, buffS.Text())
	}

	return res, buffS.Err()
}

// outputResult - выводит в stdout результат
func outputResult(output []string) {
	for _, line := range output {
		fmt.Println(line)
	}
}

// sort - выполняет сортировку в соответствии с флагами
func sort(lines []string, columnFlag int, numberFlag, reverseFlag bool) {
	slices.SortFunc(lines, func(a, b string) int {
		// Если флаг -r == true, то сортируем в обратной последовательности
		if reverseFlag {
			return compare(b, a, columnFlag, numberFlag)
		}
		return compare(a, b, columnFlag, numberFlag)
	})

}

// compare - сравнивает строки согласно переданным флагам
func compare(a, b string, column int, numberFlag bool) int {
	// Если column больше одного, то выполняем сортировку по колонке
	if column > 1 {
		// Разделяем a и b по пробелам
		splitA := strings.Fields(a)
		splitB := strings.Fields(b)
		// Присваиваем a и b новые значения согласно колонке
		if column <= len(splitA) && column <= len(splitB) {
			a = splitA[column-1]
			b = splitB[column-1]
		}
	}
	// Если флаг numberFlag == true, то проверяем числа в строках
	if numberFlag {
		comp, ok := compareNumbers(a, b)
		if ok {
			return comp
		}
	}

	return strings.Compare(a, b)
}

// compareNumbers - сравнивает числа в префиксах строк a и b,
// a > b return +1, true
// a < b return -1, true
// a == b return 0, true
// если обе строки не содержат числа в префиксе, то возвращает false
func compareNumbers(a, b string) (int, bool) {
	var numA, indexA int
	var numB, indexB int
	// Проверяем циклом где заканчивается число
	for _, symbol := range a {
		if !unicode.IsDigit(symbol) {
			break
		}
		indexA++
	}
	for _, symbol := range b {
		if !unicode.IsDigit(symbol) {
			break
		}
		indexB++
	}
	// Возвращаем false если оба цикла завершились на первой итерации
	if indexA == 0 && indexB == 0 {
		return 0, false
	}
	numA, _ = strconv.Atoi(a[:indexA])
	numB, _ = strconv.Atoi(b[:indexB])
	if numA < numB {
		return -1, true
	}
	if numA > numB {
		return +1, true
	}
	return 0, true
}
