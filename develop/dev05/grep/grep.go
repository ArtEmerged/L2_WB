package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

func RunGrep(reader io.Reader, writer io.Writer, options *Options) error {
	if options.Fixed {
		options.Pattern = regexp.QuoteMeta(options.Pattern)
	}

	if options.IgnoreCase {
		options.Pattern = "(?i)" + options.Pattern
	}

	if options.After == 0 {
		options.After = options.Context
	}
	if options.Before == 0 {
		options.Before = options.Context
	}

	if len(options.FilesName) == 0 {
		return grep(reader, writer, options)

	}
	for _, fileName := range options.FilesName {

		file, err := os.Open(fileName)
		if err != nil {
			return ErrNoSuchFile
		}
		defer file.Close()
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

func findLine(lines []string, pattern string, invert bool) (map[int]struct{}, error) {
	setLinesIndex := map[int]struct{}{}
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for index, line := range lines {
		mutch := rex.MatchString(line)
		if (mutch && !invert) || (!mutch && invert) {
			setLinesIndex[index] = struct{}{}
		}
	}
	return setLinesIndex, err
}

func grep(reader io.Reader, writer io.Writer, options *Options) error {
	splitLinesText := loadData(reader)
	setLinesIndex, err := findLine(splitLinesText, options.Pattern, options.Invert)
	if err != nil {
		return err
	}

	return writeResul(writer, splitLinesText, setLinesIndex, options)
}

func writeResul(out io.Writer, linesText []string, setLinesIndex map[int]struct{}, options *Options) error {
	outWriter := bufio.NewWriter(out)
	defer outWriter.Flush()

	var err error
	if options.Count {
		_, err = outWriter.WriteString(fmt.Sprintf("%d\n", len(setLinesIndex)))
		return err
	}

	setIndexContext := map[int]struct{}{}
	indexSlice := []int{}

	for index := range setLinesIndex {
		afterIndex := index + options.After
		beforeIndex := index - options.Before
		if beforeIndex < 0 {
			beforeIndex = 0
		}
		if afterIndex > len(linesText)-1 {
			afterIndex = len(linesText) - 1
		}

		for addIndex := beforeIndex; addIndex <= afterIndex; addIndex++ {
			if _, ok := setIndexContext[addIndex]; !ok {
				indexSlice = append(indexSlice, addIndex)
			}
			setIndexContext[addIndex] = struct{}{}
		}
	}

	sort.Ints(indexSlice)

	for _, indexLine := range indexSlice {
		if options.LineNum {
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

// loadData -  выполняет чтение из переданного файла
// возвращает []string - содержимое файла разделенное на \n, err == nil
// если произошла ошибка при открытии файла, то возвращает err
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
