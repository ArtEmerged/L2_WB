package cut

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Run - функция для запуска утилиты cut
func Run(reader io.Reader, writer io.Writer, options *Options) error {
	// Получаем все поля для вывода из флага -f
	fields, err := getFields(options.Fields)
	if err != nil {
		return err
	}

	// Возвращаем ошибку, если передали не один символ
	if utf8.RuneCountInString(options.Delimiter) != 1 {
		return ErrNotCharacter
	}

	// Читаем из stdin, если файл не передан
	if len(options.FilesName) == 0 {
		return cut(reader, writer, fields, options.Delimiter, options.Separated)
	}

	// Проходимся по всем файлам и выполняем функцию cut
	for _, fileName := range options.FilesName {
		file, err := os.Open(fileName)
		if err != nil {
			return ErrNoSuchFile
		}
		defer file.Close()
		err = cut(file, writer, fields, options.Delimiter, options.Separated)
		if err != nil {
			return err
		}
	}

	return nil
}

// cut - Формирует и записывает результат выполнения утилиты cut
func cut(reader io.Reader, writer io.Writer, fields Fields, delimiter string, separated bool) error {
	// Сплитим текст по строчкам
	splitLinesText := loadData(reader)
	outWriter := bufio.NewWriter(writer)
	defer outWriter.Flush()

	// Проходимся циклом по строчкам
	for _, line := range splitLinesText {
		// Разделяем строку на delimiter
		splitLine := strings.Split(line, delimiter)

		lenSplit := len(splitLine)
		minFromToEnd := min(fields.FromToEnd, lenSplit)

		// Получаем номер самой последней колонки, но не больше количества колонок в строке
		maxField := getMaxField(fields.FieldsList, lenSplit, fields.FromStartTo, fields.FromToEnd)

		// Пропускаем строки, которые не разделились и флаг -s == true
		if lenSplit == 1 && separated {
			continue
		}

		// Выводим строку, которая не разделилась и нет флага -s
		if lenSplit == 1 {
			if _, err := outWriter.WriteString(line + "\n"); err != nil {
				return err
			}
			continue
		}

		// Выводим колонки от 1 до NUM, если в флаг -f передали -NUM
		if fields.FromStartTo != EmptyStart {
			for numField := 1; numField <= fields.FromStartTo; numField++ {
				if err := writeOutput(outWriter, splitLine[numField-1], delimiter, numField, maxField); err != nil {
					return err
				}
			}
		}

		// Выводим переданные колонки в флаг -f
		for _, numField := range fields.FieldsList {
			// Если колонка больше -NUM и не больше NUM-
			if numField > fields.FromStartTo && numField <= minFromToEnd {
				if err := writeOutput(outWriter, splitLine[numField-1], delimiter, numField, maxField); err != nil {
					return err
				}
			}
		}

		// Выводим колонки от NUM до конца, если в флаг -f передали NUM-
		if fields.FromToEnd != EmptyEnd {
			for numField := minFromToEnd; numField <= lenSplit; numField++ {
				if err := writeOutput(outWriter, splitLine[numField-1], delimiter, numField, lenSplit); err != nil {
					return err
				}
			}
		}
		if _, err := outWriter.WriteString("\n"); err != nil {
			return err
		}
	}
	return nil
}

// writeOutput - записывает результат в буфер
func writeOutput(out *bufio.Writer, text, delimiter string, index, maxIndex int) error {
	// Если последнее поле, то delimiter не добавляем
	if index == maxIndex {
		if _, err := out.WriteString(text); err != nil {
			return err
		}
		return nil
	}
	if _, err := out.WriteString(text + delimiter); err != nil {
		return err
	}
	return nil
}

// loadData - разделяет текст по строчкам и формирует срез типа string
func loadData(reader io.Reader) []string {
	// Объявляем переменную res для записи файла построчно
	res := make([]string, 0)

	// Объявляем новый сканер buffS
	buffS := bufio.NewScanner(reader)
	// Проходим циклом по всему файлу
	for buffS.Scan() {
		// Добавляем в переменную res построчные данные типа string с помощью метода Text()
		res = append(res, buffS.Text())
	}
	return res
}

// getFields - формирует параметры для флага -f (fields) из переданных значений в флаг -f
// возвращает ошибку, если значение переданного значения невалидно
func getFields(flagFields string) (Fields, error) {
	var fromStartTo int = EmptyStart
	var fromToEnd int = EmptyEnd
	list := []int{0}
	set := map[int]struct{}{}

	listOfFields := strings.Split(flagFields, ",")

	for _, field := range listOfFields {

		if field == "" {
			return Fields{}, ErrNumberLessOne
		}
		// Получаем интервал от 1 до NUM (-NUM)
		if strings.HasPrefix(field, "-") {
			to, err := getNumber(field[1:])
			if err != nil {
				return Fields{}, err
			}
			fromStartTo = max(fromStartTo, to)
			continue
		}
		// Получаем интервал от NUM до конца (NUM-)
		if strings.HasSuffix(field, "-") {
			from, err := getNumber(field[:len(field)-1])
			if err != nil {
				return Fields{}, err
			}
			fromToEnd = min(fromToEnd, from)
			continue
		}

		interval := strings.Split(field, "-")

		// Обрабатываем число
		if len(interval) == 1 {
			column, err := getNumber(interval[0])
			if err != nil {
				return Fields{}, err
			}
			// Добавляем число в срез, если его ещё нет в множестве
			if _, ok := set[column]; !ok {
				list = append(list, column)
			}
			// Формируем множество чисел
			set[column] = struct{}{}
			continue
		}

		// Обрабатываем интервал
		if len(interval) == 2 {
			a, err := getNumber(interval[0])
			if err != nil {
				return Fields{}, err
			}
			b, err := getNumber(interval[1])
			if err != nil {
				return Fields{}, err
			}
			if b < a {
				return Fields{}, ErrInvalidRange
			}
			for column := a; column <= b; column++ {
				// Добавляем число в срез, если его ещё нет в множестве
				if _, ok := set[column]; !ok {
					list = append(list, column)
				}
				// Формируем множество чисел
				set[column] = struct{}{}
			}
			continue
		}
		// Если мы не нашли ни одно условие, то данные невалидны
		return Fields{}, ErrInvalidRange
	}
	sort.Ints(list)
	return Fields{
		FromStartTo: fromStartTo,
		FromToEnd:   fromToEnd,
		FieldsList:  list}, nil
}

// getNumber - конвертирует string в int
// возвращает ошибку ErrNotNumber если Atoi вернул ошибку
// возвращает ошибку ErrNubmerLessOne если число < 1
func getNumber(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0, ErrNotNumber
	}
	if number < 1 {
		return 0, ErrNumberLessOne
	}
	return number, nil
}

// getMaxField - определяет максимальную колонку из всех переданных в флаг -f
// которая не больше количества колонок в строке (lenSplit)
func getMaxField(fieldsList []int, lenSplit, fromStartTo, fromToEnd int) int {
	var maxField int
	for _, field := range fieldsList {
		if field <= lenSplit {
			maxField = field
		}
	}
	a := min(lenSplit, maxField)
	b := min(lenSplit, fromStartTo)
	c := min(lenSplit, fromToEnd)
	if fromToEnd == EmptyEnd {
		return max(a, b)
	}
	return max(a, max(b, c))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
