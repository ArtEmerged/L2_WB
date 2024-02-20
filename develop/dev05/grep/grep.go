package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

// Run - определяет откуда нужно читать данные и вызывает функцию grep()
func Run(reader io.Reader, writer io.Writer, options *Options) error {

	// Читаем из stdin, если файл не передан
	if len(options.FilesName) == 0 {
		return grep(reader, writer, options)
	}

	// Проходимся по всем файлам и выполняем функцию grep
	for _, fileName := range options.FilesName {
		file, err := os.Open(fileName)
		if err != nil {
			return ErrNoSuchFile
		}
		defer file.Close()
		// Если количество файлов больше одного, то печатаем путь к файлу
		if len(options.FilesName) > 1 {
			_, err = writer.Write([]byte(fileName + ":"))
			if err != nil {
				return err
			}
		}
		err = grep(file, writer, options)
		if err != nil {
			return err
		}
	}
	return nil
}

// findLine - возвращает set с индексами среза, где был найден паттерн
func findLine(lines []string, pattern string, invert bool) (map[int]struct{}, error) {
	setLinesIndex := map[int]struct{}{}
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for index, line := range lines {
		match := rex.MatchString(line)
		// Добавляем в set индексы среза, где был паттерн или где его не было, если флаг -v == true (invert)
		if (match && !invert) || (!match && invert) {
			setLinesIndex[index] = struct{}{}
		}
	}
	return setLinesIndex, err
}

// grep - функция выполняет поиск паттерна из io.Reader, формирует результат и выводит в io.Writer в соответствии с переданными флагами
func grep(reader io.Reader, writer io.Writer, options *Options) error {
	// Если передан флаг -F fixed, то мы экранируем все метасимволы регулярного выражения
	if options.Fixed {
		options.Pattern = regexp.QuoteMeta(options.Pattern)
	}

	// Если передан флаг -i Ignore Case, то добавляем к паттерну игнорирование регистра
	if options.IgnoreCase {
		options.Pattern = "(?i)" + options.Pattern
	}

	// Задаем новые значения After и Before, если они равны нулю
	if options.After == 0 {
		options.After = options.Context
	}
	if options.Before == 0 {
		options.Before = options.Context
	}

	// Сплитим текст по строчкам
	splitLinesText := loadData(reader)
	// Определяем, в каких строчках текста есть паттерн
	setLinesIndex, err := findLine(splitLinesText, options.Pattern, options.Invert)
	if err != nil {
		return err
	}

	return writeResult(writer, splitLinesText, setLinesIndex, options)
}

// writeResult - формирует результат найденного паттерна в тексте в соответствии с переданными флагами и выводит в io.Writer
func writeResult(out io.Writer, linesText []string, setLinesIndex map[int]struct{}, options *Options) error {
	outWriter := bufio.NewWriter(out)
	defer outWriter.Flush()

	var err error

	// Если флаг -c == true (count), то возвращаем количество строк с паттерном
	if options.Count {
		_, err = outWriter.WriteString(fmt.Sprintf("%d\n", len(setLinesIndex)))
		return err
	}

	setIndexContext := map[int]struct{}{}
	indexSlice := []int{}

	for index := range setLinesIndex {
		// Определяем количество строк для вывода (до и после найденного паттерна)
		afterIndex := index + options.After
		beforeIndex := index - options.Before
		if beforeIndex < 0 {
			beforeIndex = 0
		}
		if afterIndex > len(linesText)-1 {
			afterIndex = len(linesText) - 1
		}

		// Проходимся по индексам от beforeIndex до afterIndex и добавляем их в срез
		for addIndex := beforeIndex; addIndex <= afterIndex; addIndex++ {
			// Если с set еще нет индекса, то добавляем его в срез
			if _, ok := setIndexContext[addIndex]; !ok {
				indexSlice = append(indexSlice, addIndex)
			}
			setIndexContext[addIndex] = struct{}{}
		}
	}

	// Сортируем все добавленные индексы строк
	sort.Ints(indexSlice)

	// Проходимся циклом по всем индексам строк и добавляем их в буфер
	for _, indexLine := range indexSlice {
		// Если флаг -n == true (line num), то выводим номера строк
		if options.LineNum {
			// Если в строке был найден паттерн, то после номера строки ставим ":", иначе "-"
			if _, ok := setLinesIndex[indexLine]; ok {
				_, err = outWriter.WriteString(fmt.Sprintf("%d:", indexLine+1))
			} else {
				_, err = outWriter.WriteString(fmt.Sprintf("%d-", indexLine+1))
			}
			if err != nil {
				return err
			}
		}
		_, err = outWriter.WriteString(linesText[indexLine] + "\n")
		if err != nil {
			return err
		}
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
